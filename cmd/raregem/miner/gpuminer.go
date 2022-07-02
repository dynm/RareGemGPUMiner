package miner

import (
	"crypto/rand"
	"encoding/binary"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fatih/color"
	"github.com/robvanmieghem/go-opencl/cl"
)

/*
block.chainid uint256
entropy, bytes32
address(this), bytes20
msg.sender, bytes20
kind, uint256
nonce[msg.sender], uint256
salt uint256
00000000000000000000000000000000000000000000000000000000000000012d4c49471172aa8f4c5a0b3d495145b1e78105d2b42fbfb56423359c5acd4ae5dcc555576498e638decbe15f8c61c73f1a76a3d909d7b6c180aa1077665400194cc6f52950c1a884000000000000000000000000000000000000000000000000000000000000000700000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
*/

type SingleDeviceMiner struct {
	ClDevice           *cl.Device
	MinerID            int
	HashRateReportChan chan *HashRateReport
	MiningWorkChannel  chan []byte
	//Intensity defines the GlobalItemSize in a human friendly way, the GlobalItemSize = 2^Intensity
	Intensity      int
	GlobalItemSize int
	SaltChan       chan [32]byte
	// Client         clients.HeaderReporter
}

func (singleMiner *SingleDeviceMiner) Mine() {
	context, err := cl.CreateContext([]*cl.Device{singleMiner.ClDevice})
	if err != nil {
		log.Fatalln(singleMiner.MinerID, "-", err)
	}
	defer context.Release()

	commandQueue, err := context.CreateCommandQueue(singleMiner.ClDevice, 0)
	if err != nil {
		log.Fatalln(singleMiner.MinerID, "-", err)
	}
	defer commandQueue.Release()

	program, err := context.CreateProgramWithSource([]string{kernelSource})
	if err != nil {
		log.Fatalln(singleMiner.MinerID, "-", err)
	}
	defer program.Release()

	err = program.BuildProgram([]*cl.Device{singleMiner.ClDevice}, "")
	if err != nil {
		log.Fatalln(singleMiner.MinerID, "-", err)
	}

	kernel, err := program.CreateKernel("hashMessage")
	if err != nil {
		log.Fatalln(singleMiner.MinerID, "-", err)
	}
	defer kernel.Release()

	midstateObj := CreateEmptyBuffer(context, cl.MemReadOnly, 200)
	defer midstateObj.Release()
	kernel.SetArgBuffer(0, midstateObj)

	// blockTailObj := CreateEmptyBuffer(context, cl.MemReadOnly, 56)
	// defer blockTailObj.Release()
	// kernel.SetArgBuffer(1, blockTailObj)

	solutionsObj := CreateEmptyBuffer(context, cl.MemReadWrite, 256*8)
	defer solutionsObj.Release()
	kernel.SetArgBuffer(1, solutionsObj)

	solutionCountObj := CreateEmptyBuffer(context, cl.MemReadWrite, 4)
	defer solutionCountObj.Release()
	kernel.SetArgBuffer(2, solutionCountObj)

	localItemSize, err := kernel.WorkGroupSize(singleMiner.ClDevice)
	if err != nil {
		log.Fatalln(singleMiner.MinerID, "- WorkGroupSize failed -", err)
	}

	log.Println(singleMiner.MinerID, "- Global item size:", singleMiner.GlobalItemSize, "(Intensity", singleMiner.Intensity, ")", "- Local item size:", localItemSize)

	log.Println(singleMiner.MinerID, "- Initialized ", singleMiner.ClDevice.Type(), "-", singleMiner.ClDevice.Name())

	// solutions := make([]uint64, 256, 256)
	// dataPtr := unsafe.Pointer(&solutions[0])
	// dataSize := int(unsafe.Sizeof(solutions[0])) * len(solutions)
	// if _, err = commandQueue.EnqueueWriteBuffer(solutionsObj, true, 0, dataSize, dataPtr, nil); err != nil {
	// 	log.Fatalln(singleMiner.MinerID, "-", err)
	// }
	/*
		10000000000000000000000000000000000000000000000000000000000000012d4c49471172aa8f4c5a0b3d495145b1e78105d2b42fbfb56423359c5acd4ae5dcc555576498e638decbe15f8c61c73f1a76a3d909d7b6c180aa1077665400194cc6f52950c1a8840000000000000000000000000000000000000000000000000000000000000007
		00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001
		00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000080
	*/
	// blockHeader := common.Hex2Bytes("10000000000000000000000000000000000000000000000000000000000000012d4c49471172aa8f4c5a0b3d495145b1e78105d2b42fbfb56423359c5acd4ae5dcc555576498e638decbe15f8c61c73f1a76a3d909d7b6c180aa1077665400194cc6f52950c1a8840000000000000000000000000000000000000000000000000000000000000007ff000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")

	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	var work []byte
	for {
		// start := time.Now()
		continueMining := true
		select {
		case work, continueMining = <-singleMiner.MiningWorkChannel:
		default:
			// log.Println(singleMiner.MinerID, "-", "No work ready")
			// time.Sleep(time.Second * 1)

			// work, continueMining = <-singleMiner.MiningWorkChannel
			// log.Println(singleMiner.MinerID, "-", "Continuing")
		}
		if !continueMining {
			log.Println("Halting singleMiner ", singleMiner.MinerID)
			break
		}

		if len(work) < 200+32 {
			log.Println(singleMiner.MinerID, "-", "No work ready")
			continue
		}
		blockHeader := make([]byte, 200)
		copy(blockHeader, work[:200])
		targetHash := work[200 : 200+32]
		blockTail := blockHeader[136 : 136+64]
		padding := make([]byte, 136-len(blockTail))
		padding[0] = 0x01
		padding[len(padding)-1] = 0x80
		paddedTail := append(blockTail, padding...)
		d := crypto.NewKeccakState()
		d.Write(blockHeader[:136])
		midstate := GetMidState(d)
		// log.Printf("midstate(%d): %02x", len(midstate), midstate)
		// log.Printf("paded(%d): %02x", len(paddedTail), paddedTail)

		//Copy input to kernel args
		rand.Read(paddedTail[48 : 48+8])
		// log.Print("use random")

		xorInedMidstate := make([]byte, len(midstate))
		copy(xorInedMidstate, midstate)
		for i := range paddedTail[:136] {
			xorInedMidstate[i] ^= paddedTail[i]
		}
		// log.Printf("midstate:\t%02x", midstate)
		// log.Printf("midstate(xorin):\t%02x", xorInedMidstate)

		solutionCount := make([]byte, 4, 4)
		if _, err = commandQueue.EnqueueWriteBufferByte(solutionCountObj, true, 0, solutionCount, nil); err != nil {
			log.Fatalln(singleMiner.MinerID, "-", err)
		}

		if _, err = commandQueue.EnqueueWriteBufferByte(midstateObj, true, 0, xorInedMidstate[:], nil); err != nil {
			log.Fatalln(singleMiner.MinerID, "-", err)
		}

		// if _, err = commandQueue.EnqueueWriteBufferByte(blockTailObj, true, 0, paddedTail[:56], nil); err != nil {
		// 	log.Fatalln(singleMiner.MinerID, "-", err)
		// }

		start := time.Now()
		//Run the kernel
		if _, err = commandQueue.EnqueueNDRangeKernel(kernel, []int{0}, []int{singleMiner.GlobalItemSize}, []int{localItemSize}, nil); err != nil {
			log.Fatalln(singleMiner.MinerID, "-", err)
		}
		//Get output
		// dataPtr = unsafe.Pointer(&solutions[0])
		// dataSize = int(unsafe.Sizeof(solutions[0])) * len(solutions)
		out := make([]byte, 8*256)
		if _, err = commandQueue.EnqueueReadBufferByte(solutionsObj, true, 0, out, nil); err != nil {
			log.Fatalln(singleMiner.MinerID, "-", err)
		}
		// spew.Dump(out)

		if _, err = commandQueue.EnqueueReadBufferByte(solutionCountObj, true, 0, solutionCount, nil); err != nil {
			log.Fatalln(singleMiner.MinerID, "-", err)
		}
		hashRate := float64(singleMiner.GlobalItemSize) / (time.Since(start).Seconds() * 1000000)

		solutionCountN := binary.LittleEndian.Uint32(solutionCount)
		// log.Printf("solution count: %02x", solutionCount)
		// for i := 0; i < len(out)/8; i++ {
		// 	fmt.Printf("%02x\n", out[i*8:i*8+8])
		// 	if i > 24 {
		// 		break
		// 	}
		// }

		// log.Printf("solution cnt: %d", solutionCountN)
		if solutionCountN > 0 {
			for i := uint32(0); i < solutionCountN; i++ {
				var salt [32]byte
				outN := binary.BigEndian.Uint64(out[i*8 : i*8+8])
				reversed := make([]byte, 8)
				binary.BigEndian.PutUint64(reversed, outN)
				// reversed[7] = 0x0
				newHeader := append(blockHeader[:184], paddedTail[48:48+8]...)
				newHeader = append(newHeader, reversed...)
				hash := crypto.Keccak256Hash(newHeader)
				targetHash := common.BytesToHash(targetHash)
				coloredHash := ""
				if hash.Big().Cmp(targetHash.Big()) <= 0 {
					coloredHash = green(hash.Hex())
				} else {
					coloredHash = red(hash.Hex())
				}
				fmt.Printf("MinerID: %d\nHeader(%d): %02x\nTarget:\t%s\nHash:\t%s\n", singleMiner.MinerID, len(newHeader), newHeader, yellow(targetHash.Hex()), coloredHash)
				copy(salt[:], newHeader[200-32:])
				salt[0] = newHeader[135] //pass gemid
				singleMiner.SaltChan <- salt
			}
		}

		singleMiner.HashRateReportChan <- &HashRateReport{MinerID: singleMiner.MinerID, HashRate: hashRate}
		// break
	}
}

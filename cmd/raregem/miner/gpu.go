package miner

import (
	"log"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/robvanmieghem/go-opencl/cl"
)

type GPU struct {
	clDevices    []*cl.Device
	sdms         []*SingleDeviceMiner
	saltChan     chan [32]byte
	workChan     chan []byte
	hashRateChan chan *HashRateReport
}

func NewGPUMiner(platformArgs string, intensity int) *GPU {
	vendorSlice := strings.Split(platformArgs, ",")

	platforms, err := cl.GetPlatforms()
	if err != nil {
		log.Panic("[GetPlatforms] ", err)
	}

	gpu := &GPU{}
	gpu.saltChan = make(chan [32]byte)
	gpu.workChan = make(chan []byte)
	gpu.hashRateChan = make(chan *HashRateReport)

	minerID := 0
	for _, platform := range platforms {
		platormDevices, err := cl.GetDevices(platform, cl.DeviceTypeGPU)
		if err != nil {
			log.Fatalln("[GetDevices] ", err)
		}
		for _, device := range platormDevices {
			vendor := strings.ToLower(device.Vendor())
			for _, v := range vendorSlice {
				if strings.Contains(vendor, v) {
					log.Println(device.Type(), "-", device.Name())
					gpu.clDevices = append(gpu.clDevices, device)
					sdm := &SingleDeviceMiner{
						ClDevice:           device,
						MinerID:            minerID,
						Intensity:          intensity,
						GlobalItemSize:     int(math.Exp2(float64(intensity))),
						MiningWorkChannel:  gpu.workChan,
						SaltChan:           gpu.saltChan,
						HashRateReportChan: gpu.hashRateChan,
					}
					gpu.sdms = append(gpu.sdms, sdm)
					go sdm.Mine()
					minerID++
				}
			}
		}
	}
	return gpu
}

func (g *GPU) Mine(chainID *big.Int, Entropy [32]byte, Provably_Rare_Gem, From common.Address, gemKind, nonce, difficulty *big.Int) {
	var header []byte
	chainIDBytes := make([]byte, 32)
	chainID.FillBytes(chainIDBytes)
	header = append(header, chainIDBytes...)
	header = append(header, Entropy[:]...)
	header = append(header, Provably_Rare_Gem.Bytes()...)
	header = append(header, From.Bytes()...)

	gemKindBytes := make([]byte, 32)
	gemKind.FillBytes(gemKindBytes)
	header = append(header, gemKindBytes...)
	nonceBytes := make([]byte, 32)
	nonce.FillBytes(nonceBytes)
	header = append(header, nonceBytes...)

	salt := big.NewInt(0)
	saltBytes := make([]byte, 32)
	salt.FillBytes(saltBytes)
	header = append(header, saltBytes...)

	max256 := new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil)
	max256.Sub(max256, big.NewInt(1))
	target := new(big.Int).Div(max256, difficulty)
	targetHash := common.BigToHash(target)
	log.Printf("target: %s", targetHash.Hex())
	log.Printf("gemID: %d", header[135])

	for _, sdm := range g.sdms {
		sdm.MiningWorkChannel <- append(header, targetHash.Bytes()...)
	}
}

func (g *GPU) FoundChan() (foundChan chan *big.Int) {
	foundChan = make(chan *big.Int, 100)
	go func() {
		for {
			saltByte := <-g.saltChan
			saltBN := new(big.Int).SetBytes(saltByte[:])
			foundChan <- saltBN
		}
	}()
	return
}

func (g *GPU) PrintHashRate() {
	deviceNum := len(g.sdms)
	hashRateCntMap := make(map[int]bool)
	for i := 0; i < deviceNum; i++ {
		hashRateCntMap[i] = true
	}
	hashRateMap := make(map[int]float64)
	for hashRate := range g.hashRateChan {
		delete(hashRateCntMap, hashRate.MinerID)
		hashRateMap[hashRate.MinerID] = hashRate.HashRate
		if len(hashRateCntMap) == 0 {
			totalHashrate := float64(0)
			for i := 0; i < deviceNum; i++ {
				log.Printf("MinerID: %d, Hashrate: %.2f MH/s", i, hashRateMap[i])
				totalHashrate += hashRateMap[i]
			}
			log.Printf("Total Hashrate: %.2f MH/s", totalHashrate)
			// init the map
			for i := 0; i < deviceNum; i++ {
				hashRateCntMap[i] = true
			}
		}

	}
}

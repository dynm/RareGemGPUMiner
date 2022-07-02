package miner

import (
	"log"
	"math"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/robvanmieghem/go-opencl/cl"
)

func TestListDevices(t *testing.T) {
	platforms, err := cl.GetPlatforms()
	if err != nil {
		log.Panic("[GetPlatforms] ", err)
	}

	// spew.Dump(platforms)

	// var clDevice *cl.Device
	for _, platform := range platforms {
		platormDevices, err := cl.GetDevices(platform, cl.DeviceTypeGPU)
		if err != nil {
			log.Fatalln("[GetDevices] ", err)
		}
		for _, device := range platormDevices {
			log.Println(device.Type(), "-", device.Name())
			// clDevice = device
		}
	}
}

func TestKeccakGPU(t *testing.T) {
	platforms, err := cl.GetPlatforms()
	if err != nil {
		log.Panic("[GetPlatforms] ", err)
	}

	// spew.Dump(platforms)

	var clDevice *cl.Device
	for _, platform := range platforms {
		platormDevices, err := cl.GetDevices(platform, cl.DeviceTypeGPU)
		if err != nil {
			log.Fatalln("[GetDevices] ", err)
		}
		for _, device := range platormDevices {
			log.Println(device.Type(), "-", device.Name())
			clDevice = device
		}
	}
	sdm := SingleDeviceMiner{
		ClDevice:          clDevice,
		MinerID:           0,
		Intensity:         28,
		GlobalItemSize:    int(math.Exp2(float64(28))),
		MiningWorkChannel: make(chan []byte),
		SaltChan:          make(chan [32]byte),
	}
	go sdm.Mine()

	go func() {
		for salt := range sdm.SaltChan {
			log.Printf("salt: %02x", salt)
		}
	}()

	for {
		sdm.MiningWorkChannel <- common.Hex2Bytes("10000000000000000000000000000000000000000000000000000000000000012d4c49471172aa8f4c5a0b3d495145b1e78105d2b42fbfb56423359c5acd4ae5dcc555576498e638decbe15f8c61c73f1a76a3d909d7b6c180aa1077665400194cc6f52950c1a8840000000000000000000000000000000000000000000000000000000000000007ff000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	}
}

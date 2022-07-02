package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/dynm/RareGemGPUMiner/cmd/raregem/miner"
	"github.com/dynm/RareGemGPUMiner/wallet"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/fatih/color"
)

// var Provably_Rare_Gem = common.HexToAddress("0x33fFBFa82Cd2bbC5992347257c138b9a91081C66") // testnet

var (
	ACC_NUM = 3
)

func main() {
	var conn *ethclient.Client
	ctx := context.Background()
	var err error

	rpcURL := flag.String("rpc", "http://192.168.50.128:8545", "ethereum rpc url")
	accNum := flag.Int("n", 3, "derive n accounts")
	chainIDarg := flag.Int64("chainid", 1, "chain id")
	contractAddr := flag.String("contract", "0xc67ded0ec78b849e17771b2e8a7e303b4dad6dd4", "gem contract")
	gems := flag.String("gems", "8,9,18,19", "comma seperated gem kind id")
	interval := flag.Int64("interval", 10, "block interval")
	platform := flag.String("platform", "nvidia,amd", "specified opencl platform")
	intensity := flag.Int("intensity", 28, "gpu intensity")
	flag.Parse()

	_ = platform
	var Provably_Rare_Gem = common.HexToAddress(*contractAddr) // TODO: mainnet

	ACC_NUM = *accNum
	gemSlice := strings.Split(*gems, ",")
	conn, err = ethclient.Dial(*rpcURL)
	if err != nil {
		log.Panicf("dial err: %v", err)
	}

	RareGem, err := NewProvablyRareGem(Provably_Rare_Gem, conn)
	if err != nil {
		log.Panicf("NewProvablyRareGem err: %v", err)
	}

	w := wallet.NewWallet("mnemonic.txt", uint32(ACC_NUM), int64(*chainIDarg))
	auths := make([]*bind.TransactOpts, ACC_NUM)
	for i := 0; i < ACC_NUM; i++ {
		auths[i], _, _ = w.GetNextAuth(ctx, conn, nil)
	}

	officialConns := make([]*ethclient.Client, len(OfficialRPCs))
	for i := range OfficialRPCs {
		officialConns[i], err = ethclient.Dial(OfficialRPCs[i])
		if err != nil {
			log.Printf("Cannot dial %s, error: %v", OfficialRPCs[i], err)
			continue
		}
	}

	gpu := miner.NewGPUMiner(*platform, *intensity)
	go gpu.PrintHashRate()
	foundChan := gpu.FoundChan()

	var account *bind.TransactOpts
	accountLock := &sync.Mutex{}
	chainID := big.NewInt(*chainIDarg)
	log.Printf("chainID: %s", chainID)

	go func() {
		for {
			// for simplicity, only one account
			// get basic info from RareGem contract
			// nonce
			ts := time.Now()

			accountLock.Lock()
			account = auths[ts.UnixMilli()%int64(ACC_NUM)]
			accountLock.Unlock()
			log.Printf("🔑account: %s", account.From.Hex())
			balance, err := conn.BalanceAt(context.Background(), account.From, nil)
			if err != nil || balance.Cmp(big.NewInt(1e17)) <= 0 {
				log.Printf("insufficient balance")
				continue
			}
			nonce, err := RareGem.Nonce(nil, account.From)
			if err != nil {
				log.Printf("get Nonce err: %v", err)
				continue
			}
			log.Printf("account: %s, RareGem nonce: %d", account.From, nonce.Uint64())
			// get gem info
			var kind int
			kindStr := gemSlice[ts.UnixMilli()%int64(len(gemSlice))]
			kind, err = strconv.Atoi(kindStr)
			if err != nil {
				log.Print(err)
				continue
			}
			gemKind := big.NewInt(int64(kind))

			gemInfo, err := RareGem.Gems(nil, gemKind)
			if err != nil {
				log.Printf("get Gems err: %v", err)
				continue
			}
			log.Printf("gem kind id=%v, gemInfo=%+v", gemKind, gemInfo)

			gpu.Mine(chainID, gemInfo.Entropy, Provably_Rare_Gem, account.From, gemKind, nonce, gemInfo.Difficulty)
			intervalSecond := *interval
			<-time.After(time.Second * time.Duration(intervalSecond))
		}
	}()

	go func() {
		submittedDiff := make(map[int64]*big.Int)
		for found := range foundChan {
			salt := big.NewInt(0)
			log.Printf("found(with genID at byte[0]): %02x", found)
			foundBytes := make([]byte, 32)
			found.FillBytes(foundBytes)
			gemKind := big.NewInt(int64(foundBytes[0]))
			foundBytes[0] = 0x00
			salt.SetBytes(foundBytes)
			// DEBUG: check hash
			accountLock.Lock()
			callOpts := bind.CallOpts{From: account.From}
			accountLock.Unlock()
			hashResult, err := RareGem.Luck(&callOpts, gemKind, salt)
			if err != nil {
				log.Printf("get Luck err: %v", err)
				continue
			}
			log.Printf("hashResult=%s", common.BigToHash(hashResult).Hex()) // not sure about the format

			// check gem info again before submit, because mine need time
			latestGemInfo, err := RareGem.Gems(nil, gemKind)
			if err != nil {
				log.Printf("get Gems err: %v", err)
				continue
			}
			log.Printf("gem kind id=%v, latestGemInfo=%+v", gemKind, latestGemInfo)
			if sDiff, ok := submittedDiff[gemKind.Int64()]; !ok {
				submittedDiff[gemKind.Int64()] = big.NewInt(0)
			} else {
				if sDiff.Cmp(latestGemInfo.Difficulty) == 0 {
					log.Printf("same diff for gem[%d], skipping", gemKind)
					continue
				}
			}
			// submit to contract, fire tx
			tmpAccount := *account

			if *chainIDarg == 1 {
				tmpAccount.GasFeeCap = big.NewInt(2500 * 1e9) // should be large enough
			} else {
				suggestGasPrice, err := conn.SuggestGasPrice(ctx)
				if err != nil {
					log.Printf("get SuggestGasPrice err: %v", err)
					continue
				}
				amplifiedGasPrice := new(big.Int).Mul(suggestGasPrice, big.NewInt(110))
				amplifiedGasPrice.Div(amplifiedGasPrice, big.NewInt(100))
				tmpAccount.GasPrice = amplifiedGasPrice
			}
			tmpAccount.Nonce = nil

			tx, err := RareGem.Mine(&tmpAccount, gemKind, salt) // gasFeeCap not nil, then eip1559
			if err != nil {
				color.Red(fmt.Sprintf("send Mine err: %v", err))
				continue
			}
			color.Green(fmt.Sprintf("tx sent, tx hash: %s", tx.Hash()))
			submittedDiff[gemKind.Int64()] = latestGemInfo.Difficulty
			// TODO: next loop
			continue
		}
	}()

	select {}

}

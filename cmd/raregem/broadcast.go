package main

import (
	"context"
	"log"
	"sync"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	OfficialRPCs = []string{
		"https://eu.betarpc.io",
		"https://us.betarpc.io",
		"http://localhost:8545",
		"http://localhost:10545",
		"https://cloudflare-eth.com/",
	}
)

func broadcastViaRPC(ctx context.Context, officialConns []*ethclient.Client, officialRPCs []string, tx *types.Transaction) {
	wg := &sync.WaitGroup{}
	for i, c := range officialConns {
		wg.Add(1)
		if c == nil {
			continue
		}
		go func(i int, c *ethclient.Client) {
			log.Printf("send via %s", officialRPCs[i])
			err := c.SendTransaction(ctx, tx)
			if err != nil {
				log.Printf("Public Node %s err: %v", officialRPCs[i], err)
			}
			wg.Done()
		}(i, c)
	}
	wg.Wait()
}

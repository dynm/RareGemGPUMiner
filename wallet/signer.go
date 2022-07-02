package wallet

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func SignAndSend(ctx context.Context, conn *ethclient.Client, nonce uint64, auth *bind.TransactOpts, to common.Address, value *big.Int, calldata []byte, gasLimit int64, gasPrice *big.Int) (txHash common.Hash, err error) {
	if err != nil {
		log.Panic(err)
	}
	rawTx := types.NewTransaction(nonce, to, value, uint64(gasLimit), gasPrice, calldata)
	signedTx, err := auth.Signer(auth.From, rawTx)
	if err != nil {
		log.Printf("sign tx err: %v", err)
	}
	txHash = signedTx.Hash()
	err = conn.SendTransaction(ctx, signedTx)
	if err != nil {
		return
	}
	return signedTx.Hash(), nil
}

func Sign(nonce uint64, auth *bind.TransactOpts, to common.Address, value *big.Int, calldata []byte, gasLimit int64, gasPrice *big.Int) (signedTx *types.Transaction, err error) {
	if err != nil {
		log.Panic(err)
	}
	rawTx := types.NewTransaction(nonce, to, value, uint64(gasLimit), gasPrice, calldata)
	signedTx, err = auth.Signer(auth.From, rawTx)
	if err != nil {
		log.Printf("sign tx err: %v", err)
		return
	}

	return signedTx, nil
}

func SignEIP1559Tx(nonce uint64, auth *bind.TransactOpts, to common.Address, value *big.Int, calldata []byte, gasLimit uint64, gasFeeCap, gasTipCap *big.Int) (*types.Transaction, error) {
	baseTx := &types.DynamicFeeTx{
		Nonce:     nonce,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Gas:       gasLimit,
		Value:     value,
		Data:      calldata,
		To:        &to,
	}

	signedTx, err := auth.Signer(auth.From, types.NewTx(baseTx))
	if err != nil {
		log.Printf("sign tx err: %v", err)
		return nil, err
	}

	return signedTx, nil
}

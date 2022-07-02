package wallet

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

const (
	BIP44PATH = "m/44'/60'/0'/0"
)

type lastBeat struct {
	beatTime time.Time
	gasPrice *big.Int
	opts     *bind.TransactOpts
}
type Wallet struct {
	auths       []*bind.TransactOpts
	privateKeys []*ecdsa.PrivateKey
	authChan    chan *bind.TransactOpts
	authState   map[common.Hash]*lastBeat
}

func NewWallet(mnemonicPath string, accountnum uint32, chainID int64) *Wallet {
	f, err := os.Open(mnemonicPath)
	var mnemonic string
	if err != nil {
		for _, e := range os.Environ() {
			pair := strings.SplitN(e, "=", 2)
			if pair[0] == "MNEMONIC" {
				mnemonic = pair[1]
				break
			}
		}
		if mnemonic == "" {
			log.Fatalf("neither mnemonic.txt nor env %s found", "MNEMONIC")
		}
	} else {
		mnemonicBytes := make([]byte, 1024)
		len, err := f.Read(mnemonicBytes)
		if err != nil {
			log.Fatal(err)
		}
		mnemonic = string(mnemonicBytes[:len])
		f.Close()
	}

	wallet := &Wallet{
		auths:       make([]*bind.TransactOpts, accountnum),
		privateKeys: make([]*ecdsa.PrivateKey, accountnum),
		authChan:    make(chan *bind.TransactOpts),
		authState:   make(map[common.Hash]*lastBeat),
	}

	path := accounts.DefaultRootDerivationPath
	seed := bip39.NewSeed(mnemonic, "")
	key, _ := bip32.NewMasterKey(seed)
	var lastChildKey *bip32.Key
	lastChildKey = key
	for _, p := range path {
		lastChildKey, err = lastChildKey.NewChildKey(p)
		if err != nil {
			log.Fatal("BIP44 derivation", err)
		}
	}

	for i := uint32(0); i < accountnum; i++ {
		keyb, err := lastChildKey.NewChildKey(i)
		if err != nil {
			log.Fatal(err)
		}
		privateKey, err := crypto.ToECDSA(keyb.Key)
		if err != nil {
			log.Fatal(err)
		}
		auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainID))
		if err != nil {
			log.Panic(err)
		}
		wallet.privateKeys[i] = privateKey
		wallet.auths[i] = auth
		log.Printf("🔑Deriviated account: %s", auth.From.Hex())
	}
	go func() {
		i := uint32(0)
		for {
			wallet.authChan <- wallet.auths[i]
			i++
			if i >= accountnum {
				i = 0
			}
		}
	}()
	return wallet
}

func (w *Wallet) BeatTx(auth *bind.TransactOpts, tx *types.Transaction) {
	w.authState[tx.Hash()] = &lastBeat{
		beatTime: time.Now(),
		gasPrice: tx.GasPrice(),
		opts:     auth,
	}
}

func (w *Wallet) GetPrivateKey(i int) ecdsa.PrivateKey {
	return *w.privateKeys[i]
}

func (w *Wallet) NotifyAddGasPrice() {}

func (w *Wallet) GetNextAuth(ctx context.Context, conn *ethclient.Client, gasPrice *big.Int) (*bind.TransactOpts, *big.Int, error) {
	var auth *bind.TransactOpts
	if false {
		// if stuckAuth := w.getTooLongPendingAndLowGasPrice(ctx, conn, gasPrice); stuckAuth != nil {
		// auth = stuckAuth
		// golog.Infof("Use stucked auth, account=%s", auth.From.Hex())
	} else {
		auth = <-w.authChan
	}
	// nonce, err := conn.NonceAt(ctx, auth.From, nil)
	// if err != nil {
	// 	return nil, nil, fmt.Errorf("problem getting pending nonce: %+v", err)
	// }

	// gasPrice, err := conn.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	return nil, fmt.Errorf("problem getting gas price: %+v", err)
	// }

	// ethBalance, err := conn.BalanceAt(ctx, auth.From, nil)
	// if err != nil {
	// 	return nil, nil, fmt.Errorf("problem getting balance: %+v", err)
	// }

	// if ethBalance.Cmp(cost) < 0 {
	// 	return nil, nil, fmt.Errorf("insufficient ethereum to send a transaction: %v < %v", ethBalance, cost)
	// }

	// auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = nil // in wei
	// auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice
	return auth, nil, nil
}

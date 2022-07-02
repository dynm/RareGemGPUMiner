package miner

import (
	"crypto/rand"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

type CPU struct {
}

func (c *CPU) Mine(chainID *big.Int, Entropy [32]byte, Provably_Rare_Gem, From common.Address, gemKind, nonce, difficulty *big.Int, foundChan chan<- *big.Int) {

	// salt, hash, err = mine(chainID, gemInfo.Entropy, Provably_Rare_Gem, account.From, gemKind, nonce)
	/*
			    bytes32 entropy = 0x7465737400000000000000000000000000000000000000000000000000008761;
		        uint kind = 0x712d;
		        uint nonce = 0xffaa;
		        uint salt = 0xdeadbeef;
		        bytes memory data = abi.encodePacked(
		          block.chainid,
		          entropy,
		          address(this),
		          msg.sender,
		          kind,
		          nonce,
		          salt
		        );
	*/
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
	log.Printf("target:\t%s", common.BigToHash(target).Hex())
	saltChan := make(chan *big.Int, 100)
	shouldStop := false
	for i := 0; i < 48; i++ {
		go func() {
			for salt := range saltChan {
				newHeader := make([]byte, len(header))
				copy(newHeader, header)
				// log.Printf("header:\t%02x", header)
				salt.FillBytes(newHeader[len(newHeader)-32:])
				// log.Printf("newHeader:\t%02x", newHeader)
				hash := crypto.Keccak256Hash(newHeader)
				/*
						uint diff = gems[kind].difficulty;
					    require(val <= type(uint).max / diff, 'salt not good enough');
				*/

				if salt.Uint64()%10_000_000 == 0 {
					log.Printf("searching hash:\t%s, current salt: %d", hash.Hex(), salt)

				}

				if hash.Big().Cmp(target) <= 0 {
					log.Printf("found hash:\t%s, current salt: %d", hash.Hex(), salt)
					foundChan <- salt
					shouldStop = true
				}
			}
		}()
	}
	/*
		chainId				0000000000000000000000000000000000000000000000000000000000000001
		Entropy				1e2f8254a899a60fbbf20afdc7b3d18a62a8b5e2129a4ab1b049b840930ebe93
		Provably_Rare_Gem	8cf5ae0c7a5de4e22bec0cbdc7341636520d8c24
		From				f39fd6e51aad88f6f4ce6ab8827279cfffb92266
		gemKind				0000000000000000000000000000000000000000000000000000000000000007
		nonce				0000000000000000000000000000000000000000000000000000000000000000
		salt				0000000000000000000000000000000000000000000000000000000000000000
	*/
	/*
		0000000000000000000000000000000000000000000000000000000000000001
		2d4c49471172aa8f4c5a0b3d495145b1e78105d2b42fbfb56423359c5acd4ae5
		dcc555576498e638decbe15f8c61c73f1a76a3d9
		09d7b6c180aa1077665400194cc6f52950c1a884
		0000000000000000000000000000000000000000000000000000000000000007
		0000000000000000000000000000000000000000000000000000000000000000
		0000000000000000000000000000000000000000000000000000000000000000
	*/
	randomBytes := make([]byte, 32)
	rand.Read(randomBytes)
	start := new(big.Int).SetBytes(randomBytes)

	i := 0
	for {
		saltChan <- new(big.Int).Set(start)
		// log.Printf("salt: %d", start)
		start.Add(start, big.NewInt(1))

		i++
		if shouldStop {
			close(saltChan)
			break
		}
	}

}

package miner

import (
	"encoding/binary"
	"reflect"
	"unsafe"

	"github.com/ethereum/go-ethereum/crypto"
)

type MidState [25]uint64

func GetMidState(d crypto.KeccakState) []byte {
	v := reflect.ValueOf(d)
	midstatePtr := unsafe.Pointer(v.Elem().FieldByName("a").UnsafeAddr())
	midstate := *(*MidState)(midstatePtr)
	midStateBytes := make([]byte, 200)
	for i := range midstate {
		binary.LittleEndian.PutUint64(midStateBytes[i*8:i*8+8], midstate[i])
	}
	return midStateBytes
}

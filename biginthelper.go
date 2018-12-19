package utils

import (
	"io"

	"math/big"

	"fmt"

	"crypto/sha256"

	"github.com/ethereum/go-ethereum/common"
	"github.com/nkbai/log"
)

//EmptyHash all zero,invalid
var EmptyHash = common.Hash{}

//EmptyAddress all zero,invalid
var EmptyAddress = common.Address{}

//Pex short string stands for data
func Pex(data []byte) string {
	return common.Bytes2Hex(data[:4])
}

//ShaSecret is short for sha256
func Sha256(data ...[]byte) common.Hash {
	//	return crypto.Keccak256Hash(data...)
	var d = sha256.New()
	d.Reset()
	for _, b := range data {
		d.Write(b)
	}
	h := common.Hash{}
	h.SetBytes(d.Sum(nil))
	return h
}

//HPex pex for hash
func HPex(data common.Hash) string {
	return common.Bytes2Hex(data[:2])
}

//BPex bytes to string
func BPex(data []byte) string {
	return common.Bytes2Hex(data)
}

//APex pex for address
func APex(data common.Address) string {
	return common.Bytes2Hex(data[:4])
}

//APex2 shorter than APex
func APex2(data common.Address) string {
	return common.Bytes2Hex(data[:2])
}

//BigIntTo32Bytes convert a big int to bytes
func BigIntTo32Bytes(i *big.Int) []byte {
	data := i.Bytes()
	buf := make([]byte, 32)
	for i := 0; i < 32-len(data); i++ {
		buf[i] = 0
	}
	for i := 32 - len(data); i < 32; i++ {
		buf[i] = data[i-32+len(data)]
	}
	return buf
}

//ReadBigInt read big.Int from buffer
func ReadBigInt(reader io.Reader) *big.Int {
	bi := new(big.Int)
	tmpbuf := make([]byte, 32)
	_, err := reader.Read(tmpbuf)
	if err != nil {
		log.Error(fmt.Sprintf("read BigInt error %s", err))
	}
	bi.SetBytes(tmpbuf)
	return bi
}

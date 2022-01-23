package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"math"
	"math/big"
)

// hash difficulty or size
const difficulty = 18

type proofOfWork struct {
	Block  *Block
	Target *big.Int
}

func NewProof(b *Block) *proofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-difficulty))
	pow := &proofOfWork{b, target}

	return pow
}

func (pow *proofOfWork) initData(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.Data,
			ToHex(int64(nonce)),
			ToHex(int64(difficulty)),
		},
		[]byte{},
	)

	return data
}

func (pow *proofOfWork) Run() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.initData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r%x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}

	fmt.Println()

	return nonce, hash[:]
}

func (pow *proofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.initData(pow.Block.Nonce)
	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(pow.Target) == -1
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		panic(err)
	}

	return buff.Bytes()
}

package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

type blockChain struct {
	blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

func (b *Block) deriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func createBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.deriveHash()
	return block
}

func (chain *blockChain) addBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := createBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func genesis() *Block {
	return createBlock("Genesis", []byte{})
}

func main() {
	version := "Version: 1.0"
	fmt.Println("Go blockchain version: ", version)
}

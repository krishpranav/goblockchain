package main

import (
	"fmt"
)

// lets create  a new package and put all this functions there

func main() {
	version := "Version: 1.0"
	fmt.Println("Go blockchain version: ", version)

	chain := initChain()
	chain.addBlock("first block after genesis")
	chain.addBlock("second block")
	chain.addBlock("third block")

	for _, block := range chain.blocks {
		fmt.Printf("previous hash: %x\n", block.PrevHash)
		fmt.Printf("Data in block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
	}
}

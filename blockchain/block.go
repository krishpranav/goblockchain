package blockchain

type blockChain struct {
	blocks []*Block
}

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func createBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
}

func (chain *blockChain) addBlock(data string) {
	prevBlock := chain.blocks[len(chain.blocks)-1]
	new := createBlock(data, prevBlock.Hash)
	chain.blocks = append(chain.blocks, new)
}

func genesis() *Block {
	return createBlock("Genesis", []byte{})
}

func initChain() *blockChain {
	return &blockChain{[]*Block{genesis()}}
}

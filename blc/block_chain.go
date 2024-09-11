package blc

// BlockChain 区块链
type BlockChain struct {
	Blocks []*Block
}

func NewBlockChain() *BlockChain {
	genesisBlock := NewGenesisBlock([]byte("init block chain"))
	return &BlockChain{
		Blocks: []*Block{genesisBlock},
	}
}

func (bc *BlockChain) AddBlock(block *Block) {
	bc.Blocks = append(bc.Blocks, block)
}

func (bc *BlockChain) LastBlock() *Block {
	return bc.Blocks[len(bc.Blocks)-1]
}

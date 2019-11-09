package main


// 4.引入区块链
type BlockChain struct {
	blocks []*Block
}

// 定义一个区块链
func NewBlcokChain() *BlockChain {
	block := GensisBlock()
	return &BlockChain{ blocks: []*Block{block, }}
}

// 定义一个创世快
func GensisBlock() *Block {
	return NewBlock("这是一个创世快！", []byte{})
}

// 5. 添加区块
func (bc *BlockChain) AddBlock (data string) {
	// 前区快hash
	prevHash := bc.blocks[len(bc.blocks)-1].Hash
	// 添加到链中
	bc.blocks = append(bc.blocks, NewBlock(data, prevHash))
}

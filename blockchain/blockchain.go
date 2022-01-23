package blockchain

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	return &Blockchain{}
}

func (bc *Blockchain) Init() {
	b := NewBlock(nil, []byte("genesis block"))
	bc.AddBlock(b)
}

func (bc *Blockchain) GetBlock(p int) *Block {
	if p == -1 {
		return bc.Blocks[len(bc.Blocks)-1]
	}
	return bc.Blocks[p]
}

func (bc *Blockchain) GetLastBlock() *Block {
	return bc.GetBlock(-1)
}

func (bc *Blockchain) AddBlock(b *Block) {
	bc.Blocks = append(bc.Blocks, b)
}

func (bc *Blockchain) AddNewBlock(data []byte) {
	b := NewBlock(bc.GetLastBlock().Hash, data)
	bc.AddBlock(b)
}

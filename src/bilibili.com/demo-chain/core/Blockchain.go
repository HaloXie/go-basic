package core

import (
	"fmt"
	"log"
)

// Blockchain struct
type Blockchain struct {
	Blocks []*Block
}

func isValid(newBlock Block, oldBlock Block) bool {
	if newBlock.Index-1 != oldBlock.Index {
		return false
	}
	if newBlock.PrevBlockHash != oldBlock.Hash {
		return false
	}
	newHash := CalculateHash(newBlock)

	if newHash != newBlock.Hash {
		return false
	}
	return true
}

// Print => to print the all blocks
func (bc *Blockchain) Print() {
	for _, block := range bc.Blocks {
		fmt.Printf("Index:%d\n", block.Index)
		fmt.Printf("Prev.Hash:%s\n", block.PrevBlockHash)
		fmt.Printf("Curr.hash:%s\n", block.Hash)
		fmt.Printf("Data:%s\n", block.Data)
		fmt.Printf("Timespan:%d\n", block.Timestamp)
		fmt.Println()
	}
}

// NewBlockchain => create a new blockchian
func NewBlockchain() *Blockchain {
	genesisBlock := GenerateGenesisBlock()
	blocchain := Blockchain{}
	blocchain.AppendBlock(&genesisBlock)
	return &blocchain
}

// SendData => send the transcation data into the block
func (bc *Blockchain) SendData(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := GenerateNewBlock(*preBlock, data)
	bc.AppendBlock(&newBlock)
}

// AppendBlock => append a new block into the blockchain
func (bc *Blockchain) AppendBlock(newBlock *Block) {
	if len(bc.Blocks) == 0 {
		bc.Blocks = append(bc.Blocks, newBlock)
		return
	}
	if isValid(*newBlock, *bc.Blocks[len(bc.Blocks)-1]) {
		bc.Blocks = append(bc.Blocks, newBlock)
	} else {
		log.Fatal("invalid Block")
	}
}

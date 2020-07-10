package main

import (
	"core"
	"fmt"
)

func main() {

	bc := core.NewBlockChain()

	bc.AddBlock("Send 1 BTC 2 Ivan")
	bc.AddBlock("Send 2 BTC 2 Ivan")

	for _, block := range bc.Blocks {
		fmt.Printf("prev Hash:%v\n", block.PrevBlockHash)
		fmt.Printf("data:%s\n", block.Data)
		fmt.Printf("Hash:%s\n", block.Hash)
		fmt.Println()
	}
}

package core

import (
	"crypto/sha256"
	"encoding/hex"
	"time"
)

// Block struct
type Block struct {
	Index         int64  // 区块编号
	Timestamp     int64  // 时间戳
	PrevBlockHash string // 上一个区块的 Hash 值
	Hash          string // 当前区块的 Hash 值

	// 以下是区块体
	Data string // 区块数据
}

// CalculateHash => calc the block hash
func CalculateHash(b Block) string {
	blockData := string(b.Index) + string(b.Timestamp) + b.PrevBlockHash + b.Data
	hashInBytes := sha256.Sum256([]byte(blockData))
	return hex.EncodeToString(hashInBytes[:])
}

// GenerateNewBlock => create a new block
func GenerateNewBlock(preBlock Block, data string) Block {
	newBlock := Block{}
	newBlock.Index = preBlock.Index + 1
	newBlock.PrevBlockHash = preBlock.Hash
	newBlock.Timestamp = time.Now().Unix()
	newBlock.Data = data
	newBlock.Hash = CalculateHash(newBlock)

	return newBlock
}

// GenerateGenesisBlock => create a genesis block
func GenerateGenesisBlock() Block {
	preBlock := Block{}
	preBlock.Index = -1
	preBlock.Hash = ""
	return GenerateNewBlock(preBlock, "Genesis Block")
}

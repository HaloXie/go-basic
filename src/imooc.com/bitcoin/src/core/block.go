package core

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

// Block struct => keeps block header
type Block struct {
	Timestamp     int64  // 区块链创建时间戳
	Data          []byte // 区块包含的数据
	PrevBlockHash []byte // 前一个区块链的 Hash 值
	Hash          []byte // 当前区块链的 Hash 值
}

// NewGenesisBlock => create and return a genesis block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

// NewBlock => create a new block with hash
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{
		Timestamp:     time.Now().Unix(),
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
		Hash:          []byte{},
	}
	block.SetHash()

	return block
}

// SetHash => according to the block to generate the corresponding hash value
func (b *Block) SetHash() {
	// 时间戳转为字节数组
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

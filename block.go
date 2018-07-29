package main

import (
	"strconv"
	"bytes"
	"crypto/sha256"
	"time"
	"fmt"
)

type Block struct {
	Timestamp	int64
	Data 	[]byte
	PrevBlockHash []byte
	Hash  	[]byte
}
type Blockchain struct {
	blocks []*Block
}
func (b *Block)SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{timestamp, b.Data, b.PrevBlockHash}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

//初始化一个块
func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

//添加区块
func (bc *Blockchain)AddBlock(data string){
	prevBlock := bc.blocks[len(bc.blocks) - 1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block创世块", []byte{})
}

//使用创世块创建一个区块链函数
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

func main() {
	//初始化
	bc := NewBlockchain()
	//添加记录
	bc.AddBlock("Send 1 BTC to L")
	bc.AddBlock(("Send 1 BTC to R"))
	//查看记录
	for _, block := range bc.blocks {
		fmt.Printf("Prev : hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data : %s\n", block.Data)
		fmt.Printf("Hash : %x\n", block.Hash)
		fmt.Println()
	}
}
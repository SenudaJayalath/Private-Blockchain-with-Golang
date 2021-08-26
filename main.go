package main

import (
	"bytes"
	"crypto/sha256"
)

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
}

type BlockChain struct {
	block []*Block
}

func (b *Block) DeriveHash() {
	info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
	hash := sha256.Sum256(info)
	b.Hash = hash[:]
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash}
	block.DeriveHash()
	return block
}

func main() {
	CreateBlock("ffe", []byte{'1', '3', '5'})
	// var test Block
	// test.Data = []byte{'d', 'f', 'h'}
	// test.PrevHash = []byte{'1', '3', '5'}

	// test.DeriveHash()

}

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	Tranaction   string
	Nonce        int
	PreviousHash string
	BlockHash    string
}

type Chain struct {
	ChainBlock []Block
	ChainHash  string
}

func NewBlock(transaction string, Nonce int, PreviousHash string) *Block {
	s := new(Block)
	s.Tranaction = transaction
	s.Nonce = Nonce
	s.PreviousHash = PreviousHash
	s.CalculateHash(s.Tranaction + string(s.Nonce) + s.PreviousHash)
	return s
}

func (b *Block) CalculateHash(stringToHash string) {

	sum := sha256.Sum256([]byte(stringToHash))
	b.BlockHash = hex.EncodeToString(sum[:])

}

func (b *Chain) ListBlocks() {
	//A method to print all the blocks in a nice format showing block data such
	//as transaction, Nonce, previous hash, current block hash

	for i := range b.ChainBlock {
		fmt.Printf(" Tranaction is : %s \n", b.ChainBlock[i].Tranaction)
		fmt.Printf(" Nonce is : %d  \n ", (b.ChainBlock[i].Nonce))
		fmt.Println("\n Previous Hash is :    ", (b.ChainBlock[i].PreviousHash))
		// fmt.Printf("%s :\n ",b.ChainBlock[i].BlockHash)
		fmt.Println("\n  BlockHash is :  ", b.ChainBlock[i].BlockHash)

	}
}
func (b *Chain) ChangeBlock() {
	// function to change block transaction of the given block ref

	b.ChainBlock[3].Nonce = 732
}

func (b *Chain) VerifyChain() {

	b.ChainBlock[0].CalculateHash(b.ChainBlock[0].Tranaction + string(b.ChainBlock[0].Nonce) + b.ChainBlock[0].PreviousHash)

	for i := 1; i < len(b.ChainBlock); i++ {
		b.ChainBlock[i].CalculateHash(b.ChainBlock[i].Tranaction + string(b.ChainBlock[i].Nonce) + b.ChainBlock[i-1].BlockHash)
	}

	if b.ChainHash == b.ChainBlock[len(b.ChainBlock)-1].BlockHash {
		fmt.Printf("\n Block Chain is verified\n\n")
	} else {
		fmt.Printf("\n Block Chain is modofied\n\n")
	}

}

func main() {

	block1 := NewBlock("ALICE TO BOB", 2136, "0")
	block2 := NewBlock("BOB TO ALICE", 4234, block1.BlockHash)
	block3 := NewBlock("GOOFY TO SCROOGE", 3454, block2.BlockHash)
	block4 := NewBlock("SCROOGE TO GOOFY", 6654, block3.BlockHash)
	block5 := NewBlock("BITCOIN TO TESLA", 3354, block4.BlockHash)
	block_chain := new(Chain)
	block_chain.ChainBlock = append(block_chain.ChainBlock, *block1, *block2, *block3, *block4, *block5)
	block_chain.ChainHash = block5.BlockHash

	block_chain.VerifyChain()
	block_chain.ListBlocks()

	block_chain.ChangeBlock()
	block_chain.VerifyChain()
	block_chain.ListBlocks()

}

package main

import (
	"crypto/sha256"
	"fmt"
)

type block struct {
	hash        string
	transaction string
	prev        string
}

func NewBlock(transaction string, nonce int, previousHash string) *block {
	s := transaction + string(nonce) + previousHash
	b := block{
		hash:        s,
		transaction: transaction,
		prev:        previousHash,
	}
	return &b

}

func CalculateHash(stringToHash string) string {
	fmt.Printf("String received: %s\n", stringToHash)
	h := sha256.Sum256([]byte(stringToHash))

	return fmt.Sprintf("Hash: %x", h)
}

func main() {
	transaction1 := "Alice sends 1 coin to Bob"
	transcation2 := "Charlie sends 4 coins to Bob"

	ht1 := CalculateHash(transaction1)
	ht2 := CalculateHash(transcation2)

	fmt.Printf(ht1, ht2)
}

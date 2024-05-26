package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
	"errors"
)
var ErrInsufficientFunds =errors.New("Unsufficient funds!") 
const reward = 10 // The reward to the genesis adresses..  


// See https://en.bitcoin.it/wiki/Transaction You can replace
// ScriptSig and ScriptPubKey by adresses (just names)

// TXInput represents a transaction input see
type TXInput struct {
	Txid      []byte
	Vout      int
 	ScriptSig string
}

// TXOutput represents a transaction output
type TXOutput struct {
	Value        int
	ScriptPubKey string
}

// Transaction represents a Bitcoin transaction it contains a hash, as
// well as a set of input and output transactions
type Transaction struct {
	Hash   []byte
	TxIns  []TXInput
	TxOuts []TXOutput
}

func NewTransaction(hash []byte, txIns []TXInput, txOuts []TXOutput) *Transaction {
	transaction := &Transaction{
		Hash:   hash,
		TxIns:  txIns,
		TxOuts: txOuts,
	}
	return transaction
}

// Computes the Hash of a transaction
func (tx *Transaction) ComputeHash()[]byte{
	
	//variable to store the encoded data
	var encoded bytes.Buffer
	//gob package(Go Binary) encoding and decoding data in a binary format
	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	result := sha256.Sum256(encoded.Bytes())
	//Converts the array to a slice and returns the SHA-256 hash of the transaction
	return result[:]
}

// NewCoinbaseTX creates a new coinbase transaction
func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward for %s", to)
	}

	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{reward, to}
	tx := NewTransaction(nil, []TXInput{txin}, []TXOutput{txout})
	tx.Hash = tx.ComputeHash()
	return tx
}

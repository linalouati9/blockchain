package main

import (

)

// Blockchain implements interactions with a DB
type Blockchain struct {
	GHash []byte // Hash of the Genesis Block
	Chain []Block// Slice of blocks
}

// AddBlock adds a new block with the provided transactions, the block
// is mined before addition 
func (bc *Blockchain) AddBlock(transactions []*Transaction) {
	zeroBits := 16 
	prevBlock := []byte{}
	if len(bc.Chain) > 0 {
		prevBlock = bc.Chain[len(bc.Chain)-1].Hash
	}
	block := NewBlock(transactions, prevBlock, true, zeroBits)
	bc.Chain = append(bc.Chain, *block)
}

// NewSendTransaction creates a set of transactions to transfert the
// amount from sender to receiver. If the sender has not the required
// amount, an error is returned 
func (bc *Blockchain) NewTransfertTX(from, to string, amount int) (*Transaction, error) {
	var inputs []TXInput
	var outputs []TXOutput

	balanceSender := bc.GetBalance(from)
	if(balanceSender < amount){
		return nil, ErrInsufficientFunds
	}
	inputs = append(inputs, TXInput{[]byte{}, balanceSender, from})
	outputs = append(outputs, TXOutput{amount, to})
	if(balanceSender > amount){
		outputs = append(outputs, TXOutput{balanceSender - amount, from})
	}
	tx := NewTransaction(nil, inputs, outputs)
	return tx, nil
}

// CreateBlockchain creates a new blockchain, evey adress in adresses
// is given the initial 
func NewBlockchain(addresses []string) *Blockchain {
    zeroBits := 16
    var txs []*Transaction

    for _, address := range addresses {
		tx := NewCoinbaseTX(address, "")
        txs = append(txs, tx)
    }

    genesisBlock := NewGBlock(txs, zeroBits)
    bc := NewBlockchainFromGB(genesisBlock)
    return bc
}

// creates a new blockchain given a valid genesis block
func NewBlockchainFromGB(genesis *Block) *Blockchain {
	return &Blockchain{
		GHash: genesis.Hash,
		Chain: []Block{*genesis},
	}
	
}

func (bc *Blockchain) GetBalance(address string) int {
    balance := 0

    for _, block := range bc.Chain {
        for _, tx := range block.Transactions {
            for _, txout := range tx.TxOuts {
                if txout.ScriptPubKey == address {
                    balance += txout.Value
                }
            }

            for _, txin := range tx.TxIns {
                if txin.ScriptSig == address {
                    balance -= txin.Vout
                }
            }
        }
    }

    return balance
}


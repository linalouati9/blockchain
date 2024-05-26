package main

import ("crypto/sha256"
	"math/big"
	"time"
)
// Block keeps block headers
type Block struct {
	Hash          []byte
	PrevBlockHash []byte
	Transactions  []*Transaction
	Timestamp     int64
	Nonce         int
}

// HashTransactions returns a hash of the transactions in the block
func (b *Block) HashTXs() []byte {
	var txHashes [][]byte
	for _, tx := range b.Transactions {
		txHashes = append(txHashes, tx.Hash)
	}
	result:= sha256.Sum256(Serialize(txHashes))
	
	return result[:]
}

// NewBlock creates and returns Block, for testing purposes, mining
// can be activated or disabled using the boolean flag mine If mine is
// set to true, the nonce is computed so that the hash start with
// zeroBits
func NewBlock(txs []*Transaction, prevBHash []byte, mine bool, zeroBits int) *Block {
	block := &Block{nil,  prevBHash, txs, time.Now().Unix() , 0}
	if mine{
		block.Mine(zeroBits)
	}else{
		block.Nonce = 0
	}
	return block
}

//the first block in a blockchain
// Creates and returns genesis Block, its hash must start with zeroBits
func NewGBlock(cbtx []*Transaction, zeroBits int) *Block {
	genesisBlock := &Block{nil, []byte{}, cbtx, time.Now().Unix(), 0}
	genesisBlock.Mine(zeroBits)
	return genesisBlock
}


// true if the block is correclty Hashed 
func (block *Block) IsCorrectlyHashed(zeroBits int) bool {
	hash := block.HashTXs()
	hash = append(hash, block.PrevBlockHash...)
	hash = append(hash, IntToHex(block.Timestamp)...)
	hash = append(hash, IntToHex(int64(block.Nonce))...)

	hashInBytes := sha256.Sum256(hash)
	hashInt := new(big.Int).SetBytes(hashInBytes[:])
	target := new(big.Int).Exp(big.NewInt(2), big.NewInt(int64(256-zeroBits)), nil)

	return hashInt.Cmp(target) == -1
}

// Hashes a block, private function 
// private function because it starts with a lowercase  
func (block *Block) computeHash() []byte {
	contents := Serialize([][]byte{
		block.PrevBlockHash,
		block.HashTXs(),
		IntToHex(block.Timestamp),
		IntToHex(int64(block.Nonce)),
	})
	result:=sha256.Sum256(contents)
	return result[:]
}		

// Computes and sets the hash of "block"
func (block *Block) SetHash(){
	block.Hash = block.computeHash()
}		

// Mines a block : iterates over nonces until the hash starts with the
// number of zeros defined by zeroBits		
func (block *Block) Mine(zeroBits int) {
	for {
		block.Nonce++
		hash := block.computeHash()

		if block.IsCorrectlyHashed(zeroBits) {
			block.Hash = hash
			break
		}
	}
}
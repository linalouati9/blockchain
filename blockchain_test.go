// for more information about testing in Go, see
// https://ieftimov.com/posts/testing-in-go-go-test/
package main

import (
	"os"
	"errors"
	"testing";
	"fmt";
	//"strconv"
)
var note float64 =.0

// tests if a []byte starts with a given number of zeros 
func TestStartsWithXZeros(t *testing.T){
	var test_note=1.5
	b0 := []byte{0x0F}
	if !StartsWithXZeros(b0,0){
		t.Error("isValidHash has Problem with b0")
		test_note=test_note-.25
	}
	if !StartsWithXZeros(b0,1){
		t.Error("isValidHash has Problem with b0")
		test_note=test_note-.25
	}
	if StartsWithXZeros(b0,7){
		t.Error("isValidHash has Problem with b0")
		test_note=test_note-.25
	}
	
	b1 := []byte{0x01, 0x00, 0x01}
	if StartsWithXZeros(b1,8){
		t.Error("isValidHash has Problem with b1")
		test_note=test_note-.25
	}

	b2 := []byte{0x00, 0x00, 0x01}
	if StartsWithXZeros(b2,24){
		t.Error("isValidHash has Problem with b2")
		test_note=test_note-.25
	}
	if !StartsWithXZeros(b2,23){
		t.Error("isValidHash has Problem with b2")
		test_note=test_note-.25
	}	
	note=note+test_note
}

func TestEqualSlices(t *testing.T){
	var test_note=0.75
	s := []byte{0x1,0x1,0x2,0x1,0x1,0x1}
	if !EqualSlices(s[0:2],s[3:5]){
		t.Error("TestEqualSlices failed")
		test_note=test_note-0.25
	}	
	
	if EqualSlices(s[0:2],s[3:6]){
		t.Error("TestEqualSlices failed")
		test_note=test_note-0.25
	}	

	if EqualSlices(s[1:3],s[3:5]){
		t.Error("TestEqualSlices failed")
		test_note=test_note-0.25
	}
	note=note+test_note
}

func TestIntToHex(t *testing.T){
	var test_note=0.75
	if !EqualSlices(IntToHex(0), []byte{0,0,0,0,0,0,0,0}){
		t.Error("TestIntToHex failed")
		test_note=test_note-0.25

	}	
	
	var p int64= 16777215  
	want:= []byte{0,0,0,0,0,0xFF,0xFF,0xFF}
	
	if !EqualSlices(IntToHex(p),want){
		t.Error("TestIntToHex failed")
		test_note=test_note-0.5

	}
	note=note+test_note
}

// Coinbase transacations are the genesis transactions, they don't
// need TXInputs, contrary to other transactions. Since mining is free
// in this Blockchain, and there is no inflation, the number of
// bitcoins in the blockchain remains constant.

func TestNewCoinbaseTX(t *testing.T){
	var test_note=3.0
	have1:=*NewCoinbaseTX("ammar","Reward for ammar 404")
	have2:=*NewCoinbaseTX("ammar","")
	
	txin1  := TXInput{[]byte{}, -1, "Reward for ammar 404"}
	txin2 := TXInput{[]byte{}, -1, "Reward for ammar"}
	txout := TXOutput{10, "ammar"}
	want1 := Transaction{nil, []TXInput{txin1}, []TXOutput{txout}}
	want2 := Transaction{nil, []TXInput{txin2}, []TXOutput{txout}}
	
	if EqualTransactions(have1,want1){
		t.Error("TestNewCoinbaseTX failed")
		test_note=test_note-1
		
	}
	want1.Hash=want1.ComputeHash()
	if !EqualTransactions(have1,want1){
		t.Error("TestNewCoinbaseTX failed")
		test_note=test_note-1
	}
	
	want2.Hash=want2.ComputeHash()

	if !EqualTransactions(have2,want2){
		t.Error("TestNewCoinbaseTX failed")
		test_note=test_note-1

	}
	note=note+test_note
}


// creates a block with no transactions, nonce and timestamps are set
// to zero
func TestCreateEmptyBlock(t *testing.T){
	var test_note=1.0
	want := Block{nil, []byte{0x0}, nil, 0, 0}
	
	got:=NewBlock(nil, []byte{0x0},false, 0)
	got.Timestamp=0
	if !EqualBlocks(*got,want){
		fmt.Println(want)
		fmt.Println(got)
		t.Error("TestCreateEmptyBlock failed")
		test_note=test_note-1

	}
	note=note+test_note
}

// Creates a Genesis Block and a blockchain that contains . The
// Genesis Block is the only entry point for bitcoins
func TestCreateGBlock(t *testing.T){
	var test_note=3.0
	difficulty:=16
	cbtx:=NewCoinbaseTX("ammar","Reward for ammar 404")
	genesisblock:=NewGBlock([]*Transaction{cbtx},difficulty);
	bc:=NewBlockchainFromGB(genesisblock)
	if !genesisblock.IsCorrectlyHashed(difficulty)||
		bc.GetBalance("ammar")!=10{
			t.Error("TestCreateGBlock failed")
			test_note=test_note-3
	
		}
	note=note+test_note
}


func TestCreateGBlock2(t *testing.T){
	var test_note=1.0
	difficulty:=18
	cbtx1:=NewCoinbaseTX("ammar","Reward for ammar 404")
	cbtx2:=NewCoinbaseTX("ammar","Reward for ammar 404")
	genesisblock:=NewGBlock([]*Transaction{cbtx1,cbtx2},difficulty);
	bc:=NewBlockchainFromGB(genesisblock)
	
	if !genesisblock.IsCorrectlyHashed(difficulty)||
		bc.GetBalance("ammar")!=20 {
		t.Error("TestCreateGBlock failed")
		test_note=0
	}
	fmt.Println(bc)

	note=note+test_note
}

func TestSerialization(t *testing.T){
	var test_note=2.0

	difficulty:=16
	cbtx1:=NewCoinbaseTX("ammar","Reward for ammar 404")
	cbtx2:=NewCoinbaseTX("ammar","Reward for ammar 404")
	cbtx3:=NewCoinbaseTX("ali","Reward for ali 404")
	
	genesisblock:=NewGBlock([]*Transaction{cbtx1,cbtx2,cbtx3},difficulty);
	bc:=NewBlockchainFromGB(genesisblock)

	err:=SaveBlockchain(bc,"bctest.json")
	if err!=nil {
		t.Error("SaveBlockchain Failed")
	}
	
	if _, err := os.Stat("bctest.json");errors.Is(err, os.ErrNotExist)  {
		t.Error("File was not written in the disk")
		test_note=test_note-1
	}
		
	bc2,err:=LoadBlockchain("bctest.json")	
	
	if err!=nil||bc2.GetBalance("ammar")!=20 ||bc2.GetBalance("ali")!=10{
		t.Error("TestSerialization failed")
		test_note=test_note-1
	}
	err = os.Remove("bctest.json")
	if err !=nil  {
		t.Error("Can't remove file from disk")
	}
	note=note+test_note
}

func TestSuccessfulTransfert(t *testing.T){
	var test_note=2.0
	bc:= NewBlockchain([]string{"ammar","ali"})
	tx,err:=bc.NewTransfertTX("ali","ammar", 5)
	bc.AddBlock([]*Transaction{tx})

	if err !=nil || bc.GetBalance("ammar")!=15 || bc.GetBalance("ali")!=5{
		t.Error("TestSuccTransfert failed")
		test_note=test_note-2

	}
	note=note+test_note
}

func TestImpossibleTransfert(t *testing.T){
	var test_note=2.0
	bc:= NewBlockchain([]string{"ammar","ali"})
	tx,err:=bc.NewTransfertTX("ali","ammar", 15)
	if tx!=nil {
		bc.AddBlock([]*Transaction{tx})
	}
	
	if !errors.Is(err,ErrInsufficientFunds) || bc.GetBalance("ammar")!=10 || bc.GetBalance("ali")!=10{
		t.Error("TestSuccTransfert failed")
		test_note=test_note-2

	}
	note=note+test_note	
}

func TestMain(m *testing.M) {
	m.Run()
	fmt.Println("\n===================================")
	fmt.Println("========   Résumé   ===============")
	fmt.Printf("  Note actuelle : %2.2f/17\n",note)
	fmt.Println("===================================")	
}

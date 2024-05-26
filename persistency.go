package main

import (
	"os"
	"errors"
	"encoding/json" //to save/load the blockchain using json encoding  
    //"io/ioutil" for reading and writing files
)

var ErrInexistantBC =errors.New("No existing Blockchain found.  Create one first.")

// true if the file is existent
func bcFileExists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

// LoadBlockchain loads a blockchain from a file
func LoadBlockchain(file string) (*Blockchain, error) {
	if !bcFileExists(file) {
		return nil, ErrInexistantBC
	}

	// Open the file
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close() //the file is closed after its content has been read

	// Decode JSON data into a Blockchain struct
	var blockchain Blockchain
	decoder := json.NewDecoder(f)
	if err := decoder.Decode(&blockchain); err != nil {
		return nil, err
	}

	return &blockchain, nil
}

func SaveBlockchain(bc *Blockchain, file string) error {
	// Open the file for writing
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	// Encode the Blockchain struct to JSON
	encoder := json.NewEncoder(f)
	if err := encoder.Encode(bc); err != nil {
		return err
	}

	return nil
}

package main

import (
	"encoding/binary"
)

// IntToHex converts an int64 to a byte array of length 8
// 
func IntToHex(num int64) []byte {
    bytes := make([]byte, 8)
    binary.BigEndian.PutUint64(bytes, uint64(num))
	return bytes
}

// true if the hash starts with zeroBits zeros, note that the hash is
// a slice of *bytes* but we want zeroBits *bits* (a byte has 8 bits)
func StartsWithXZeros(hash []byte, zeroBits int) bool {
	// Calculate the number of complete bytes needed for zeroBits
	completeBytes := zeroBits / 8

	// Iterate over the complete bytes
	for i := 0; i < completeBytes; i++ {
		if hash[i] != 0 {
			return false
		}
	}

	// Check remaining bits
	remainingBits := zeroBits % 8
	if remainingBits > 0 && hash[completeBytes]>>(8-remainingBits) != 0 {
		return false
	}

	return true
}
 
func EqualSlices(a, b []byte) bool {
	// Check if the slices have the same length
	if len(a) != len(b) {
		return false
	}

	// Compare each byte in the slices
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	// If all bytes match, the slices are equal
	return true
}

func EqualMaps(a, b map[string]int) bool {
	// Check if the maps have the same number of entries
    if len(a) != len(b) {
		return false
	}

	// Iterate over the keys in map 'a'
	for key, valueA := range a {
		// Check if the key is present in map 'b'
		valueB, ok := b[key]
		if !ok {
			// Key is missing in map 'b'
			return false
		}

		// Check if the values are equal
		if valueA != valueB {
			return false
		}
	}

	// All keys and values match, maps are equal
	return true
}

func EqualTransactions(a, b Transaction) bool {
    return EqualSlices(a.Hash, b.Hash)
}

func EqualBlocks(a,b Block) bool{
	return EqualSlices(a.Hash,b.Hash)
}

// Serializes a slice of byte slices by converting it to a byte slice so
// needed to easily hash data
func Serialize(input [][]byte) []byte {
	var result []byte

	for _, slice := range input {
		result = append(result, slice...)
	}

	return result
}
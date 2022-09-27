package hash

import (
	"strconv"
	"testing"
)

func TestHashing(t *testing.T) {
	hash := New(3, func(key []byte) uint32 {
		i, _ := strconv.Atoi(string(key))
		return uint32(i)
	})

	//Given the above hash function, this will give replicas with "hashes":
	//2,4,6,12,14,16,22,24,26
	hash.Add("6", "4", "2")
}

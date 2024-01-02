package randutils

import (
	"crypto/rand"
	"encoding/hex"
	"math/big"
)

// RandStr generates a random string double of length len unless len is less than 8
func RandStr(len int) string {
	if len < 8 {
		len = 8
	}
	b := make([]byte, len)
	if _, err := rand.Read(b); err != nil {
		// rand.Read should never error unless we run out of entropy
		panic(err)
	}

	return hex.EncodeToString(b)
}

// RandInt64 generates a random int64
func RandInt64() int64 {
	i, err := rand.Int(rand.Reader, big.NewInt(1000000))
	if err != nil {
		panic(err)
	}
	return i.Int64()
}

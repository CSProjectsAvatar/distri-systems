package utils

import (
	"crypto/sha256"
	"fmt"
)

func Hash(input string) string {
	h := sha256.New()

	h.Write([]byte(input))
	res := h.Sum(nil)

	// data := binary.BigEndian.Uint64(res)
	// convert res to string
	data := fmt.Sprintf("%x", res)
	return data
}

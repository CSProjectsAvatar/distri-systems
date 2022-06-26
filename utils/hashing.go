package utils

import (
	"crypto/sha256"
	"encoding/binary"
)

func Hash(input string) uint64 {
	h := sha256.New()

	h.Write([]byte(input))
	res := h.Sum(nil)

	data := binary.BigEndian.Uint64(res)
	return data
}

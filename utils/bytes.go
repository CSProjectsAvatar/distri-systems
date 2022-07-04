package utils

import (
	"bytes"
	"fmt"
	"hash"
)

// Checks if x is in (a, b] interval. The interval can be circular, i.e,
// b < a, which is equivalent to (a, MAX] + [0, b].
func InIntervalRIncluded(x, a, b []byte) bool {
	return InInterval(x, a, b) || bytes.Equal(x, b)
}

// Checks if x is in (a, b) interval. The interval can be circular, i.e,
// b < a, which is equivalent to (a, MAX] + [0, b).
func InInterval(x, a, b []byte) bool {
	aLtx := bytes.Compare(a, x) < 0
	xLtb := bytes.Compare(x, b) < 0

	switch bytes.Compare(a, b) {
	case 1:
		return aLtx || xLtb // a < x || x < b
	case -1:
		return aLtx && xLtb // x in (a, b)
	default:
		return bytes.Compare(a, x) != 0 // (a, a) means Universe but a, so x belongs to (a, a) <=> x != a
	}
}

// Sha1Sized returns the first size bits of applying a hash obtained from the given factory
// to the given string.
func Sha1Sized(s string, hFactory func() hash.Hash, size uint) []byte {
	h := hFactory()

	if size > uint(h.Size())*8 || size%8 != 0 {
		panic(fmt.Sprintf("size must be a multiple of 8 less than or equal to %d (the hash size)", h.Size()))
	}

	if _, err := h.Write([]byte(s)); err != nil {
		panic(err)
	}
	return h.Sum(nil)[:size/8]
}

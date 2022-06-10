package utils

import "bytes"

// Checks if x is in (a, b] interval. The interval can be circular, i.e,
// b < a, which is equivalent to (a, MAX] + [0, b].
func InIntervalRIncluded(x, a, b []byte) bool {
	return InInterval(x, a, b) || bytes.Equal(x, b)
}

// Checks if x is in (a, b) interval. The interval can be circular, i.e,
// b < a, which is equivalent to (a, MAX] + [0, b).
func InInterval(x, a, b []byte) bool { // @audit test it
	aLtx := bytes.Compare(a, x) < 0
	xLtb := bytes.Compare(x, b) < 0

	switch bytes.Compare(a, b) {
	case 1:
		return aLtx || xLtb // a < x || x < b
	case -1:
		return aLtx && xLtb // x in (a, b)
	default:
		return false
	}
}

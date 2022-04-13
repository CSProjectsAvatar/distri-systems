package utils

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestUniqueInt(t *testing.T) {
	input := []int{1, -1, 1, 3, 2, 2, 7, 6, 3}
	expect := []int{1, -1, 3, 2, 7, 6}
	ans := Unique(input)
	lenAns, lenExp := len(ans), len(expect)

	if lenAns != lenExp {
		t.Errorf("Expected length %v, got %v", lenExp, lenAns)
	}

	for _, x := range expect {
		if !slices.Contains(ans, x) {
			t.Errorf("Element %v not found in %v", x, ans)
		}
	}
}

func TestUniqueStr(t *testing.T) {
	input := []string{"andy", "el tipazo", "crudo", "crudo", "andy"}
	expect := []string{"andy", "el tipazo", "crudo"}
	ans := Unique(input)
	lenAns, lenExp := len(ans), len(expect)

	if lenAns != lenExp {
		t.Errorf("Expected length %v, got %v", lenExp, lenAns)
	}

	for _, x := range expect {
		if !slices.Contains(ans, x) {
			t.Errorf("Element %v not found in %v", x, ans)
		}
	}
}

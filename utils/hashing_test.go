package utils

import "testing"

func TestHash(t *testing.T) {
	input := "andy"
	input2 := "nandy"
	input3 := "andy"

	ans := Hash(input)
	ans2 := Hash(input2)
	ans3 := Hash(input3)

	if ans == ans2 {
		t.Errorf("Expected different hashes, got %v", ans)
	}
	if ans != ans3 {
		t.Errorf("Expected same hashes, got %v", ans)
	}
}

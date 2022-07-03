package utils

import "fmt"

type Tuple[T comparable, U any] struct {
	First  T
	Second U
}

// JoinToTuples make a slice of tuples with the players and their games won
func JoinToTuples[T comparable, U any](slice1 []T, slice2 []U) []Tuple[T, U] {
	if len(slice1) != len(slice2) {
		panic(fmt.Sprintf("Slices must be of the same length: %d != %d", len(slice1), len(slice2)))
	}
	tuples := make([]Tuple[T, U], len(slice1))
	for i, value := range slice1 {
		tuples[i] = Tuple[T, U]{value, slice2[i]}
	}
	return tuples
}

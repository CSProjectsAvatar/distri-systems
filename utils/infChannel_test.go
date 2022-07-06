package utils

import (
	"log"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInfChannel(t *testing.T) {
	require := require.New(t)
	limit := 1000
	in, out := MakeInf[int]()
	lastVal := -1
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for val := range out {
			require.Equal(lastVal+1, val)
			lastVal = val
		}
		wg.Done()
		log.Println("Finished Reading")
	}()

	for i := 0; i < limit; i++ {
		in <- i
	}
	close(in)
	log.Println("Finished Writing")
	wg.Wait()

	require.Equal(limit-1, lastVal)
}

func TestConsume(t *testing.T) {
	require := require.New(t)
	toConsume := make(chan int, 5)
	toConsume <- 1
	toConsume <- 2
	toConsume <- 3
	toConsume <- 4
	toConsume <- 5

	require.Equal(5, len(toConsume))
	Consume[int](toConsume)

	require.Equal(0, len(toConsume))
}

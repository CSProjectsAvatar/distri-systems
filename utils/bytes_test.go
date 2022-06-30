package utils

import (
	"encoding/hex"
	"fmt"
)

func ExampleInInterval() {
	x, err := hex.DecodeString("d118e5a3cc15b182d1286373a60c787e58d3166f")
	if err != nil {
		panic(err)
	}
	a, err := hex.DecodeString("a458e9a3e1c30c664054ce477d28b14dc26def85")
	if err != nil {
		panic(err)
	}
	b, err := hex.DecodeString("bb3807ea1a2e7962e147c3b2ce0a08577286944b")
	if err != nil {
		panic(err)
	}
	fmt.Println(InInterval(x, a, b))

	// Output:
	// false
}

package chord

import "math/big"

// Finger Table.
type Ftable []*FingerRow

// A row of the Finger Table.
type FingerRow struct {
	Id   []byte      // hash of (n + 2^i) mod 2^m (i>=0)
	Node *RemoteNode // first node on circle that succeeds Id
}

// Computes (n + 2^i) mod 2^m
func FingerId(n []byte, i, m int) []byte {
	id := new(big.Int).SetBytes(n) // bigint id

	_2 := big.NewInt(2)
	_2pow_i := big.Int{}
	_2pow_i.Exp(_2, big.NewInt(int64(i)), nil) // 2^i

	sum := big.Int{}
	sum.Add(id, &_2pow_i) // n + 2^i

	_2pow_m := big.Int{}
	_2pow_m.Exp(_2, big.NewInt(int64(m)), nil) // 2^m

	id.Mod(&sum, &_2pow_m) // (n + 2^i) mod 2^m

	return id.Bytes()
}

func NewFingerRow(id []byte, node *RemoteNode) *FingerRow {
	return &FingerRow{
		Id:   id,
		Node: node,
	}
}

package dht

import "math/big"

// Creates a finger table of m rows when each successor is the given node.
func newFtable(node *RemoteNode, m int) ftable {
	table := make([]*fingerRow, m)
	for i := range table {
		// node is always the successor since it's the only node in the ring (maybe until ensureFtable() is called)
		table[i] = newFingerRow(fingerId(node.Id, i, m), node)
	}
	return table
}

func newFingerRow(id []byte, node *RemoteNode) *fingerRow {
	return &fingerRow{
		Id:   id,
		Node: node,
	}
}

// Computes (n + 2^i) mod 2^m
func fingerId(n []byte, i, m int) []byte {
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

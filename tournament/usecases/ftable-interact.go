package usecases

import (
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain/chord"
)

// Creates a finger table of m rows when each successor is the given node.
func newFtable(node *chord.RemoteNode, m int) chord.Ftable {
	table := make([]*chord.FingerRow, m)
	for i := range table {
		// node is always the successor since it's the only node in the ring (maybe until ensureFtable() is called)
		table[i] = chord.NewFingerRow(chord.FingerId(node.Id, i, m), node)
	}
	return table
}

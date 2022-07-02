package usecases

import (
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain/chord"
)

// Creates a finger table of m rows when each successor is the given node.
func newFtable(node *chord.RemoteNode, m uint) chord.Ftable {
	return make([]*chord.FingerRow, m)
}

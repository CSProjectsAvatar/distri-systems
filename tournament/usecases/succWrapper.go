package usecases

import (
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain/chord"
	"strconv"
)

type SuccWrapper struct {
	node *chord.Node
}

func NewSuccWrapper(node *chord.Node) *SuccWrapper {
	return &SuccWrapper{node: node}
}
func (sw *SuccWrapper) GetSuccessor() string {
	remote := sw.node.GetSuccessor()
	return remote.Ip + ":" + strconv.Itoa(int(remote.Port))
}

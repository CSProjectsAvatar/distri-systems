package tests

import (
	"log"
	"testing"

	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/tournament/usecases"
	"github.com/stretchr/testify/require"
)

func TestCreateFirstDefeat(t *testing.T) {
	log.Println("> Create AllvsAll Tree <")
	tm := usecases.NewMockTourMngr()
	tm.TInfo.Type_ = domain.First_Defeat
	tree := tm.Tree()

	require := require.New(t)
	require.Equal(2, len(tree.Children))
}
func TestCreateAllvsAllTree(t *testing.T) {
	log.Println("> Create AllvsAll Tree <")
	tm := usecases.NewMockTourMngr()
	tm.TInfo.Type_ = domain.All_vs_All
	tree := tm.Tree()

	require := require.New(t)
	require.Equal(4, len(tree.Children))
}

func TestCreateGroupsTree(t *testing.T) {
	log.Println("> Create Groups Tree <")
	tm := usecases.NewMockTourMngr()
	// Double Players
	tm.TInfo.Players = append(tm.TInfo.Players, tm.TInfo.Players...)
	tm.TInfo.Type_ = domain.Groups
	tree := tm.Tree()

	require := require.New(t)
	require.Equal(2, len(tree.Children))
	require.Equal(4, len(tree.Children[0].Children))
}

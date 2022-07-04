package tests

import (
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain/chord"
	"github.com/CSProjectsAvatar/distri-systems/tournament/infrastruct"
	"github.com/CSProjectsAvatar/distri-systems/tournament/interfaces"
	"github.com/CSProjectsAvatar/distri-systems/tournament/usecases"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestDhtTourMngr(t *testing.T) {
	os.Remove("logrus.log")
	removeDbs()

	nodes := ringOf(5, t)
	dhtStr := infrastruct.NewTestDht[string](5, nodes[0].RemoteNode)
	dhtInfos := infrastruct.NewTestDht[[]*domain.TournInfo](5, nodes[0].RemoteNode)
	dhtMatches := infrastruct.NewTestDht[[]*domain.Pairing](5, nodes[0].RemoteNode)

	mngr := &usecases.DhtTourDataMngr{
		DhtStr:     dhtStr,
		DhtInfos:   dhtInfos,
		DhtMatches: dhtMatches,
	}
	t.Run("infos", SubTestTournInfos(mngr))
	t.Run("no unfinished tournaments", SubTestNonUnfinished(mngr))
	t.Run("matches", SubTestMatches(mngr))
	t.Run("files", SubTestFiles(mngr))
	t.Run("creator", SubTestTournCreator(mngr))

	for _, n := range nodes {
		require.Nil(t, n.Stop())
	}
}

func SubTestTournCreator(mngr *usecases.DhtTourDataMngr) func(t *testing.T) {
	return func(t *testing.T) {
		f1 := &interfaces.TournFile{
			Name:   "player1.py",
			Data:   []byte(`print("habla matador")`),
			IsGame: false,
		}
		f2 := &interfaces.TournFile{
			Name:   "player2.py",
			Data:   []byte(`print("escucha matador")`),
			IsGame: false,
		}
		f3 := &interfaces.TournFile{
			Name:   "chess.py",
			Data:   []byte(`print("mira matador")`),
			IsGame: true,
		}
		info, err := interfaces.SaveTournament("all-vs-all", domain.All_vs_All, []*interfaces.TournFile{f1, f2, f3}, mngr)
		require.Nil(t, err)

		assert.Equal(t, "chess.py", info.Name)

		f1Content, err := mngr.File(info.ID, "player1.py")
		require.Nil(t, err)
		assert.Equal(t, `print("habla matador")`, f1Content)

		f2Content, err := mngr.File(info.ID, "player2.py")
		require.Nil(t, err)
		assert.Equal(t, `print("escucha matador")`, f2Content)

		f3Content, err := mngr.File(info.ID, "chess.py")
		require.Nil(t, err)
		assert.Equal(t, `print("mira matador")`, f3Content)

		savedInfo, err := mngr.GetTournInfo(info.ID)
		require.Nil(t, err)
		assert.Equal(t, *info, *savedInfo)
	}
}

func SubTestFiles(mngr *usecases.DhtTourDataMngr) func(t *testing.T) {
	return func(t *testing.T) {
		f1 := "habla matador"
		f2 := "fuego con tol mun2"
		f3 := "eso depende"

		require.Nil(t, mngr.SaveFiles("t1", &map[string]string{
			"f1": f1,
			"f2": f2,
			"f3": f3,
		}))

		v1, err := mngr.File("t1", "f1")
		require.Nil(t, err)
		assert.Equal(t, f1, v1)

		v2, err := mngr.File("t1", "f2")
		require.Nil(t, err)
		assert.Equal(t, f2, v2)

		v3, err := mngr.File("t1", "f3")
		require.Nil(t, err)
		assert.Equal(t, f3, v3)
	}
}

func SubTestMatches(mngr *usecases.DhtTourDataMngr) func(t *testing.T) {
	return func(t *testing.T) {
		p1 := &domain.Player{"andy"}
		p2 := &domain.Player{"omar"}
		//p3 := &domain.Player{"aylin"}
		p4 := &domain.Player{"celedonio"}

		m1 := &domain.Pairing{
			Winner:  domain.Draw,
			ID:      "m1",
			TourId:  "tour-1",
			Player1: p1,
			Player2: p4,
		}

		m2 := &domain.Pairing{
			Winner:  domain.Draw,
			ID:      "m2",
			TourId:  "tour-1",
			Player1: p1,
			Player2: p2,
		}

		m3 := &domain.Pairing{
			Winner:  domain.Draw,
			ID:      "m3",
			TourId:  "tour-1",
			Player1: p4,
			Player2: p2,
		}
		require.Nil(t, mngr.SaveMatch(m1))
		require.Nil(t, mngr.SaveMatch(m2))
		require.Nil(t, mngr.SaveMatch(m3))

		ms, err := mngr.Matches("tour-1")
		require.Nil(t, err)
		assert.Equal(t, []*domain.Pairing{m1, m2, m3}, ms)
	}
}

func SubTestTournInfos(mngr *usecases.DhtTourDataMngr) func(t *testing.T) {
	return func(t *testing.T) {
		inf1 := &domain.TournInfo{
			ID:      "tour-1",
			Name:    "tour-1",
			Players: []*domain.Player{{Id: "manolo"}, {"el loco"}, {"de la mata"}},
		}
		inf2 := &domain.TournInfo{
			ID:      "tour-2",
			Name:    "tour-2",
			Players: []*domain.Player{{Id: "manolo"}, {"el loco"}, {"de la mata"}},
		}
		inf3 := &domain.TournInfo{
			ID:      "tour-3",
			Name:    "tour-3",
			Players: []*domain.Player{{Id: "manolo"}, {"el loco"}, {"de la mata"}},
			Winner:  &domain.Player{Id: "manolo"},
		}
		require.Nil(t, mngr.SetTournInfo(inf3))
		require.Nil(t, mngr.SetTournInfo(inf1))
		require.Nil(t, mngr.SetTournInfo(inf2))

		inf, err := mngr.GetTournInfo("tour-1")
		require.Nil(t, err)
		require.Equal(t, *inf1, *inf)

		inf, err = mngr.GetTournInfo("tour-2")
		require.Nil(t, err)
		require.Equal(t, *inf2, *inf)

		val, err := mngr.UnfinishedTourn()
		require.Nil(t, err)
		assert.Equal(t, "tour-1", val)
	}
}

func SubTestNonUnfinished(mngr *usecases.DhtTourDataMngr) func(*testing.T) {
	return func(t *testing.T) {
		inf1 := &domain.TournInfo{
			ID:      "tour-1",
			Name:    "tour-1",
			Players: []*domain.Player{{Id: "manolo"}, {"el loco"}, {"de la mata"}},
			Winner:  &domain.Player{Id: "andy"},
		}
		inf2 := &domain.TournInfo{
			ID:      "tour-2",
			Name:    "tour-2",
			Players: []*domain.Player{{Id: "manolo"}, {"el loco"}, {"de la mata"}},
			Winner:  &domain.Player{Id: "el loco"},
		}
		inf3 := &domain.TournInfo{
			ID:      "tour-3",
			Name:    "tour-3",
			Players: []*domain.Player{{Id: "manolo"}, {"el loco"}, {"de la mata"}},
			Winner:  &domain.Player{Id: "manolo"},
		}
		require.Nil(t, mngr.SetTournInfo(inf3))
		require.Nil(t, mngr.SetTournInfo(inf1))
		require.Nil(t, mngr.SetTournInfo(inf2))

		val, err := mngr.UnfinishedTourn()
		require.Nil(t, err)
		assert.Equal(t, "", val)
	}
}

func ringOf(amount uint, t *testing.T) []*chord.Node {
	log := infrastruct.NewLogger().ToFile()

	entry, err := usecases.NewNode(
		localConfig(8001),
		nil,
		log)
	require.Nil(t, err)

	ans := []*chord.Node{entry}
	for i := uint(0); i < amount-1; i++ {
		o, err := usecases.NewNode(
			localConfig(uint(8002+i)),
			entry.RemoteNode,
			log,
		)
		require.Nil(t, err)
		ans = append(ans, o)
	}
	time.Sleep(time.Second * time.Duration(amount*3))
	return ans
}

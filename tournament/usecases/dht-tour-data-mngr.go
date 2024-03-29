package usecases

import (
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain/chord"
	"math/rand"
)

type DhtTourDataMngr struct {
	DhtStr     *Dht[string]
	DhtInfos   *Dht[[]*domain.TournInfo]
	DhtMatches *Dht[[]*domain.Pairing]
}

func (data *DhtTourDataMngr) SetTournInfo(info *domain.TournInfo) error {
	infos, err := data.DhtInfos.Get(tourInfosKey())
	if err != nil {
		infos = []*domain.TournInfo{}
	}
	idx := len(infos)
	for i := range infos {
		if infos[i].ID == info.ID {
			idx = i
			break
		}
	}
	if idx == len(infos) {
		infos = append(infos, info)
	} else {
		infos[idx] = info
	}
	return data.DhtInfos.Set(tourInfosKey(), infos)
}

func tourInfosKey() string {
	return "_tour-infos"
}

func (data *DhtTourDataMngr) SaveFiles(tourId string, files *map[string]string) error {
	m := *files
	for k, v := range m {
		if err := data.DhtStr.Set(k+"_"+tourId, v); err != nil {
			return err
		}
	}
	return nil
}

func (data *DhtTourDataMngr) File(tourId string, fileName string) (string, error) {
	ans, err := data.DhtStr.Get(fileName + "_" + tourId)
	if err != nil {
		return "", err
	}
	return ans, nil
}

func (data *DhtTourDataMngr) FileGroup(tourId string, fileNames []string) (map[string]string, error) {
	ans := make(map[string]string)
	for _, fileName := range fileNames {
		file, err := data.File(tourId, fileName)
		if err != nil {
			return nil, err
		}
		ans[fileName] = file
	}
	return ans, nil
}

func matchesKey(tourId string) string {
	return "_tour-" + tourId
}

func (data *DhtTourDataMngr) SaveMatch(match *domain.Pairing) error {
	val, err := data.DhtMatches.Get(matchesKey(match.TourId))
	if err != nil {
		val = []*domain.Pairing{}
	}
	val = append(val, match)
	return data.DhtMatches.Set(matchesKey(match.TourId), val)
}

func (data *DhtTourDataMngr) Matches(tourId string) ([]*domain.Pairing, error) {
	ans, err := data.DhtMatches.Get(matchesKey(tourId))
	if err == chord.ErrKeyNotFound {
		return []*domain.Pairing{}, nil
	} else if err != nil {
		return nil, err
	}
	return ans, nil
}

func (data *DhtTourDataMngr) GetTournInfo(tourId string) (*domain.TournInfo, error) {
	infos, err := data.DhtInfos.Get(tourInfosKey())
	if err != nil {
		return nil, err
	}
	for _, inf := range infos {
		if inf.ID == tourId {
			return inf, nil
		}
	}
	return nil, ErrInfoNotFound
}

func (data *DhtTourDataMngr) UnfinishedTourn() (string, error) {
	infos, err := data.DhtInfos.Get(tourInfosKey())

	if err != nil {
		if err.Error() == chord.ErrKeyNotFound.Error() {
			return "", ErrNotAnyUnfTournmnt
		}
		return "", err
	}
	not_run_candidates := []*domain.TournInfo{}
	for _, inf := range infos {
		if inf.Winner == nil {
			not_run_candidates = append(not_run_candidates, inf)
		}
	}
	if len(not_run_candidates) == 0 {
		return "", ErrNotAnyUnfTournmnt
	}
	// select a random tourmnt
	idx := rand.Intn(len(not_run_candidates))
	return not_run_candidates[idx].ID, nil

	//return "", ErrNotAnyUnfTournmnt
}

func (data *DhtTourDataMngr) GetAllIds() ([]string, error) {
	infos, err := data.DhtInfos.Get(tourInfosKey())
	if err != nil {
		return nil, err
	}
	ans := []string{}
	for _, inf := range infos {
		ans = append(ans, inf.ID)
	}
	return ans, nil
}

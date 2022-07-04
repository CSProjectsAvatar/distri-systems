package usecases

import (
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain/chord"
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
	if err != nil {
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
	if err == chord.ErrKeyNotFound {
		return "", nil
	} else if err != nil {
		return "", err
	}
	for _, inf := range infos {
		if inf.Winner == nil {
			return inf.ID, nil
		}
	}
	return "", nil
}

package usecases

import (
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain/chord"
)

type DataMap struct{}

func (dm *DataMap) LowerEq(upper []byte) []*chord.Data {
	return nil
}

func (dm *DataMap) Delete(data []*chord.Data) {

}

func (dm *DataMap) Save(data []*chord.Data) {

}

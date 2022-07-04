package usecases

import (
	"bytes"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain/chord"
	"sync"
)

type DataMap struct {
	data map[string]string
	mtx  sync.RWMutex
}

func (dm *DataMap) LowerEq(upper []byte) ([]*chord.Data, error) {
	var data []*chord.Data
	dm.mtx.RLock()
	defer dm.mtx.RUnlock()

	for k, v := range dm.data {
		bk := []byte(k)
		if bytes.Compare(bk, upper) <= 0 {
			data = append(data, &chord.Data{Key: bk, Value: v})
		}
	}
	return data, nil
}

func (dm *DataMap) Delete(data []*chord.Data) error {
	dm.mtx.Lock()
	defer dm.mtx.Unlock()

	for _, d := range data {
		delete(dm.data, string(d.Key))
	}
	return nil
}

func (dm *DataMap) Save(data []*chord.Data) error {
	dm.mtx.Lock()
	defer dm.mtx.Unlock()
	for _, d := range data {
		dm.data[string(d.Key)] = d.Value
	}
	return nil
}

func (dm *DataMap) Get(key []byte) (string, error) {
	dm.mtx.RLock()
	val, ok := dm.data[string(key)]
	dm.mtx.RUnlock()

	if !ok {
		return "", chord.ErrKeyNotFound
	}
	return val, nil
}

func NewDataMap() *DataMap {
	return &DataMap{
		data: make(map[string]string),
	}
}

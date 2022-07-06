package infrastruct

import (
	"fmt"
	"github.com/CSProjectsAvatar/distri-systems/tournament/domain/chord"
	"github.com/tidwall/buntdb"
	"log"
)

type BuntDb struct {
	*buntdb.DB
}

func NewBuntDb(name string) *BuntDb {
	db, err := buntdb.Open(fmt.Sprintf("%s.db", name))
	if err != nil {
		log.Print(err)
		panic(err)
	}
	return &BuntDb{DB: db}
}

func (db *BuntDb) LowerEq(upper []byte) ([]*chord.Data, error) {
	var data []*chord.Data
	err := db.View(func(tx *buntdb.Tx) error {
		return tx.DescendLessOrEqual("", string(upper), func(key, value string) bool {
			data = append(data, &chord.Data{Key: []byte(key), Value: value})
			return true
		})
	})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (db *BuntDb) Delete(data []*chord.Data) error {
	return db.Update(func(tx *buntdb.Tx) error {
		for _, d := range data {
			if _, err := tx.Delete(string(d.Key)); err != nil {
				return err
			}
		}
		return nil
	})
}

func (db *BuntDb) Save(data []*chord.Data) error {
	return db.Update(func(tx *buntdb.Tx) error {
		for _, d := range data {
			if _, _, err := tx.Set(string(d.Key), d.Value, nil); err != nil {
				return err
			}
		}
		return nil
	})
}

func (db *BuntDb) Get(key []byte) (string, error) {
	var val string
	err := db.View(func(tx *buntdb.Tx) error {
		var err error
		val, err = tx.Get(string(key))
		return err
	})
	if err != nil {
		if err.Error() == "not found" {
			return "", chord.ErrKeyNotFound
		}
		return "", err
	}
	return val, nil
}

package store

import (
	"encoding/json"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/tidwall/buntdb"
)

type Store struct {
	db *buntdb.DB
}

type StoreSingleResult struct {
	Id   string                 `json:"_id"`
	Data map[string]interface{} `json:"data"`
}

type StoreUpdateResult struct {
	Id       string                 `json:"_id"`
	Data     map[string]interface{} `json:"data"`
	PrevData map[string]interface{} `json:"prevData"`
	Replaced bool                   `json:"replaced"`
}

type StoreDeleteResult struct {
	Id       string                 `json:"_id"`
	PrevData map[string]interface{} `json:"prevData"`
}

type StoreMultiResult struct {
	Data []StoreSingleResult `json:"data"`
}

func OpenStore(file string) Store {
	// Open the data.db file. It will be created if it doesn't exist.
	db, err := buntdb.Open(filepath.FromSlash(file))
	if err != nil {
		logrus.WithError(err).Fatal("Could not open the database")
	}

	return Store{db: db}
}

func (s Store) Close() {
	s.db.Close()
}

func (s Store) Get(id string) (result StoreSingleResult, err error) {
	result = StoreSingleResult{
		Id:   id,
		Data: make(map[string]interface{}),
	}
	if err = s.db.View(func(tx *buntdb.Tx) error {
		var data map[string]interface{}
		val, err := tx.Get(id)
		if err != nil {
			logrus.WithError(err).WithFields(map[string]interface{}{
				"id": id,
			}).Error("Could not get item")
			return err
		}

		err = json.Unmarshal([]byte(val), &data)
		if err != nil {
			logrus.WithError(err).WithFields(map[string]interface{}{
				"id": id,
			}).Error("Could not unmarshal")
			return err
		}

		result.Data = data
		return nil
	}); err != nil {
		logrus.WithError(err).WithField("id", id).Error("Could not find item")
	}

	return result, err
}

func (s Store) Find(params map[string]interface{}) (result StoreMultiResult, err error) {
	result = StoreMultiResult{
		Data: make([]StoreSingleResult, 0),
	}
	err = s.db.View(func(tx *buntdb.Tx) error {
		err = tx.Ascend("", func(key, value string) bool {
			var data map[string]interface{}
			err := json.Unmarshal([]byte(value), &data)
			if err != nil {
				logrus.WithError(err).WithFields(map[string]interface{}{
					"key":   key,
					"value": value,
				}).Error("Could not unmarshal")
				return false
			}

			result.Data = append(result.Data, StoreSingleResult{
				Id:   key,
				Data: data,
			})
			return true
		})
		return err
	})

	return result, err
}

func (s Store) Add(data map[string]interface{}) (result StoreSingleResult, err error) {
	result = StoreSingleResult{
		Id:   "",
		Data: data,
	}
	id := uuid.New().String()
	str, err := json.Marshal(data)
	if err != nil {
		return result, err
	}

	if err = s.db.Update(func(tx *buntdb.Tx) error {
		_, _, err = tx.Set(id, string(str), nil)
		return err
	}); err != nil {
		logrus.WithError(err).Error("Could not add item")
	}

	result.Id = id

	return result, err
}

func (s Store) Update(id string, data map[string]interface{}) (result StoreUpdateResult, err error) {
	result = StoreUpdateResult{
		Id:   id,
		Data: data,
	}
	str, err := json.Marshal(data)
	if err != nil {
		return result, err
	}

	if err = s.db.Update(func(tx *buntdb.Tx) error {
		var prevVal map[string]interface{}
		prevValS, replaced, err := tx.Set(id, string(str), nil)
		if err != nil {
			return err
		}

		err = json.Unmarshal([]byte(prevValS), &prevVal)
		result.PrevData = prevVal
		result.Replaced = replaced
		return err
	}); err != nil {
		logrus.WithError(err).Error("Could not add item")
	}

	return result, err
}

func (s Store) Delete(id string) (result StoreDeleteResult, err error) {
	result = StoreDeleteResult{
		Id: id,
	}

	if err = s.db.Update(func(tx *buntdb.Tx) error {
		var prevVal map[string]interface{}
		prevValS, err := tx.Delete(id)
		if err != nil {
			return err
		}

		err = json.Unmarshal([]byte(prevValS), &prevVal)
		result.PrevData = prevVal
		return err
	}); err != nil {
		logrus.WithError(err).Error("Could not add item")
	}

	return result, err
}

package storage

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

var _ Storage = (*fsStorage)(nil)

type fsStorage struct {
	filename string
	m        *sync.Mutex
}

func (f *fsStorage) Save(v interface{}) {
	go func() {
		f.m.Lock()
		file, err := os.OpenFile(f.filename, os.O_CREATE|os.O_WRONLY, os.ModePerm)
		storageFunc := func() {
			if err != nil {
				log.Println(err)
				return
			}
			err = json.NewEncoder(file).Encode(v)
			if err != nil {
				log.Println(err)
				return
			}
		}
		storageFunc()
		_ = file.Close()
		f.m.Unlock()
	}()
}

func (f *fsStorage) Load(v interface{}) error {
	f.m.Lock()
	file, err := os.Open(f.filename)
	storageFunc := func() error {
		if err != nil {
			return err
		}
		err = json.NewDecoder(file).Decode(v)
		if err != nil {
			return err
		}
		return nil
	}
	err = storageFunc()
	_ = file.Close()
	f.m.Unlock()
	return err
}

func NewFileStorage(filename string) (Storage, error) {
	_, err := os.OpenFile(filename, os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	return &fsStorage{
		filename: filename,
		m:        &sync.Mutex{},
	}, nil
}

var _ = NewFileStorage

// Copyright 2020 Rogchap. All Rights Reserved.

package db

import (
	"path/filepath"

	badger "github.com/dgraph-io/badger/v2"
	"google.golang.org/protobuf/proto"
)

type Store struct {
	db *badger.DB
}

// NewStore creates a new store to save user data
func NewStore(path string) (*Store, error) {
	dbPath := filepath.Join(path, "db")
	db, err := badger.Open(badger.DefaultOptions(dbPath))
	if err != nil {
		return nil, err
	}

	return &Store{db}, nil
}

func (s *Store) GetWorkspace(key string) (*Workspace, error) {
	w := &Workspace{}
	err := s.db.View(func(txn *badger.Txn) error {
		data, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		return data.Value(func(val []byte) error {
			return proto.Unmarshal(val, w)
		})
	})
	return w, err
}

func (s *Store) SetWorkspace(key string, w *Workspace) error {
	data, err := proto.Marshal(w)
	if err != nil {
		return err
	}
	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), data)
	})
}

func (s *Store) Get(key []byte) (val []byte, rtnErr error) {
	rtnErr = s.db.View(func(txn *badger.Txn) error {
		data, err := txn.Get(key)
		if err != nil {
			return err
		}
		val, err = data.ValueCopy(val)
		return err
	})
	return
}

func (s *Store) Set(key, val []byte) error {
	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, val)
	})
}

func (s *Store) Close() {
	if s == nil {
		return
	}
	s.db.Close()
}

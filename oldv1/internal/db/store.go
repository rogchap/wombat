// Copyright 2020 Rogchap. All Rights Reserved.

package db

import (
	"path/filepath"

	badger "github.com/dgraph-io/badger/v2"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

type dbLogger struct {
	*zap.SugaredLogger
}

func (l *dbLogger) Warningf(template string, args ...interface{}) {
	l.Warnf(template, args...)
}

// Store is a wrapper to a DB to store data to disk
type Store struct {
	db *badger.DB
}

// NewStore creates a new store to save user data
func NewStore(path string, logger *zap.SugaredLogger) (*Store, error) {
	dbPath := filepath.Join(path, "db")
	opts := badger.DefaultOptions(dbPath)
	opts = opts.WithLogger(&dbLogger{logger})
	db, err := badger.Open(opts)
	if err != nil {
		return nil, err
	}

	return &Store{db}, nil
}

// GetWorkspace will retrieve the data for the workspace and unmarshal to the Workspace struct
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

// SetWorkspace will store the workspace data in the store
func (s *Store) SetWorkspace(key string, w *Workspace) error {
	data, err := proto.Marshal(w)
	if err != nil {
		return err
	}
	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), data)
	})
}

// Get is a generic function to retrive the raw data in the store by key
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

// Set is a generic function to store raw data for a given key
func (s *Store) Set(key, val []byte) error {
	return s.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, val)
	})
}

// Close will close the underling data store
func (s *Store) Close() {
	if s == nil {
		return
	}
	s.db.Close()
}

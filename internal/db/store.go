// Copyright 2020 Rogchap. All Rights Reserved.

package db

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"google.golang.org/protobuf/proto"
)

type Store struct {
	dbPath string
}

// NewStore is a poor mans DB, should replace with domething like BadgerDB
func NewStore(path string) *Store {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0700)
	}
	dbPath := filepath.Join(path, "workspace.db")
	return &Store{dbPath}
}

func (s *Store) Get() *Workspace {
	raw, err := ioutil.ReadFile(s.dbPath)
	if err != nil {
		println(err.Error())
		return nil
	}
	w := &Workspace{}
	if err := proto.Unmarshal(raw, w); err != nil {
		println(err.Error())
		return nil
	}
	return w
}

func (s *Store) Put(w *Workspace) {
	raw, err := proto.Marshal(w)
	if err != nil {
		println(err.Error())
		return
	}
	if err := ioutil.WriteFile(s.dbPath, raw, 0644); err != nil {
		println(err.Error())
	}
}

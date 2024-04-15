package database

import (
	bolt "go.etcd.io/bbolt"
)

var (
	db  *bolt.DB
	err error
)

func Init() (*bolt.DB, error) {
	db, err = bolt.Open(DBFileName, DBFileMode, nil)
	if err != nil {
		return nil, err
	}
	return db, nil
}

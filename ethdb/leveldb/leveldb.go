package leveldb

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type database struct {
	fn string
	db *leveldb.DB
}

func NewDatabase(file string, opt *opt.Options) (*database, error) {
	db, err := leveldb.OpenFile(file, opt)
	if err != nil {
		return nil, err
	}

	return &database{fn: file, db: db}, nil
}

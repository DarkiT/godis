package database

import (
	"github.com/darkit/godis/datastruct/dict"
)

func makeTestDB() *DB {
	return &DB{
		data:       dict.MakeConcurrent(dataDictSize),
		versionMap: dict.MakeConcurrent(dataDictSize),
		ttlMap:     dict.MakeConcurrent(ttlDictSize),
		addAof:     func(line CmdLine) {},
	}
}

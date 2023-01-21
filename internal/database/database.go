package database

import (
	"github.com/gizmo-ds/misstodon/internal/database/buntdb"
	"github.com/gizmo-ds/misstodon/internal/database/memory"
)

type DbType = string

const (
	DbTypeMemory DbType = "memory"
	DbTypeBuntDb DbType = "buntdb"
)

type Database interface {
	Get(server, key string) (string, bool)
	Set(server, key, value string, expire int64) error
	Close() error
}

func NewDatabase(dbType, address string) Database {
	switch dbType {
	case DbTypeBuntDb:
		return buntdb.NewDatabase(address)
	case DbTypeMemory:
		return memory.NewDatabase()
	default:
		panic("unknown database type")
	}
	return nil
}

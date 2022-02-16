package database

// db 只有数据操作，也就是只会保存玩家数据落地

import (
	. "db"
)

type IDBHandler interface {
	CreateTable(mysql *Mysql) error
	PrimaryKey() (interface{}, interface{})
	Select(mysql *Mysql, keys ...interface{}) (bool, error)
	Insert(mysql *Mysql, args ...interface{}) (bool, error)
	Updata(mysql *Mysql, args ...interface{}) (bool, error)
	Delete(mysql *Mysql, keys ...interface{}) (bool, error)
	FromBytes(blob []byte) error
	ToBytes() ([]byte, error)
}

type IDBHandleManager interface {
	Register(key string, dataDb IDBHandler)
	GetDataDb(key string) IDBHandler
}
type DBHandleManager struct {
	dataHandlers map[string]IDBHandler
}

func NewDBHandleManager() *DBHandleManager {
	return &DBHandleManager{
		dataHandlers: make(map[string]IDBHandler),
	}
}

func (d *DBHandleManager) Register(key string, dataDb IDBHandler) {
	if _, ok := d.dataHandlers[key]; ok {
		return
	}
	d.dataHandlers[key] = dataDb
}

func (d *DBHandleManager) GetDataDb(key string) IDBHandler {
	if _, ok := d.dataHandlers[key]; !ok {
		return nil
	}
	return d.dataHandlers[key]
}

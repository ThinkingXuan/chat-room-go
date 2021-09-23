package myleveldb

import (
	"github.com/golang/glog"
	"github.com/syndtr/goleveldb/leveldb"
)

var (
	db *leveldb.DB
)

//InitLevelDB 初始化leveldb数据库
func InitLevelDB() (err error) {
	db, err = leveldb.OpenFile("path/to/db", nil)
	if err != nil  {
		glog.Error(err)
		return err
	}
	return nil
}

// Close 关闭数据库
func Close() {
	err := db.Close()
	if err != nil {
		return
	}
}

package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/viper"
)

var (
	db *gorm.DB //deal normal database
)

// Model base
type Model struct {
	ID        string `json:"id" gorm:"primary_key;not null"`
	CreatedAt string `json:"created_at"`
}

//InitSQLite 初始化数据库
func InitSQLite() (err error) {

	// 获取SQLite存储路径
	SQLitePath := viper.GetString("sqlite_path")

	db, err = gorm.Open("sqlite3", SQLitePath)
	if err != nil {
		return
	}
	return db.DB().Ping()
}

// Close 关闭数据库
func Close() {
	err := db.Close()
	if err != nil {
		return
	}
}

// InitDBTable 初始化数据库表
func InitDBTable() {

	//禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	db.SingularTable(true)
	//开启日志
	db.LogMode(true)
	//模型绑定
	db.AutoMigrate(&User{})

}

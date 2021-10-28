package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"time"
)

var (
	db *gorm.DB //deal normal database
)

// Model base
type Model struct {
	ID        uint64    `json:"id" gorm:"primary_key;not null;index:idx_id"`
	CreatedAt time.Time `json:"created_at" gorm:"index:idx_create_at"`
}

//InitMySql 初始化数据库
func InitMySql() (err error) {

	// 获取SQLite存储路径
	//SQLitePath := viper.GetString("sqlite_path")
	//
	//db, err = gorm.Open("sqlite3", SQLitePath)
	//if err != nil {
	//	return
	//}

	dbURL, dbPort, dbName, dbPassword, dbDbName := viper.GetString("mysql.url"), viper.GetString("mysql.port"), viper.GetString("mysql.username"), viper.GetString("mysql.password"), viper.GetString("mysql.dbname")
	url1 := dbName + ":" + dbPassword + "@tcp(" + dbURL + ":" + dbPort + ")/" + dbDbName + "?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", url1)
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
func InitDBTable(maxConn int, maxIdleConn int) {

	//禁用默认表名的复数形式，如果置为 true，则 `User` 的默认表名是 `user`
	db.SingularTable(true)
	//开启日志
	db.LogMode(false)
	//模型绑定
	db.AutoMigrate(&User{}, &Room{}, &Message{})

	db.DB().SetMaxOpenConns(maxConn)     //设置最大的连接
	db.DB().SetMaxIdleConns(maxIdleConn) //设置最大的空闲连接数

}

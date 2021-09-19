package myleveldb

import (
	"github.com/golang/glog"
)

// 存储方式
// user.{username}: userBytes

var userKey = "users"

// CreateUser create a user
func CreateUser(username string, newUserBytes []byte) error {
	key := userKey + "." + username
	err := db.Put([]byte(key), newUserBytes, nil)
	if err != nil {
		glog.Error(err)
		return err
	}
	return nil
}

//func SelectUserByUsername(username string) (*User, int64) {
//	var u User
//	rowAffect := db.Model(&User{}).Where("username = ?", username).First(&u).RowsAffected
//	return &u, rowAffect
//}
//
//func SelectResUserByUsername(username string) (*rr.ResUser, int64) {
//	var u rr.ResUser
//	rowAffect := db.Table("user").Where("username = ?", username).First(&u).RowsAffected
//	return &u, rowAffect
//}

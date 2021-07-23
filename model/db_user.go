package model

import (
	"chat-room-go/api/router/rr"
	"github.com/golang/glog"
)

// User user model
type User struct {
	Model
	Username  string `json:"username,omitempty" gorm:"unique;not null;type:varchar(100);"`
	FirstName string `json:"firstName,omitempty" gorm:"not null;type:varchar(20)"`
	LastName  string `json:"lastName,omitempty" gorm:"not null;type:varchar(20)"`
	Email     string `json:"email,omitempty" gorm:"not null;type:varchar(20)"`
	Password  string `json:"password,omitempty" gorm:"not null;type:varchar(100)"`
	Phone     string `json:"phone,omitempty" gorm:"not null;type:varchar(20)"`
}

// CreateUser create a user
func CreateUser(newUser *rr.ReqUser) error {

	user := &User{
		Username:  newUser.Username,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Email:     newUser.Email,
		Password:  newUser.Password,
		Phone:     newUser.Phone,
	}
	//sql := "insert into user(`id`,`created_at`,`username`,`first_name`,`last_name`,`email`,`password`,`phone`) values(?,?,?,?,?,?,?,?);"
	//err := db.Exec(sql, user.ID, user.CreatedAt, user.Username, user.FirstName, user.LastName, user.Email, user.Password, user.Phone).Error
	//if err != nil {
	//	glog.Error(err)
	//	return err
	//}

	err := db.Model(&User{}).Create(user).Error
	if err != nil {
		glog.Error(err)
		return err
	}
	return nil
}

func SelectUserByUsername(username string) (*User, int64) {
	var u User
	rowAffect := db.Model(&User{}).Where("username = ?", username).First(&u).RowsAffected
	return &u, rowAffect
}

func SelectResUserByUsername(username string) (*rr.ResUser, int64) {
	var u rr.ResUser
	rowAffect := db.Table("user").Where("username = ?", username).First(&u).RowsAffected
	return &u, rowAffect
}

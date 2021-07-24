package model

import (
	"chat-room-go/api/router/rr"
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

	// 去重检测
	//var oldUser User
	//err := db.Table("user").Where("username = ?", newUser.Username).First(&oldUser).Error
	//if err != nil {
	//	return err
	//}
	//
	//if oldUser.Username == newUser.Username {
	//	return errors.New("user exist")
	//}

	err := db.Model(&User{}).Create(user).Error
	if err != nil {
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

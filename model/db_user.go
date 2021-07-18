package model

import (
	"chat-room-go/api/router/rr"
	"chat-room-go/util"
)

// User user model
type User struct {
	Model
	Username  string `json:"username,omitempty" gorm:"not null"`
	FirstName string `json:"firstName,omitempty" gorm:"not null"`
	LastName  string `json:"lastName,omitempty" gorm:"not null"`
	Email     string `json:"email,omitempty" gorm:"not null"`
	Password  string `json:"password,omitempty" gorm:"not null"`
	Phone     string `json:"phone,omitempty" gorm:"not null"`
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
	user.ID = util.GetSnowflakeID()
	user.CreatedAt = util.GetNowTime()
	return db.Model(User{}).Create(user).Error
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

package model

import "github.com/google/uuid"

type UserToken struct {
	UserId uint   `json:"user_id"`
	Token  string `gorm:"unique" json:"token"`
}

func NewToken(u *User) (*UserToken, error) {
	ut := &UserToken{}
	ut.UserId = u.ID
	ut.Token = uuid.New().String()
	err := db.Create(&ut).Error
	if err != nil {
		return nil, err
	}
	return ut, nil
}

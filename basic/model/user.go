package model

import (
	"github.com/SteveZhangF/brewer/basic/errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	UserName     string `json:"username"`
	Password     string `json:"-"`
	IsGuest      bool
	Level        int
	Roles        []string `json:"roles" gorm:"-"`
	Introduction string   `json:"introduction"`
	Avatar       string   `json:"avatar"`
	Token        string   `json:"-" gorm:"-"`

	// roles: ['editor'],
	// introduction: 'I am an editor',
	// avatar: 'https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif',
	// name: 'Normal Editor'
}

func (u *User) AfterFind(tx *gorm.DB) (err error) {
	u.Roles = []string{"admin"}
	return
}

func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}

func (user *User) Login(username string, password string) (*UserToken, error) {
	u := &User{}
	err := db.Where("user_name = ?", username).First(u).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.NotFound.Error()
	}
	if err != nil {
		return nil, err
	}

	err = u.CheckPassword(password)
	if err != nil {
		return nil, errors.AuthInvalidToken.Error()
	}
	token, err := NewToken(u)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func (user *User) SetPassword(oldPassword string, newPassword string) error {
	err := user.CheckPassword(oldPassword)
	if err != nil {
		return err
	}
	err = user.HashPassword(newPassword)
	if err != nil {
		return err
	}
	err = db.Save(user).Error
	if err != nil {
		return err
	}
	return nil
}

func UserByToken(token string) (*User, error) {
	ut := &UserToken{}
	err := db.Where("token = ?", token).First(ut).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.NotFound.Error()
	}
	if err != nil {
		return nil, err
	}
	u := &User{}
	err = db.Where("id = ?", ut.UserId).First(u).Error
	if err == gorm.ErrRecordNotFound {
		return nil, errors.NotFound.Error()
	}
	if err != nil {
		return nil, err
	}
	u.Token = token
	return u, nil
}

func (user *User) Logout() error {
	if user.Token == "" {
		return errors.NotFound.Error()
	}
	ut := &UserToken{}
	db.Where("token = ?", user.Token).Delete(ut)
	return nil
}

func (user *User) Create() error {
	return db.Create(user).Error
}

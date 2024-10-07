package model

import (
	"kenneth/backend/basic"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	db = basic.GetDatabase()
	db = db.Debug()

	db.AutoMigrate(User{})
	db.AutoMigrate(UserToken{})
	db.AutoMigrate(Audit{})
}

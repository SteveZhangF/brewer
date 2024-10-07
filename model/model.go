package model

import (
	"time"

	"gorm.io/gorm"
)

// Model a basic GoLang struct which includes the following fields: ID, CreatedAt, UpdatedAt, DeletedAt
// It may be embedded into your model or you may build your own model without it
//
//	type User struct {
//	  gorm.Model
//	}
type Model struct {
	ID uint `gorm:"primarykey" json:"id"`
	BasicModel
}

type BasicModel struct {
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
	Remarks   string         `gorm:"type:varchar(255);default:'';comment:备注" json:"remarks"`
	Class     string         `gorm:"type:varchar(32);default:'';comment:分类" json:"__class__"`
}

type CodeModel struct {
	Code string `gorm:"type:varchar(32);default:'';comment:编码" json:"code"`
	Name string `gorm:"type:varchar(32);default:'';comment:名称" json:"name"`
	Model
}

func (model *BasicModel) AfterFind(tx *gorm.DB) (err error) {
	model.Class = "BasicModel"
	return nil
}

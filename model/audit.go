package model

type Audit struct {
	Model
	UserId uint   `json:"user_id"`
	Name   string `json:"name"`
	Data   string `json:"data" gorm:"type:text"`
	Method string `json:"method"`
	URI    string `json:"uri"`
}

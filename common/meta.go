package common

import "time"

type Meta struct {
	Id         int
	CreateTime int64 `json:"created_at" gorm:"column:created_at"`
	UpdateTime int64 `json:"updated_at" gorm:"column:updated_at"`
}

func CreateNewMeta() *Meta {
	return &Meta{
		Id:         0,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}
}

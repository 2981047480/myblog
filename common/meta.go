package common

import "time"

type Meta struct {
	Id         int
	CreateTime int64
	UpdateTime int64
}

func CreateNewMeta() *Meta {
	return &Meta{
		Id:         0,
		CreateTime: time.Now().UnixMicro(),
		UpdateTime: time.Now().UnixMicro(),
	}
}

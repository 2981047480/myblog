package impl

import (
	"vblog/app/user"
	"vblog/common/config"

	"gorm.io/gorm"
)

type CreateTokenImpl struct {
	i *gorm.DB
	user.Service
}

// func NewCreateTokenImpl() *CreateTokenImpl {
// 	db, _ := config.ReadDBConf(config.DBConfigFile).GetConn()
// 	return &CreateTokenImpl{
// 		i: db,
// 	}
// }

func (i *CreateTokenImpl) 
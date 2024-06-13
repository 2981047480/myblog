package impl

import (
	"context"
	"vblog/app/token"
	"vblog/app/user"
	"vblog/common/config"
	"vblog/ioc"

	"gorm.io/gorm"
)

type CreateTokenImpl struct {
	DB   *gorm.DB
	User user.Service
}

func NewCreateTokenImpl() *CreateTokenImpl {
	db, _ := config.ReadDBConf(config.DBConfigFile).GetConn()
	return &CreateTokenImpl{
		DB: db,
	}
}
func (i *CreateTokenImpl) Init() {
	db, _ := config.ReadDBConf(config.DBConfigFile).GetConn()
	i.DB = db
	i.User = ioc.ControllerImpl.Get(user.AppName).(user.Service)
}

func (i *CreateTokenImpl) IssueToken(ctx *context.Context, in *token.IssueTokenRequest) (*token.Token, error) {
	// 1、第一步先检查用户账号密码是否能匹配，再次之前，首先肯定要把用户给查出来
	queryrequest := user.NewQueryUserRequest()
	queryrequest.Username = in.Username
	i.Init()
	user_set, err := i.User.QueryUser(ctx, queryrequest)
	if err != nil {
		return nil, err
	}
	// 判断下user_set是否为空
	if len(user_set.Items) == 0 {
		return nil, token.ErrAuthFailed
	}

	// 2、这个时候查出来了，就可以对比了
	u := user_set.Items[0]
	if err := u.ValidatePassword(in.Password); err != nil {
		return nil, err
	}

	// 3、查出来没问题就可以颁发证书了
	token, err := i.IssueToken(ctx, in)
	if err != nil {
		return nil, err
	}

	// 4、这里第一次写的时候忘了令牌入库了
	if err := i.DB.WithContext(*ctx).Save(token).Create(in).Error; err != nil {
		return nil, err
	}

	// 5、做完操作返回
	return token, nil
}

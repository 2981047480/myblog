package impl

import (
	"context"
	"fmt"
	"vblog/app/token"
	"vblog/app/user"
	"vblog/app/user/impl"
	"vblog/common/config"
	"vblog/exception"
	"vblog/ioc"

	"gorm.io/gorm"
)

func init() {
	ioc.ControllerImpl.Register(token.AppName, &TokenServiceImpl{})
}

type TokenServiceImpl struct {
	DB   *gorm.DB
	User user.Service
}

func NewServiceImpl(u *impl.UserServiceImpl) *TokenServiceImpl {
	db, _ := config.C().GetConn()
	return &TokenServiceImpl{
		DB:   db,
		User: u,
	}
}
func (i *TokenServiceImpl) Init() error {
	db, _ := config.C().GetConn()
	i.DB = db
	i.User = ioc.ControllerImpl.Get(user.AppName).(user.Service) // 这里没初始化user 前面的user是初始化过的 这里没初始化
	return nil
}

func (i *TokenServiceImpl) IssueToken(ctx *context.Context, in *token.IssueTokenRequest) (*token.Token, error) {
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
	token := token.NewToken(*u)

	// 4、这里第一次写的时候忘了令牌入库了
	if err := i.DB.WithContext(*ctx).Create(token).Error; err != nil {
		return nil, err
	}

	// 5、做完操作返回
	return token, nil
}

func (i *TokenServiceImpl) RevolkToken(ctx *context.Context, in *token.RevolkTokenRequest) (*token.Token, error) {
	// 1、不能说随便来一个要取消就给他取消，验证下at和rt是否一致
	tk := token.DefaultToken()
	query := i.DB.WithContext(*ctx)
	err := query.Where("access_token=?", in.Access_token).First(tk).Error
	if err == gorm.ErrRecordNotFound {
		return nil, exception.ErrNotFound("Token not found error!")
	}

	// 判断at和rt是否一致
	if in.Refresh_token != tk.RefreshToken {
		return nil, fmt.Errorf("Accesstoken和Refreshtoken不一致")
	}

	// 2、校验完成后可以直接把这一行删掉
	if err := i.DB.WithContext(*ctx).Where("access_token=?", in.Access_token).Delete(token.Token{}).Error; err != nil {
		return nil, err
	}

	return tk, nil
}

func (i *TokenServiceImpl) ValicateToken(ctx *context.Context, in *token.ValicateTokenRequest) (*token.Token, error) {
	// 1、验证下at和rt是否一致
	tk := token.DefaultToken()
	query := i.DB.WithContext(*ctx)
	err := query.Where("access_token=?", in.Access_token).First(tk).Error
	if err == gorm.ErrRecordNotFound {
		return nil, exception.ErrNotFound("Token not found error!")
	}

	// 2、分别验证Refresh token和Access token是否过期
	// 之所以先验证Refresh token，是因为如果Refresh token过期了就没必要看Access token了，反之则不然
	if err := tk.RefreshExpired(); err != nil {
		return nil, err
	}
	if err := tk.AccessExpired(); err != nil {
		return nil, err
	}

	// 3、满足at和rt一致、且都未过期，可以把token返回出去，验证成功
	// 补充一下 这里之前忘了给role值了 导致一直返回role为0了
	queryuser := user.NewQueryUserRequest()
	queryuser.Username = tk.UserName
	u, err := i.User.QueryUser(ctx, queryuser)
	if err != nil {
		return nil, err
	}
	if u.Total == 0 || len(u.Items) == 0 {
		return nil, fmt.Errorf("user not found")
	}
	tk.Role = u.Items[0].R
	return tk, nil
}

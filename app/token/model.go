package token

import (
	"time"
	"vblog/app/user"
	"vblog/common"

	"github.com/rs/xid"
)

type Token struct {
	// 该Token是颁发
	UserId int `json:"user_id" gorm:"column:user_id"`
	// 人的名称， user_name
	UserName string `json:"username" gorm:"column:username"`
	// 办法给用户的访问令牌(用户需要携带Token来访问接口)
	AccessToken string `json:"access_token" gorm:"column:access_token"`
	// 过期时间(2h), 单位是秒
	AccessTokenExpiredAt int `json:"access_token_expired_at" gorm:"column:access_token_expired_at"`
	// 刷新Token
	RefreshToken string `json:"refresh_token" gorm:"column:refresh_token"`
	// 刷新Token过期时间(7d)
	RefreshTokenExpiredAt int `json:"refresh_token_expired_at" gorm:"column:refresh_token_expired_at"`

	// 创建时间
	CreatedAt int64 `json:"created_at" gorm:"column:created_at"`
	// 更新实现
	UpdatedAt int64 `json:"updated_at" gorm:"column:updated_at"`

	// 额外补充信息, gorm忽略处理
	Role user.Role `json:"role" gorm:"-"`

	common.Print
}
 
func NewToken(u user.User) *Token {
	return &Token{
		UserId:                u.Id,
		UserName:              u.Username,
		AccessToken:           xid.New().String(),
		AccessTokenExpiredAt:  3600, // 一小时
		RefreshToken:          xid.New().String(),
		RefreshTokenExpiredAt: 3600 * 4, // 四小时
		CreatedAt:             time.Now().Unix(),
		UpdatedAt:             time.Now().Unix(),
		Role:                  u.R,
	}
}

func DefaultToken() *Token {
	return &Token{}
}

func (t *Token) CreateTime() time.Time {
	return time.Unix(t.CreatedAt, 0)
}

func (t *Token) UpdateTime() time.Time {
	return time.Unix(t.UpdatedAt, 0)
}

func (t *Token) AccessTimeDuration() time.Duration {
	return time.Duration(t.AccessTokenExpiredAt * int(time.Second))
}

func (t *Token) RefreshTimeDuration() time.Duration {
	return time.Duration(t.RefreshTokenExpiredAt * int(time.Second))
}

func (t *Token) AccessExpired() error {
	expiretime := t.CreateTime().Add(t.AccessTimeDuration())
	if time.Since(expiretime).Seconds() > 0 {
		return ErrAccessTokenExpired
	}
	return nil
}

func (t *Token) RefreshExpired() error {
	expiretime := t.CreateTime().Add(t.RefreshTimeDuration())
	if time.Since(expiretime).Seconds() > 0 {
		return ErrRefreshTokenExpired
	}
	return nil
}

// type BaseToken struct {
// 	CreateTime int64 `gorm:"column:created_at"`
// 	ExpireTime int64
// 	*common.Print
// }

// func NewBaseToken(t int64) *BaseToken {
// 	return &BaseToken{
// 		CreateTime: time.Now().Unix(),
// 		ExpireTime: t,
// 	}
// }

// type AccessToken struct {
// 	*BaseToken
// }

// type RefreshToken struct {
// 	*BaseToken
// }

// type Token struct {
// 	Username string
// 	UserId int
// 	AccessToken
// 	RefreshToken
// } 想了想这样入库的时候可能会有坑 因此不采取这种方式了

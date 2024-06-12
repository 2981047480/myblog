package user

import (
	"encoding/json"
	"fmt"
	"log"
	"vblog/common"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserRequest struct {
	// 用户名
	Username string `json:"username" gorm:"column:username"`
	// 密码
	Password string `json:"password" gorm:"column:password"`
	// Nickname string `json:"nickname" gorm:"nickname"`
	R     Role              `json:"role" gorm:"column:role"`
	Label map[string]string `json:"label" gorm:"column:label;serializer:json"`
}

func (u *CreateUserRequest) Validator() error {
	if u.Username == "" {
		return fmt.Errorf("用户名必填")
	}

	return nil
}

// 对用户名做加密处理
func (u *CreateUserRequest) HashPassword() error {
	cyptoPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}
	u.Password = string(cyptoPass)
	return nil
}

func (u *CreateUserRequest) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// user属性：
// name string
// nickname string
// password string
// createdtime int64
// changedtime int64
// id int
// role Role
type User struct {
	*common.Meta
	*CreateUserRequest
}

func CreateNewUserRequest() *CreateUserRequest {
	return &CreateUserRequest{
		Username: "default",
		// Nickname: "default",
		Password: "default",
		R:        visitor,
		Label:    map[string]string{},
	}
}

func CreateNewUser(u *CreateUserRequest) *User {
	return &User{
		Meta:              common.CreateNewMeta(),
		CreateUserRequest: u,
	}
}

func (u *User) String() string {
	str, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
		return ""
	}
	return string(str)
}

func (u *User) Tablename() string {
	return "users"
}

type UserSet struct {
	Total int64   `json:"total"`
	Items []*User `json:"items"`
}

func CreateNewUserSet() *UserSet {
	return &UserSet{
		Items: []*User{},
	}
}

package user

import (
	"encoding/json"
	"log"
	"vblog/common"
)

type UserRequest struct {
	// 用户名
	Username string `json:"username" gorm:"username"`
	// 密码
	Password string `json:"password" gorm:"password"`
	Nickname string `json:"nickname" gorm:"nickname"`
	R        Role   `json:"role" gorm:"role"`
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
	*UserRequest
}

func CreateNewUserRequest() *UserRequest {
	return &UserRequest{
		Username: "default",
		Nickname: "default",
		Password: "default",
		R:        visitor,
	}
}

func CreateNewUser() *User {
	return &User{
		Meta:        common.CreateNewMeta(),
		UserRequest: CreateNewUserRequest(),
	}
}

func (u *User) String() string {
	str, err := json.Marshal(u)
	if err != nil {
		log.Println(err)
		return "111"
	}
	return string(str)
}

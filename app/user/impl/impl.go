package impl

import (
	"context"
	"fmt"
	"log"
	"vblog/app/user"
	"vblog/common/config"
	"vblog/ioc"

	"gorm.io/gorm"
)

var filename = "/Users/zephyrzhao/Documents/vblog/myblog/common/config/db.yaml"

func NewUserServiceImpl() *UserServiceImpl {
	db, _ := config.ReadDBConf(filename).GetConn()
	return &UserServiceImpl{
		database: db,
	}
}

type UserServiceImpl struct {
	database *gorm.DB
	Test     string
}

func init() {
	ioc.ControllerImpl.Register(user.AppName, &UserServiceImpl{})
}

func (u *UserServiceImpl) Init() error {
	var err error
	u.database, err = config.ReadDBConf(filename).GetConn()
	if err != nil {
		log.Println("get DB failed")
		return err
	}
	return nil
}

func (u *UserServiceImpl) CreateUser(ctx *context.Context, in *user.CreateUserRequest) (*user.User, error) {
	// 校验请求合法性 这个后面补
	if in.Username == "" {
		return nil, fmt.Errorf("Username不能为空")
	}

	// 密码加密
	if err := in.HashPassword(); err != nil {
		return nil, err
	}

	// 创建user对象
	fmt.Printf("%#v", in)
	ins := user.CreateNewUser(in)

	// 入库
	if err := u.database.WithContext(*ctx).Save(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (u *UserServiceImpl) QueryUser(ctx *context.Context, in *user.QueryUserRequest) (*user.UserSet, error) {
	// 先创建实例
	set := user.CreateNewUserSet()

	// 初始化db对象，搞query
	query := u.database.Model(user.User{}).WithContext(*ctx)

	// 构造query 相当于加 Where 子句
	if in.Username != "" {
		query = query.Where("username=?", in.Username)
	}

	// 查询，在查询之前，需要先统计一下count，否则后面query会失效
	if err := query.Count(&set.Total).Error; err != nil {
		log.Println("查询total失败")
		return nil, err
	}

	// 查询返回userset,相当于加了个limit子句
	if err := query.Offset(in.OffSet()).Limit(in.PageSize).Find(&set.Items).Error; err != nil {
		log.Println("查询userset失败")
		return nil, err
	}
	return set, nil
}

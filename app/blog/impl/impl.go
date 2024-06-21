package impl

import (
	"log"
	"vblog/app/blog"
	"vblog/common/config"
	"vblog/ioc"

	"gorm.io/gorm"
)

func init() {
	ioc.ControllerImpl.Register(blog.AppName, &Blogimpl{})
}

type Blogimpl struct {
	*gorm.DB
}

func (i *Blogimpl) Init() error {
	var err error
	i.DB, err = config.ReadDBConf(config.Filename).GetConn()
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	yaml "gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 首先得想好要读什么玩意
// 这里读取预计用yaml相关的库

func Default() *Config {
	return &Config{
		Database: &Database{
			Addr:      "127.0.0.1",
			Port:      3306,
			Password:  "123456",
			Username:  "admin",
			DBname:    "vblog",
			Charset:   "utf8mb4",
			ParseTime: "True",
			Loc:       "Local",
		},
		Application: &Application{
			Host:   "127.0.0.1",
			Port:   8080,
			Domain: "http://127.0.0.1",
		},
	}
}

type Config struct {
	// Addr      string `yaml:"addr"`
	// Port      int    `yaml:"port"`
	// Password  string `yaml:"password"`
	// Username  string `yaml:"username"`
	// DBname    string `yaml:"dbname"`
	// Charset   string `yaml:"charset"`
	// ParseTime string `yaml:"parsetime"`
	// Loc       string `yaml:"loc"`
	*Database    `yaml:"mysql"`
	*Application `yaml:"app"`
}

type Application struct {
	Host   string `toml:"host" yaml:"host" json:"host"`
	Port   int    `toml:"port" yaml:"port" json:"port"`
	Domain string `toml:"domain" yaml:"domain" json:"domain"`

	server *gin.Engine
	lock   sync.Mutex
	root   gin.IRouter
}

func (a *Application) Address() string {
	return fmt.Sprintf("%s:%d", a.Host, a.Port)
}

func (a *Application) GinServer() *gin.Engine {
	a.lock.Lock()
	defer a.lock.Unlock()

	if a.server == nil {
		a.server = gin.Default()
	}

	return a.server
}

func (a *Application) GinRootRouter() gin.IRouter {
	r := a.GinServer()

	if a.root == nil {
		a.root = r.Group("vblog").Group("api").Group("v1")
	}

	return a.root
}

func (a *Application) Start() error {
	r := a.GinServer()
	return r.Run(a.Address())
}

func ReadDBConf(file string) (c *Config) {
	// fp, err := os.Open(file)
	// if err != nil {
	// 	log.Printf("Config file %v is not exists, reading from env", file)
	// }
	// var content []byte
	// fp.Read(content)
	content, err := os.ReadFile(file)
	if err != nil {
		log.Printf("Config file %v is not exists, reading from env", file)
	}

	// decoder := yaml.NewDecoder()
	err = yaml.Unmarshal([]byte(content), &c)
	if err != nil {
		log.Println(err)
	}
	return c
}

func LoadFromYaml() error {
	c := ReadDBConf(Filename)
	if c == nil {
		return fmt.Errorf("read config error")
	}
	return nil
}

func (c *Config) DSN() string {
	connstr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=%v&parseTime=%v&loc=%v",
		c.Username, c.Password, c.Addr, c.Database.Port, c.DBname, c.Charset, c.ParseTime, c.Loc)
	return connstr
}

func (d *Config) GetConn() (*gorm.DB, error) {
	d.Database.lock.Lock()
	defer d.Database.lock.Unlock()

	if d.db != nil {
		return d.db, nil
	}
	dsn := d.DSN()
	var err error
	d.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Get db conn failed:\n", err)
		return nil, err
	}
	return d.db, nil
}

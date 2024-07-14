package config_test

import (
	"fmt"
	"log"
	"testing"
	"vblog/common/config"

	yaml "gopkg.in/yaml.v3"
)

var (
	config_file = "/Users/zephyrzhao/Documents/vblog/myblog/etc/db.yaml"
)

func TestConfig(t *testing.T) {
	d := config.ReadDBConf(config_file)
	log.Printf("Config: %#v", d)
	log.Println(d.Database)
}

func TestCreateDB(t *testing.T) {
	d := config.ReadDBConf(config_file)
	c, _ := yaml.Marshal(d)
	fmt.Println(string(c))
}

func TestToString(t *testing.T) {
	// d := db.CreateNewDBimpl()
	d := config.ReadDBConf(config_file)
	fmt.Println(d.DSN())
}

func TestGetConn(t *testing.T) {
	c := config.ReadDBConf(config_file)
	fmt.Println(c)
	fmt.Println(c.GetConn())
}

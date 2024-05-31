package config_test

import (
	"fmt"
	"log"
	"testing"
	"vblog/common/config"
	"vblog/common/db"

	yaml "gopkg.in/yaml.v3"
)

func TestConfig(t *testing.T) {
	d := config.ReadDBConf("/Users/zhaozephyr/Documents/vblog/common/config/db.yaml")
	log.Printf("Config: %#v", d)
}

func TestCreateDB(t *testing.T) {
	d := db.CreateNewDBimpl()
	c, _ := yaml.Marshal(d)
	fmt.Println(string(c))
}

func TestToString(t *testing.T) {
	d := db.CreateNewDBimpl()
	fmt.Println(d.DSN())
}

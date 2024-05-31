package config

import (
	"log"
	"os"
	"vblog/common/db"

	yaml "gopkg.in/yaml.v3"
)

// 首先得想好要读什么玩意
// 这里读取预计用yaml相关的库

func ReadDBConf(file string) (d *db.Database) {
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
	err = yaml.Unmarshal([]byte(content), &d)
	if err != nil {
		log.Println(err)
	}
	return d
}

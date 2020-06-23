package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type MysqlYaml struct {
	Mysql `yaml:"mysql"`
}

type Mysql struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
}

func MysqlConfig() MysqlYaml {
	var mysqlYaml MysqlYaml
	file, err := ioutil.ReadFile("config/config.yml")
	if err != nil {
		log.Printf("yaml file get err: #%v", err)
	}
	yaml.Unmarshal(file, &mysqlYaml)
	return mysqlYaml
}

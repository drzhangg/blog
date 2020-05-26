package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var Cfg *Yaml

type Yaml struct {
	Pgsql struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Dbname   string `yaml:"dbname"`
	}
}

func Config() *Yaml {
	var pgSqlConf = new(Yaml)
	file, err := ioutil.ReadFile("config/conf.yml")
	if err != nil {
		log.Printf("yaml file get err: #%v", err)
	}
	yaml.Unmarshal(file, &pgSqlConf)
	return pgSqlConf
}

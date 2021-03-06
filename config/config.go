package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var Cfg YamlConfig

type YamlConfig struct {
	Pgsql `yaml:"pgsql"`
}

type Pgsql struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
}

func Config() YamlConfig {
	var yamlConfig YamlConfig
	file, err := ioutil.ReadFile("config/config.yml")
	if err != nil {
		log.Printf("yaml file get err: #%v", err)
	}
	yaml.Unmarshal(file, &yamlConfig)
	return yamlConfig
}

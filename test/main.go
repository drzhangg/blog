package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Test struct {
	Pgsql `yaml:"pgsql"`
}

type Pgsql struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Dbname   string `yaml:"dbname"`
}

func main() {
	f, err := ioutil.ReadFile("../config/config.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(f))

	var p Test
	err = yaml.Unmarshal(f, &p)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("p:", p)
}

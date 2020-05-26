package common

import (
	"blog/config"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

var GetEngine *gorm.DB

var DataBase = "host=%s port=%s user=%s dbname=%s password=%s"

func GetDb() {
	DataBase = fmt.Sprintf(DataBase, config.Config().Pgsql.Host, config.Config().Pgsql.Port, config.Config().Pgsql.Dbname, config.Config().Pgsql.Password)
	db, err := gorm.Open("postgres", DataBase)
	if err != nil {
		log.Println("gorm open failed, err:", err)
	}
	GetEngine = db
}

func init() {
	go GetDb()
}

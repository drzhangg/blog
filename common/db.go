package commonpackage

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var Engine *gorm.DB

var DataBase = "host=%s port=%s user=%s dbname=%s password=%s"

func GetDb() {
	DataBase = fmt.Sprintf(DataBase)
	gorm.Open("postgres", DataBase)

}

func init() {
	go GetDb()
}

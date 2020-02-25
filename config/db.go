package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var Db *gorm.DB

func InitDb() (err error) {
	connStr := fmt.Sprintf("user=%v password=%v dbname=%v sslmode=disable", "postgres", "root", "go-web")
	if Db, err = gorm.Open("postgres", connStr); err != nil {
		log.Fatal("connect pgsql failed, err:", err)
		return err
	}
	//defer Db.Close()

	fmt.Println("connect pgsql success")
	Db.SingularTable(true)
	return
}

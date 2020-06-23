package common

import (
	"blog/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)


//var dataBase = "user=%s  password=%s host=%s port=%d  dbname=%s  sslmode=disable"
var dataBase = "postgres://%s:%s@%s:%d/%s?sslmode=disable"

//user:password@/dbname?charset=utf8&parseTime=True&loc=Local
var mysqlDatabase = "%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local"

func GetPgDb() *gorm.DB {
	dataBase = fmt.Sprintf(dataBase, config.Config().Pgsql.User, config.Config().Pgsql.Password, config.Config().Pgsql.Host, config.Config().Pgsql.Port, config.Config().Dbname)
	//fmt.Println("database:",dataBase)
	db, err := gorm.Open("postgres", dataBase)
	if err != nil {
		log.Println("gorm open failed, err:", err)
	}
	return db
}

func GetMqDb() *gorm.DB {
	user := config.MysqlConfig().User
	pwd := config.MysqlConfig().Password
	host := config.MysqlConfig().Host
	port := config.MysqlConfig().Port
	dbname := config.MysqlConfig().Dbname
	mysqlDatabase = fmt.Sprintf(mysqlDatabase, user, pwd, host, port, dbname)
	fmt.Println(mysqlDatabase)
	db, err := gorm.Open("mysql", mysqlDatabase)
	if err != nil {
		log.Println("gorm open failed, err:", err)
	}

	db.SingularTable(true)
	return db
}

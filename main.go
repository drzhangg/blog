package main

import (
	"blog/config"
	"blog/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Article struct {
	Id         int    `db:"id"`
	Title      string `db:"title"`
	CreateTime string `db:"create_time"`
	Content    string `db:"content"`
	Author     string `db:"author"`
	ImageUrl   string `db:"imageurl"`
}

func main() {
	//初始化pgsql
	if err := config.InitDb(); err != nil {
		log.Fatal(err)
		return
	}

	var article Article
	config.Db.Where("id = ?",1).Find(&article)
	fmt.Println("article:",article)

	router := gin.Default()

	router.Static("/static/", "./static")
	router.LoadHTMLGlob("views/*")

	router.GET("/", controller.IndexHandler)
	router.Run(":8089")
}

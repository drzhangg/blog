package main

import (
	"blog/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("/static/", "./static")
	router.LoadHTMLGlob("views/*")

	router.GET("/", controller.IndexHandler)
	router.Run(":8089")
}

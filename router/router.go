package router

import (
	"blog/router/index"
	"github.com/gin-gonic/gin"
)

func Run() {

	r := gin.Default()
	r.Static("static","static/")
	r.StaticFile("static","static/")
	r.LoadHTMLGlob("static/view/*")

	r.GET("/", index.Index)

	r.Run(":8083")
}

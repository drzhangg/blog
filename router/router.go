package router

import (
	"blog/router/index"
	"github.com/gin-gonic/gin"
	"log"
)

func Run() {

	r := gin.Default()
	r.Static("static", "static/")
	r.StaticFile("static", "static/")
	r.LoadHTMLGlob("static/view/*")

	r.GET("/", index.Index)

	r.POST("/login", func(c *gin.Context) {
		var u = struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{}
		if err := c.ShouldBindJSON(&u); err != nil {
			log.Println(err)
			return
		}
		log.Println("userinfo:", u)
	})

	r.Run(":8081")
}

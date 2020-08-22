package router

import (
	"blog/common/middlerware"
	"blog/router/index"
	"github.com/gin-gonic/gin"
	"log"
)

func Run() {

	r := gin.Default()

	r.Static("/static/", "./static")
	//r.StaticFile("static", "static/")
	r.LoadHTMLGlob("views/*")

	r.GET("/", index.Index)

	v1 := r.Group("", middlerware.Cors())
	v1.POST("/login", func(c *gin.Context) {
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

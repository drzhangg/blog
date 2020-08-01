package router

import (
	"blog/common/cron"
	"blog/router/index"
	"github.com/gin-gonic/gin"
)

func Run() {

	r := gin.Default()
	r.LoadHTMLGlob("static/view/*")

	r.GET("/", index.Index)


	cron.InitCron()

	r.Run(":8083")
}

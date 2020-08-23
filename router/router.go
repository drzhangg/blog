package router

import (
	"blog/controller"
	"github.com/gin-gonic/gin"
)

func Run() {

	r := gin.Default()

	r.Static("/static/", "./static")
	r.LoadHTMLGlob("views/*")

	r.GET("/", controller.IndexHandle)
	r.GET("/category/", controller.CategoryList)
	r.GET("/article/new/", controller.NewArticle)
	r.POST("/article/submit/", controller.ArticleSubmit)
	r.GET("/article/detail/", controller.ArticleDetail)
	r.POST("/upload/file/", controller.UploadFile)
	r.GET("/leave/new/", controller.LeaveNew)
	r.GET("/about/me/", controller.AboutMe)
	r.POST("/comment/submit/", controller.CommentSubmit)
	r.POST("/leave/submit/", controller.LeaveSubmit)


	r.Run(":8081")
}

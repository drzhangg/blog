package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//var count int
	//common.GetEngine.Table("test").Count(&count)
	//fmt.Println(count)

	r := gin.Default()

	r.POST("/test", func(c *gin.Context) {
		n1 := c.Request.PostFormValue("name")
		fmt.Println("nnn:", n1)
		buf := make([]byte, 1024)
		n, _ := c.Request.Body.Read(buf)
		var data = make(map[string]interface{})
		json.Unmarshal(buf[:n], &data)

		c.JSON(http.StatusOK, gin.H{
			"name": data["name"],
			"age":  data["age"],
		})
	})

	r.Run(":8081")
}

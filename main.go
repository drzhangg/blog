package main

import (
	"blog/common"
	"blog/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//var count int
	//common.GetEngine.Table("test").Count(&count)
	//fmt.Println(count)

	var test = struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}{}
	//common.GetEngine.Table("test").Where("id = 1").First(&test)
	common.GetMqDb().Table("test").Where("id = 1").First(&test)
	fmt.Println(test)

	var t = model.Test{Name: "zhang"}
	err := common.GetMqDb().Table("test").Save(&t).Error
	if err != nil {
		fmt.Println("save err:",err)
		return
	}

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

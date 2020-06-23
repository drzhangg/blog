package main

import (
	"blog/common"
	"blog/model"
	"fmt"
)

func main() {
	var user []model.User

	common.GetMqDb().Find(&user)
	fmt.Println(user)

}

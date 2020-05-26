package main

import (
	"blog/common"
	"fmt"
)

func main() {
	var count int
	common.GetEngine.Table("test").Count(&count)
	fmt.Println(count)
}

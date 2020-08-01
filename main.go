package main

import (
	"blog/common/cron"
	"blog/router"
	"fmt"
)

func main() {
	cron.InitCron()
	router.Run()

	//cron to upload git
	go func() {
		fmt.Println(1)
	}()

}

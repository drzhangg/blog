package main

import (
	"blog/common/cron"
	"blog/router"
)

func main() {
	cron.InitCron()
	router.Run()

	//cron to upload git

}

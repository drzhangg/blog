package cron

import (
	"blog/common/cmd"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
)

func InitCron() {
	c := cron.New()
	fmt.Println("in------")

	//每分钟提交一次代码
	c.AddFunc("*/1 * * * ?", func() {
		err := cmd.PushGithub("update")
		if err != nil {
			log.Println(err)
			return
		}
	})
	c.Start()
}

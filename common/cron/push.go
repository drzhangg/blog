package cron

import (
	"blog/common/cmd"
	"github.com/robfig/cron/v3"
	"log"
)

func InitCron() {
	c := cron.New()

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

package main

import (
	"github.com/robfig/cron"
	"mygoproject/cmd/remindadddelaccount"
	"mygoproject/cmd/remindexpireaccount"
)

func main() {
	c := cron.New()
	//c.AddFunc("0 */1 * * * ?",remindexpireaccount.Remindaccount)
	//每天10点执行:账号即将过期提醒
	c.AddFunc("0 0 10 * * ?",remindexpireaccount.Remindaccount)
	//每天9点执行：新增删除账号信息
	c.AddFunc("0 0 9 * * ?",remindadddelaccount.Adddelaccount)

	go c.Start()
	select {

	}
}

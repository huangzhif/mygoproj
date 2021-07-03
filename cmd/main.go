package main

import (
	"github.com/robfig/cron"
	"mygoproject/cmd/remindexpireaccount"
)

func main() {
	c := cron.New()
	c.AddFunc("0 */1 * * * ?",remindexpireaccount.Remindaccount)
	//每天10点执行
	//c.AddFunc("0 0 10 * * ?",remindexpireaccount.Remindaccount)

	go c.Start()
	select {

	}
}

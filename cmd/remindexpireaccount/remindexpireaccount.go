/*
目的：用于发邮件提醒将于7天后到期的账户
创建者：huangzhifu
创建时间：20210623
*/
package remindexpireaccount

import (
	"mygoproject/internal/db"
	"mygoproject/internal/email"
	"mygoproject/internal/getconfig"
	"mygoproject/internal/logger"
	"strings"
)

func Remindaccount() {
	logger.Info.Println("开始执行")
	var ab db.Database
	ab = new(db.Mysql)
	db := ab.Init()
	bl, ok := ab.Query(db, "SELECT * FROM auth_user u LEFT JOIN t_profileUser p ON u.id = p.user_id WHERE u.is_active = 1 and datediff(p.disabledate,CURRENT_DATE ()) = 7")
	if ok {
		config := getconfig.InitConfigure()
		subject := config.Get("remindexpireaccount.subject").(string)
		temp := config.Get("remindexpireaccount.temp").(string)
		cc := config.Get("remindexpireaccount.cc").(string)
		ip := config.Get("normal.ip").(string)
		// 把 配置文件的抄送列表转化为切片
		var ccslice []string
		if cc == "" {
			ccslice = []string{}
		} else {
			ccslice = strings.Split(cc, ",")
		}

		for _, raw := range bl {

			stru := struct {
				Username    string
				Account     string
				Disabledate string
				IP          string
			}{Username: raw["last_name"],
				Account:     raw["username"],
				Disabledate: raw["disabledate"],
				IP:          ip,
			}

			err := email.SendEmail([]string{raw["email"]}, ccslice, subject, temp, &stru)
			if err != nil {
				logger.Error.Println("发送失败：", err, raw["email"])

			} else {
				logger.Info.Println("发送成功：", raw["email"])
			}

		}
	} else {
		logger.Info.Println("没有找到数据")
	}

	db.Close()

}

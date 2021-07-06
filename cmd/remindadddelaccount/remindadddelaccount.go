package remindadddelaccount

import (
	"bytes"
	"fmt"
	"html/template"
	"mygoproject/internal/db"
	"mygoproject/internal/email"
	"mygoproject/internal/getconfig"
	"mygoproject/internal/logger"
	"strings"
	"time"
)

const layout = "2006-01-02"

var istyle = map[string]string{"0": "<span style='color:red'>删除</span>", "1": "新增"}

func Adddelaccount() {
	t := time.Now()
	yesterday := t.AddDate(0, 0, -1).Format(layout)
	config := getconfig.InitConfigure()
	sql := config.Get("remindadddelaccount.sql").(string)
	var idb db.Database
	idb = new(db.Mysql)
	thisdb := idb.Init()

	ret, ok := idb.Query(thisdb, fmt.Sprintf(sql, yesterday))
	if ok {

		subject := config.Get("remindadddelaccount.subject").(string)
		temp := config.Get("remindadddelaccount.temp").(string)
		to := config.Get("remindadddelaccount.to").(string)
		cc := config.Get("remindadddelaccount.cc").(string)

		toslice := strings.Split(to, ",")

		var ccslice []string
		if cc == "" {
			ccslice = []string{}
		} else {
			ccslice = strings.Split(cc, ",")
		}

		var tbody bytes.Buffer
		for _, raw := range ret {
			tbody.WriteString("<tr style='height:30px'>")
			tbody.WriteString("<td style='border:1px solid #000000'>" + raw["source"] + "</td>")
			tbody.WriteString("<td style='border:1px solid #000000'>" + raw["mainpart"] + "</td>")
			tbody.WriteString("<td style='border:1px solid #000000'>" + raw["name"] + "</td>")
			tbody.WriteString("<td style='border:1px solid #000000'>" + raw["username"] + "</td>")
			tbody.WriteString("<td style='border:1px solid #000000'>" + istyle[raw["style"]] + "</td>")
			tbody.WriteString("</tr>")
		}

		stru := struct {
			Tbody interface{}
		}{
			Tbody: template.HTML(tbody.String()),
		}

		err := email.SendEmail(toslice, ccslice, subject, temp, &stru)
		if err != nil {
			logger.Error.Println("发送失败：", err)

		} else {
			logger.Info.Println("发送成功：",toslice)
		}
	} else {
		logger.Info.Println("没有找到数据")
	}

	defer thisdb.Close()
}

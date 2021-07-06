package email

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"html/template"
	"mygoproject/internal/getconfig"
	"net/smtp"
)

func SendEmail(to []string, cc []string, subject string, temp string, stru interface{}) error {
	/*
		参数信息：
		to : 发送对象
	    cc : 抄送对象
		subject : 邮件主题
		temp : 邮件模板地址
		stru : 结构体 stru := struct {
				Username 	string
				Account  	string
				Disabledate string
			}{	Username: "username",
				Account: "account",
				Disabledate: "disabledate",
			}
	*/
	/*获取基本信息*/
	config := getconfig.InitConfigure()
	user := config.Get("email.user").(string)
	pwd := config.Get("email.pwd").(string)
	host := config.Get("email.host").(string)
	from := config.Get("email.from").(string)
	port := config.Get("email.port").(string)

	e := email.NewEmail()
	e.From = from
	e.To = to
	e.Cc = cc
	e.Subject = subject
	t, err := template.ParseFiles(temp)
	if err != nil {
		panic(err)
	}

	body := new(bytes.Buffer)
	t.Execute(body, &stru)
	e.HTML = body.Bytes()
	er := e.SendWithTLS(fmt.Sprintf("%v:%v", host, port), smtp.PlainAuth("", user, pwd, host), &tls.Config{ServerName: host})
	return er
}

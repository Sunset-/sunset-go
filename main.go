package main

import (
	"fmt"
	"github.com/Sunset-/sunset-go/components/mails"
	"io/ioutil"
	"time"
)

var mailTemplate = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
    </head>
    <body>
        <h4>亲爱的{{.ToUserName}},您好！</h4>
        <div>{{.Message}}</div>
        </br>
        <div>
            {{.FromUserName}} </br>
            {{.TimeDate}}
        </div>
    </body>
</html>
`

func main() {

	f, _ := ioutil.ReadFile("./节点视频链路流程.pptx")
	err := mails.SendMail(mails.MailAuth{
		Host:        "smtp.qq.com",
		Smtp:        "smtp.qq.com:587",
		MailAccount: "770203139@qq.com",
		Password:    "jjjcixqzkjsebdgh",
	},
		"770203139@qq.com", "385970211@qq.com", "三封测试邮件", mailTemplate, struct {
			FromUserName string
			ToUserName   string
			TimeDate     string
			Message      string
		}{
			FromUserName: "go语言",
			ToUserName:   "志洋",
			TimeDate:     time.Now().Format("2006/01/02"),
			Message:      "golang是世界上最好的语言！",
		}, []*mails.MailAttach{
			&mails.MailAttach{
				FileName: "节点视频链路流程.pptx",
				Body:     f,
			},
		})
	fmt.Println(err)
}

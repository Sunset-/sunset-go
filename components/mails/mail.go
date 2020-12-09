package mails

import (
	"bytes"
	"github.com/jordan-wright/email"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"html/template"
	"io/ioutil"
	"mime"
	"net/smtp"
	"path/filepath"
)

type MailAuth struct {
	Host        string
	Smtp        string
	MailAccount string
	Password    string
}
type MailAttach struct {
	FileName string
	Body     []byte
}

func SendMail(auth MailAuth, fromUser, toUser, subject, tpl string, bodyParams interface{}, files []*MailAttach) error {
	// NewEmail返回一个email结构体的指针
	e := email.NewEmail()
	// 发件人
	e.From = fromUser
	// 收件人(可以有多个)
	e.To = []string{toUser}
	// 邮件主题
	e.Subject = subject
	// 解析html模板
	tp := template.New("temp")
	t, err := tp.Parse(tpl)
	if err != nil {
		return err
	}
	// Buffer是一个实现了读写方法的可变大小的字节缓冲
	body := new(bytes.Buffer)
	// Execute方法将解析好的模板应用到匿名结构体上，并将输出写入body中
	err = t.Execute(body, bodyParams)
	if err != nil {
		return err
	}
	// html形式的消息
	e.HTML = body.Bytes()
	// 以路径将文件作为附件添加到邮件中
	if len(files) > 0 {
		for _, f := range files {
			fname, err := ioutil.ReadAll(transform.NewReader(bytes.NewReader([]byte(f.FileName)), simplifiedchinese.GBK.NewEncoder()))
			if err != nil {
				return err
			}
			ct := mime.TypeByExtension(filepath.Ext(f.FileName))
			_, err = e.Attach(bytes.NewReader(f.Body), string(fname), ct)
			if err != nil {
				return err
			}
		}
	}
	// 发送邮件(如果使用QQ邮箱发送邮件的话，passwd不是邮箱密码而是授权码)
	return e.Send(auth.Smtp, smtp.PlainAuth("", auth.MailAccount, auth.Password, auth.Host))
	//return e.Send("smtp.qq.com:587", smtp.PlainAuth("", "xxx@qq.com", "passwd", "smtp.qq.com"))
}

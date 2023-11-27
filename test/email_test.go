package test

import (
	"github.com/go-gomail/gomail"
	"testing"
)

func TestSendEmail(t *testing.T) {
	//e := email.NewEmail()
	//e.From = "Get <3551906947@qq.com>"
	//e.To = []string{"255429862@qq.com"}
	////e.Bcc = []string{"test_bcc@example.com"}
	////e.Cc = []string{"test_cc@example.com"}
	//e.Subject = "Awesome Subject"
	//e.Text = []byte("Text Body is, of course, supported!")
	//e.HTML = []byte("<b>123456</b>")
	////err := e.Send("smtp.163.com:456", smtp.PlainAuth("", "getcharzhaopan@163.com", "XYQHBVSIUAXXRHQ", "smtp.163.com"))
	//err := e.SendWithTLS("8.130.86.26:456", smtp.PlainAuth("", "3551906947@qq.com", "123456789zwm", "smtpdm.aliyun.com"), &tls.Config{InsecureSkipVerify: true, ServerName: "smtpdm.aliyun.com"})
	//if err != nil {
	//	t.Fatal(err)
	//}
	// 发件人邮箱地址和授权码
	from := "3551906947@qq.com"
	password := "zxltqxejwbuqdbgh"

	// 收件人邮箱地址
	to := "2946142439@qq.com"

	// Create a new message
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Test Subject")
	m.SetBody("text/plain", "刘柳彤你在干嘛")
	// Create a new mailer
	d := gomail.NewDialer("smtp.qq.com", 587, from, password)
	// Send the email
	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	}

	println("Email sent successfully!")
}

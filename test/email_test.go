package test

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func TestEmail(t *testing.T) {
	e := email.NewEmail()
	e.From = "six feet under <2237616362@qq.com>"
	e.To = []string{"fqhslycopene@gmail.com"}
	e.Subject = "验证码发送测试"
	e.HTML = []byte("您的验证码是：<b>123456</b>")
	err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", "2237616362@qq.com", "kcsvfkkzdzyuecdb", "smtp.qq.com"))
	if err != nil {
		fmt.Println(err)
	}
}

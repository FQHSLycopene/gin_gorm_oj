package utils

import (
	"github.com/jordan-wright/email"
	"net/smtp"
)

func SendCode(code, ToEmailstr string) error {
	e := email.NewEmail()
	e.From = "six feet under <2237616362@qq.com>"
	e.To = []string{ToEmailstr}
	e.Subject = "验证码发送"
	e.HTML = []byte("您的验证码是：<b>" + code + "</b>")
	err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", "2237616362@qq.com", "kcsvfkkzdzyuecdb", "smtp.qq.com"))
	return err
}

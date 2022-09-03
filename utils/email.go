package utils

import (
	"fmt"
	"github.com/jordan-wright/email"
	"math/rand"
	"net/smtp"
	"strings"
	"time"
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
func GetCode() string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < 6; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

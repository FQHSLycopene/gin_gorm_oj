package test

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"testing"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

//生成Token
func TestGenerateToken(t *testing.T) {
	mySigningKey := []byte("AllYourBase")
	c := MyClaims{
		Username: "qimiao",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, _ := token.SignedString(mySigningKey)
	fmt.Println(tokenString)
	fmt.Println("=========================================================")
	Token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	fmt.Println(Token.Claims.(*MyClaims))
}

//解析Token
func TestAnalyseToken(t *testing.T) {

}

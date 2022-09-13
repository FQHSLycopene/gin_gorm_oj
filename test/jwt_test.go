package test

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"testing"
)

type MyClaims struct {
	Username string `json:"username"`
	Identity string `json:"identity"`
	jwt.RegisteredClaims
}

//生成Token
func TestGenerateToken(t *testing.T) {
	mySigningKey := []byte("gin_gorm_oj")
	c := MyClaims{
		Username: "Lycopene",
		Identity: "a593a94e-f389-49e2-97eb-e8202abca57e",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	fmt.Println(token)
	tokenString, _ := token.SignedString(mySigningKey)
	fmt.Println(tokenString)
	fmt.Println("=========================================================")
	Token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("gin_gorm_oj"), nil
	})
	if err != nil {
		fmt.Printf("%s", err)
		return
	}
	fmt.Println(Token.Claims.(*MyClaims).Identity)
}

//解析Token
func TestAnalyseToken(t *testing.T) {

}

package utils

import (
	"crypto/md5"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
)

//生成MD5
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

var mySigningKey = []byte("gin_gorm_oj")

type MyClaims struct {
	jwt.RegisteredClaims
	name     string
	identity string
}

func GenerateToken(identity, name string) (string, error) {

	c := MyClaims{
		name:     name,
		identity: identity,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

//解析Token
func TestAnalyseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(*MyClaims), err
}

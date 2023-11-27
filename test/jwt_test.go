package test

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"testing"
)

type userClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
	Identity string `json:"identity"`
}

var myKey = []byte("gin_gorm_oj_key")

// 生成token
func TestGenerateToken(t *testing.T) {
	UserClaims := &userClaims{
		Identity:       "user_1",
		Name:           "Get",
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(tokenString)
}

// 解析token
func TestAnalyseToken(t *testing.T) {
	tokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJuYW1lIjoiR2V0IiwiaWRlbnRpdHkiOiJ1c2VyXzEifQ.vlEobDwJyJQlK9_4MXsviq-lddisowaP-qUIk1K6zks"
	UserClaims := new(userClaims)
	claims, err := jwt.ParseWithClaims(tokenString, UserClaims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if claims.Valid {
		fmt.Println(UserClaims)
	}
}

package helper

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-gomail/gomail"
)

type UserClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
	Identity string `json:"identity"`
}

var myKey = []byte("gin_gorm_oj_key")

// 生成token
func GenerateToken(identity, name string) (string, error) {
	UserClaims := &UserClaims{
		Identity:       identity,
		Name:           name,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析token
func AnalyseToken(tokenString string) (*UserClaims, error) {
	UserClaims := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, UserClaims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return UserClaims, nil
}

// 加密
func GetMd5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

// 发送验证码
func SendCode(toUserEmail, code string) error {
	// 发件人邮箱地址和授权码
	from := "3551906947@qq.com"
	password := "zxltqxejwbuqdbgh"

	// 收件人邮箱地址
	to := toUserEmail

	// Create a new message
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Test Subject")
	m.SetBody("text/plain", "您的验证码是:"+code)
	// Create a new mailer
	d := gomail.NewDialer("smtp.qq.com", 587, from, password)
	// Send the email
	err := d.DialAndSend(m)
	return err
}

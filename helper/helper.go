package helper

import (
	"crypto/md5"
	"fmt"
	"gin_gorm_oj/define"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-gomail/gomail"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type UserClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
	IsAdmin  int    `json:"is_admin"`
	Identity string `json:"identity"`
}

var myKey = []byte("gin_gorm_oj_key")

// 生成token
func GenerateToken(identity, name string, is_admin int) (string, error) {
	UserClaims := &UserClaims{
		Identity:       identity,
		Name:           name,
		IsAdmin:        is_admin,
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
func GetUUID() string {
	return uuid.NewV4().String()
}
func GetRand() string {
	rand.Seed(time.Now().UnixNano())
	rand.Intn(10)
	s := ""
	for i := 0; i < 6; i++ {
		s += strconv.Itoa(rand.Intn(10))
	}
	return s
}

// code save
func CodeSave(code []byte) (string, error) {
	dirName := "code/" + GetUUID()
	path := dirName + "/main.go"
	err := os.Mkdir(dirName, 0777)
	if err != nil {
		return "", err
	}
	f, err := os.Create(path)
	if err != nil {
		return "", err
	}
	f.Write(code)
	defer f.Close()
	return path, nil
}

// CheckGoCodeValid
// 检查golang代码的合法性
func CheckGoCodeValid(path string) (bool, error) {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return false, err
	}
	code := string(b)
	for i := 0; i < len(code)-6; i++ {
		if code[i:i+6] == "import" {
			var flag byte
			for i = i + 7; i < len(code); i++ {
				if code[i] == ' ' {
					continue
				}
				flag = code[i]
				break
			}
			if flag == '(' {
				for i = i + 1; i < len(code); i++ {
					if code[i] == ')' {
						break
					}
					if code[i] == '"' {
						t := ""
						for i = i + 1; i < len(code); i++ {
							if code[i] == '"' {
								break
							}
							t += string(code[i])
						}
						if _, ok := define.ValidGolangPackageMap[t]; !ok {
							return false, nil
						}
					}
				}
			} else if flag == '"' {
				t := ""
				for i = i + 1; i < len(code); i++ {
					if code[i] == '"' {
						break
					}
					t += string(code[i])
				}
				if _, ok := define.ValidGolangPackageMap[t]; !ok {
					return false, nil
				}
			}
		}
	}
	return true, nil
}

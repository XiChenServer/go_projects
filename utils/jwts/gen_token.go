package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
	"time"
	"virus/global"
)

func GenToken(user JwtPayLoad) (string, error) {
	MySecret = []byte(global.Config.Jwy.Secret)
	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.Jwy.Expires))),
			Issuer:    global.Config.Jwy.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}

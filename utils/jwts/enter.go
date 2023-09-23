package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
)

type JwtPayLoad struct {
	//Username string `json:"username"`
	NickName string `json:"nickName"`
	Role     int    `json:"role"`
	UserID   uint   `json:"user_id"`
}

var MySecret []byte

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}

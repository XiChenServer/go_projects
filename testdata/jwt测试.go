package main

import (
	"fmt"
	"virus/core"
	"virus/global"
	"virus/utils/jwts"
)

func main() {
	core.InitConf()
	global.Log = core.InitLogger()

	fmt.Println(global.Config.Jwy.Secret)
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   1,
		Role:     1,
		Username: "少年",
		NickName: "dfggg",
	})
	fmt.Println(token, err)
	claime, err := jwts.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IuWwkeW5tCIsIm5pY2tOYW1lIjoiZGZnZ2ciLCJyb2xlIjoxLCJ1c2VyX2lkIjoxLCJleHAiOjE2OTU0NTczNTUuMTAwNTU4LCJpc3MiOiJ4eCJ9.6n4q55aGKUu0MeG5pzj4A5HWfl2k7A_Dol8uv1HO4XQ")
	fmt.Println(claime, err)
}

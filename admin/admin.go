package main

import (
	"context"
	"flag"
	"fmt"
	"iot-platform/user/rpc/types/user"
	"net/http"

	"iot-platform/admin/internal/config"
	"iot-platform/admin/internal/handler"
	"iot-platform/admin/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/admin-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("token") == "" {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("111Unauthorized"))
				return
			}
			auth, err := ctx.RpcUser.Auth(context.Background(), &user.UserAuthRequest{Token: r.Header.Get("token")})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized"))
				return
			}
			ctx.AuthUser = auth
			next(w, r)
		}
	})
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}

package main

import (
	"flag"
	"fmt"
	"iot-platform/device/internal/mqtt"

	"iot-platform/device/internal/config"
	"iot-platform/device/internal/server"
	"iot-platform/device/internal/svc"
	"iot-platform/device/types/device"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/device.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		device.RegisterDeviceServer(grpcServer, server.NewDeviceServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()
	go mqtt.NewMqttServer(c.Mqtt.Broker)
	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

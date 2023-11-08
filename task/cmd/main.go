package main

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"grpc-todolist-tmp/user/config"
	"grpc-todolist-tmp/user/discovery"
	"grpc-todolist-tmp/user/internal/handler"
	"grpc-todolist-tmp/user/internal/repository"
	"grpc-todolist-tmp/user/internal/service"
	"net"
)

func main() {
	config.InitConfig()
	repository.InitDB()
	//etcd path addr
	etcdAddress := []string{viper.GetString("etcd.address")}
	//服务注册
	etcdRegister := discovery.NewRegister(etcdAddress, logrus.New())
	grocAddress := viper.GetString("server.grpcAddress")
	userNode := discovery.Server{
		Name: viper.GetString("service"),
		Addr: grocAddress,
	}
	server := grpc.NewServer()
	defer server.Stop()
	service.RegisterUserServiceServer(server, handler.NewUserService())
	lis, err := net.Listen("tcp", grocAddress)
	if err != nil {
		panic(err)
	}
	if _, err = etcdRegister.Register(userNode, 10); err != nil {
		panic(err)
	}
	if err = server.Serve(lis); err != nil {
		panic(err)
	}
}

package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	pb "hello-protobuf/proto/hello" // 引入编译生成的包
	"net"
	_ "net"
)

const (
	// Address gRPC服务地址
	Address = "localhost:50052"
)

// 定义helloService并实现约定的接口
type helloService struct{}

// HelloService Hello服务
var HelloService = helloService{}

// SayHello 实现Hello服务接口
func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)

	return resp, nil
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		println("报错:", err)
	}
	creds, err := credentials.NewServerTLSFromFile("C:/Users/admin/go/src/hello-protobuf/keys/server.pem", "C:/Users/admin/go/src/hello-protobuf/keys/server.key")
	if err != nil {
		println("服务器tls报错:", err)
	}
	// 实例化grpc Server
	s := grpc.NewServer(grpc.Creds(creds))

	// 注册HelloService
	pb.RegisterHelloServer(s, HelloService)

	s.Serve(listen)
}

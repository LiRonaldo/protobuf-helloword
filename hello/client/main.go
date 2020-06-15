package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	pb "hello-protobuf/proto/hello" // 引入proto包
)

const (
	// Address gRPC服务地址
	Address = "localhost:50052"
)

func main() {
	//tls 链接
	creds, err := credentials.NewClientTLSFromFile("C:/Users/admin/go/src/hello-protobuf/keys/server.pem", "WWW.BAIDU.COM")
	if err != nil {
		println("客户端tls错误：", err)
	}

	// 连接
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds))
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloClient(conn)

	// 调用方法
	req := &pb.HelloRequest{Name: "gRPC"}
	res, err := c.SayHello(context.Background(), req)

	if err != nil {
		grpclog.Fatalln(err)
	}

	println(res.Message)
}

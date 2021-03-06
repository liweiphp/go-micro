package main

import (
	"context"
	goods "d10/grpc"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

const (
	Addr = "192.168.56.102:9002"
)

func main() {
	cres,err := credentials.NewClientTLSFromFile("../keys/server.pem","www.123.com")
	if err != nil {
		grpclog.Fatalln("crenditial error:", err)
	}
	conn,err := grpc.Dial(Addr, grpc.WithTransportCredentials(cres))
	//conn, err := grpc.Dial(Addr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("conn error:", err)
	}
	defer conn.Close()
	//初始化客户端
	c := goods.NewGoodsServiceClient(conn)

	req := &goods.Goods{Id: 1}
	//调用server方法
	res, err := c.Get(context.Background(), req)
	if err != nil {
		grpclog.Fatalln("get func error:", err)
	}
	fmt.Println(res)
}

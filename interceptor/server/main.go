package main

import (
	"context"
	goods "d10/grpc"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"net"
)

const Addr = "192.168.56.102:9002"
type GoodsService struct {}

var GoodsServices = GoodsService{}

func (g GoodsService)Get(ctx context.Context, args *goods.Goods) (*goods.Goods, error)  {
	//Auth(ctx)
	fmt.Println("args:", args)
	replyGoods := &goods.Goods{
		Id : 1,
		Name : "test",
	}
	return replyGoods,nil
}

func main() {
	listen,err := net.Listen("tcp", Addr)
	if err != nil {
		fmt.Println("error:", err)
	}
	defer listen.Close()

	var serverOpt []grpc.ServerOption
	cres,err := credentials.NewServerTLSFromFile("../keys/server.pem", "../keys/server.key")
	if err != nil {
		grpclog.Fatal("credential error:", err)
	}
	serverOpt = append(serverOpt, grpc.Creds(cres))
	serverOpt = append(serverOpt, grpc.UnaryInterceptor(interceptor))
	s := grpc.NewServer(serverOpt...)
	//s := grpc.NewServer()

	goods.RegisterGoodsServiceServer(s, GoodsServices)

	fmt.Println("listen on :"+Addr)
	s.Serve(listen)

}

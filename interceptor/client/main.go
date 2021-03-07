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
	OpenSSL = true
)

//自定义验证器
type customCreads struct {}

//自定义认证接口
func (c customCreads)GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error){
	return map[string]string{
		"Id":"1",
		"token":"wwwaaa",
	},nil
}
//是否开启openssl验证
func (c customCreads)RequireTransportSecurity() bool  {
	return OpenSSL
}



func main() {
	var err error
	var opts []grpc.DialOption
	if OpenSSL {
		//开启openssl
		cres,err := credentials.NewClientTLSFromFile("../keys/server.pem","www.123.com")
		if err != nil {
			grpclog.Fatalln("crenditial error:", err)
		}
		//证书验证
		opts = append(opts, grpc.WithTransportCredentials(cres))
	}else {
		//正常验证
		opts = append(opts, grpc.WithInsecure())
	}
	//自定义验证
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCreads)))
	opts = append(opts, grpc.WithUnaryInterceptor(interceptor))
	conn, err := grpc.Dial(Addr, opts...)
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

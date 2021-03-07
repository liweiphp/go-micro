package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Println("客户端拦截器")
	//客户端调用
	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}

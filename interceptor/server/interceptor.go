package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

func interceptor(ctx context.Context, req interface{},info *grpc.UnaryServerInfo,handler grpc.UnaryHandler) (interface{},error) {
	Auth(ctx)
	fmt.Println("拦截器")
	return handler(ctx,req)
}

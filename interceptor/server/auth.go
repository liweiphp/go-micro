package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/metadata"
)

func Auth(ctx context.Context) (bool,error) {
	md,ok := metadata.FromIncomingContext(ctx)
	if !ok {
		fmt.Println("无token信息")
		return false,errors.New("无token信息")
	}
	fmt.Println("token:",md["token"],"id:",md["id"])
	return true,nil
}

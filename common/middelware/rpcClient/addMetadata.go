package rpcclient

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type UserInfo struct {
	UserId  string
	GroupId string
}

// rpc 客户端拦截器
func ClientAddMetadataInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	fmt.Println("=========客户端拦截器开始：==========")
	fmt.Printf("req: %v\n", req)
	fmt.Printf("method: %s\n", method)

	ctx = AddMd(ctx)
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err != nil {
		return err
	}

	fmt.Println("=========客户端拦截器 结束：==========")
	return nil
}

func AddMd(ctx context.Context) context.Context {
	ui := UserInfo{
		UserId:  "124",
		GroupId: "23",
	}

	data, _ := json.Marshal(ui)
	// 添加元数据
	md := metadata.New(map[string]string{"Grpc-Metadata-GroupId": "23", "Grpc-Metadata-UserInfo": string(data)})
	// 按顺序添加
	if ctx.Value("UserId") != nil {
		md.Append("UserId", strconv.Itoa(ctx.Value("UserId").(int)))
	}

	if ctx.Value("UserRole") != nil {
		md.Append("UserRole", ctx.Value("UserRole").(string))
	}
	ctx = metadata.NewOutgoingContext(ctx, md)
	return ctx
}

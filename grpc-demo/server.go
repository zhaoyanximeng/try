package main

import (
	"context"
	"fmt"
	"github.com/casbin/casbin/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"grpc-demo/src/pbfiles"
	"grpc-demo/src/services"
	"log"
	"net"
)

/*var AuthMap map[string]string

func init() {
	AuthMap = make(map[string]string)
	AuthMap["admin"] = "/ProdService/GetProd"
}*/

func checkToken(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// 获取grpc请求携带的元数据信息
	if md, ok := metadata.FromIncomingContext(ctx); !ok {
		return nil, status.Error(codes.Unauthenticated, "metadata error")
	} else {
		if tokens, ok := md["token"]; ok {
			fmt.Println(tokens[0]) // token验证
		} else {
			return nil, status.Error(codes.Unauthenticated, "token error")
		}
	}

	return handler(ctx, req)
}

var E *casbin.Enforcer

func init() {
	e,err := casbin.NewEnforcer("casbin/model.conf","casbin/p.csv")
	if err != nil {
		log.Fatal(err)
	}
	E = e
}

// 一般先获得token，再通过token获得用户角色，验证权限
func RABC(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unavailable, "请求错误")
	}
	tokens := md.Get("token")
	if len(tokens) != 1 {
		return nil,status.Errorf(codes.Unauthenticated,"没有权限（token）")
	}
	b,err := E.Enforce(tokens[0],info.FullMethod)
	if !b || err != nil {
		return nil,status.Errorf(codes.Unauthenticated,"没有权限")
	}

	return handler(ctx,req)
}

func main() {
	//grpc.UnaryInterceptor()拦截器
	s := grpc.NewServer(grpc.UnaryInterceptor(RABC))
	pbfiles.RegisterProdServiceServer(s, services.NewProdService())

	listen, _ := net.Listen("tcp", ":8080")
	if err := s.Serve(listen); err != nil {
		log.Fatal(err)
	}
}

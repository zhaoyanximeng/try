package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"grpc-demo/src/lib"
	"grpc-demo/src/pbfiles"
	"log"
)

func main() {
	conn, err := grpc.DialContext(context.Background(),
		"localhost:8080",
		grpc.WithInsecure(),
		grpc.WithPerRPCCredentials(lib.NewAuth("guanliyuan")), //设置身份凭证
	)
	if err != nil {
		log.Fatal(err)
	}
	client := pbfiles.NewProdServiceClient(conn)
	resp, err := client.UpdateProd(context.Background(), &pbfiles.ProdRequest{Id: 120})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.String())
}

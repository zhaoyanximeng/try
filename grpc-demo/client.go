package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-demo/src/pbfiles"
	"io"
	"log"
	"time"
)

func main() {
	conn, err := grpc.DialContext(context.Background(),
		"localhost:8080",
		grpc.WithInsecure(),
		//grpc.WithPerRPCCredentials(lib.NewAuth("guanliyuan")), //设置身份凭证
	)
	if err != nil {
		log.Fatal(err)
	}
	/*client := pbfiles.NewProdServiceClient(conn)
	resp, err := client.UpdateProd(context.Background(), &pbfiles.ProdRequest{Id: 120})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.String())*/

	client := pbfiles.NewUserServiceClient(conn)
	stream,err := client.GetUserScoreByStream(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	id := 0
	for id < 10 {
		time.Sleep(time.Second * 1)

		id += 1
		req := &pbfiles.UserScoreRequest{
			UserId:               int32(id),
		}
		err = stream.Send(req)
		if err != nil {
			log.Fatal(err)
		}

		res,err := stream.Recv()
		if err != nil && err != io.EOF {
			log.Fatal(err)
		} else if err == io.EOF {
			break
		}

		log.Printf("id:%d,score:%d",res.UserInfo.UserId,res.UserInfo.UserScore)
	}

	err = stream.CloseSend()
	if err != nil {
		log.Fatal(err)
	}
}

package services

import (
	"grpc-demo/src/pbfiles"
	"io"
	"log"
	"math/rand"
)

type UserService struct {
	pbfiles.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (U *UserService) GetUserScoreByStream(stream pbfiles.UserService_GetUserScoreByStreamServer) error {
	rand.Seed(100)
	for {
		req,err := stream.Recv()
		if err != nil && err != io.EOF {
			log.Fatal(err)
		} else if err == io.EOF {
			break
		}

		score := rand.Int31n(100)
		err = stream.Send(&pbfiles.UserScoreResponse{
			UserInfo:             &pbfiles.UserInfo{
				UserId:               req.UserId,
				UserScore:            score,
			},
		})
		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

package services

import (
	"context"
	"grpc-demo/src/pbfiles"
	"log"
)

type ProdService struct {
	pbfiles.UnimplementedProdServiceServer
}

func NewProdService() *ProdService {
	return &ProdService{}
}

func (p *ProdService) GetProd(c context.Context, req *pbfiles.ProdRequest) (*pbfiles.ProdResponse, error) {
	log.Println("rpc请求进来了")

	if err := req.Validate(); err != nil {
		return nil, err
	}

	rsp := &pbfiles.ProdResponse{
		Result: &pbfiles.ProdModel{Id: req.Id, Name: "test"},
	}

	return rsp, nil
}

func (p *ProdService) UpdateProd(c context.Context, req *pbfiles.ProdRequest) (*pbfiles.ProdResponse, error) {
	return &pbfiles.ProdResponse{
		Result:               &pbfiles.ProdModel{
			Id:                   req.Id,
			Name:                 "success",
		},
	}, nil
}

func (p *ProdService) GetProdStream(*pbfiles.ProdRequest, pbfiles.ProdService_GetProdStreamServer) error {
	return nil
}

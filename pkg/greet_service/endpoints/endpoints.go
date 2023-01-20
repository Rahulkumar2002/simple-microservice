package endpoints

import (
	"context"
	"errors"
	"log"

	"github.com/Rahulkumar2002/simple-microservice/api/pb"
	"github.com/Rahulkumar2002/simple-microservice/pkg/greet_service"
	"github.com/go-kit/kit/endpoint"
)

type Set struct {
	GreetEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc greet_service.Service) Set {
	return Set{
		GreetEndpoint: MakeGreetEndpoint(svc),
	}
}

func MakeGreetEndpoint(svc greet_service.Service) endpoint.Endpoint {
	log.Printf("Inside MakeGreetEndpoint")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GreetRequest)

		resp, err := svc.GreetName(ctx, &pb.Name{Name: req.Name})
		if err != nil {
			return GreetResponse{resp.Message, err.Error()}, nil
		}

		return GreetResponse{resp.Message, ""}, nil
	}
}

func (s *Set) GreetName(ctx context.Context, name string) (string, error) {
	log.Printf("Inside Greet Set!!!")
	resp, err := s.GreetEndpoint(ctx, GreetRequest{Name: name})
	if err != nil {
		return "", err
	}
	getResp := resp.(GreetResponse)
	if getResp.Err != "" {
		return "", errors.New(getResp.Err)
	}
	return getResp.Message, nil
}

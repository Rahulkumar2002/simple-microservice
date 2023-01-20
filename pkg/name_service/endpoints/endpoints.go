package endpoints

import (
	"context"
	"errors"
	"log"

	"github.com/Rahulkumar2002/simple-microservice/pkg/name_service"
	"github.com/go-kit/kit/endpoint"
)

type Set struct {
	GiveNameEndpoint endpoint.Endpoint
}

func NewEndpointSet(svc name_service.Service) Set {
	return Set{
		GiveNameEndpoint: MakeGiveNameEndpoint(svc),
	}
}

func MakeGiveNameEndpoint(svc name_service.Service) endpoint.Endpoint {
	log.Printf("Inside MakeGiveNameEndpoint!!!")
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GiveNameRequest)
		message, err := svc.GiveName(ctx, req.Name)
		if err != nil {
			return GiveNameResponse{message, err.Error()}, nil
		}

		return GiveNameResponse{message, ""}, nil
	}
}

func (s *Set) GiveName(ctx context.Context, name string) (string, error) {
	log.Printf("Inside GiveName Set!!!")
	resp, err := s.GiveNameEndpoint(ctx, GiveNameRequest{Name: name})
	if err != nil {
		return "", err
	}
	getResp := resp.(GiveNameResponse)
	if getResp.Err != "" {
		return "", errors.New(getResp.Err)
	}
	return getResp.Message, nil
}

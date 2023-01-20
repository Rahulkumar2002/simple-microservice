package greet_service

import (
	"context"

	"github.com/Rahulkumar2002/simple-microservice/api/pb"
)

type Service interface {
	GreetName(ctx context.Context, name *pb.Name) (*pb.GiveReply, error)
}

package greet_service

import (
	"context"
	"log"

	"github.com/Rahulkumar2002/simple-microservice/api/pb"
)

type greetService struct{}

func NewGreetService() Service {
	return &greetService{}
}

func (g *greetService) GreetName(_ context.Context, name *pb.Name) (*pb.GiveReply, error) {
	log.Println("Inside the Greet microservice!!!")

	return &pb.GiveReply{Message: "Hello " + name.Name, Err: ""}, nil

}

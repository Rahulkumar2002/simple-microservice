package name_service

import (
	"context"
	"log"

	pb "github.com/Rahulkumar2002/simple-microservice/api/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	Client pb.GreetServiceClient
)

func createConnection() (pb.GreetServiceClient, *grpc.ClientConn) {
	conn, err := grpc.Dial("localhost:8082", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Did not connect=%v", err)
	}

	return pb.NewGreetServiceClient(conn), conn
}

type nameService struct{}

func NewNameService() Service {
	return &nameService{}
}

func (n *nameService) GiveName(_ context.Context, name string) (string, error) {
	// store the name in the sql database and send it to the greet service:
	log.Println("Inside the Name microservice!!!")
	Client, conn := createConnection()
	defer conn.Close()
	log.Printf("Client=%v", Client)
	resp, err := CallGreetName(Client, name)
	if err != nil {
		return "", err
	}
	return resp, nil
}

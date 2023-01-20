package name_service

import (
	"context"
	"log"

	pb "github.com/Rahulkumar2002/simple-microservice/api/pb"
	"go.mongodb.org/mongo-driver/bson"
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
	log.Println("Inside the Name microservice!!!")
	mongoClient := DBConnection()
	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	Client, conn := createConnection()
	defer conn.Close()

	coll := mongoClient.Database("names").Collection("name")
	result, err := coll.InsertOne(context.TODO(), bson.D{
		{Key: "name", Value: name},
	})
	if err != nil {
		return "", err
	}
	log.Printf("Result of MongoDB insertion: %v", result)

	resp, err := CallGreetName(Client, name)
	if err != nil {
		return "", err
	}

	return resp, nil
}

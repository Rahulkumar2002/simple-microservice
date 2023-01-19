package main

import (
	"log"

	pb "github.com/Rahulkumar2002/simple-microservice/api/pb/name_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	port   = ":8081"
	Client pb.NameServiceClient
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect:  %v", err)
	}
	defer conn.Close()

	Client = pb.NewNameServiceClient(conn)
}

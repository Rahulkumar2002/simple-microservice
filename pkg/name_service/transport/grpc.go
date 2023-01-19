package transport

import (
	"context"
	"log"
	"time"

	pb "github.com/Rahulkumar2002/simple-microservice/api/pb/name_service"
)

func CallGreetName(client pb.NameServiceClient, name *pb.Name) string {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	log.Printf("GreetName is Called")
	res, err := client.GreetName(ctx, name)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	return res.Message
}

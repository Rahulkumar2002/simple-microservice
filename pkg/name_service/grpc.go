package name_service

import (
	"context"
	"errors"
	"log"
	"time"

	pb "github.com/Rahulkumar2002/simple-microservice/api/pb"
)

func CallGreetName(client pb.GreetServiceClient, name string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	log.Printf("GreetName RPC is Called")
	res, err := client.GreetName(ctx, &pb.Name{Name: name})
	log.Printf("Response : %v", res)
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
		return "", err
	}
	if res.Err != "" {
		log.Fatalf("Error in creating response: %v", res.Err)
		return res.Message, errors.New(res.Err)
	}
	return res.Message, nil
}

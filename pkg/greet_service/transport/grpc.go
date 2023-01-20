package transport

import (
	"context"

	pb "github.com/Rahulkumar2002/simple-microservice/api/pb"
	"github.com/Rahulkumar2002/simple-microservice/pkg/greet_service/endpoints"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	pb.UnimplementedGreetServiceServer
	greetName grpctransport.Handler
}

func NewGRPCServer(ep endpoints.Set) pb.GreetServiceServer {
	return &grpcServer{
		greetName: grpctransport.NewServer(
			ep.GreetEndpoint,
			decodeGRPCGreetRequest,
			decodeGRPCGreetResponse,
		),
	}
}

func (g *grpcServer) GreetName(ctx context.Context, r *pb.Name) (*pb.GiveReply, error) {
	_, rep, err := g.greetName.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return rep.(*pb.GiveReply), nil
}

func decodeGRPCGreetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*pb.Name)
	return endpoints.GreetRequest{Name: req.Name}, nil
}

func decodeGRPCGreetResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(endpoints.GreetResponse)
	return &pb.GiveReply{Message: reply.Message, Err: reply.Err}, nil
}

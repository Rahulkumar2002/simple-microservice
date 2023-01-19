package name_service

import (
	"context"
	"os"

	pb "github.com/Rahulkumar2002/simple-microservice/api/pb/name_service"
	mainService "github.com/Rahulkumar2002/simple-microservice/cmd/name_service"
	transport "github.com/Rahulkumar2002/simple-microservice/pkg/name_service/transport"
	"github.com/go-kit/log"
)

type nameService struct{}

func NewNameService() Service {
	return &nameService{}
}

func (n *nameService) GiveName(_ context.Context, name string) (string, error) {
	// store the name in the sql database and send it to the greet service:
	logger.Log("Inside the Name microservice!!!")
	pbName := &pb.Name{
		Name: name,
	}
	resp := transport.CallGreetName(mainService.Client, pbName)
	return resp, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}

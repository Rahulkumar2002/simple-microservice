package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Rahulkumar2002/simple-microservice/pkg/name_service"
	"github.com/Rahulkumar2002/simple-microservice/pkg/name_service/endpoints"
	"github.com/Rahulkumar2002/simple-microservice/pkg/name_service/transport"
	"github.com/go-kit/log"
	"github.com/oklog/oklog/pkg/group"
)

var (
	defaultHTTPPort = "8081"
)

func main() {
	var (
		logger   log.Logger
		httpAddr = net.JoinHostPort("localhost", defaultHTTPPort)
	)

	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)

	var (
		service     = name_service.NewNameService()
		eps         = endpoints.NewEndpointSet(service)
		httpHandler = transport.NewHTTPHandler(eps)
	)

	var g group.Group
	{
		httpListener, err := net.Listen("tcp", httpAddr)
		if err != nil {
			logger.Log("transport", "HTTP", "during", "Listen", "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Log("transport", "HTTP", "addr", httpAddr)
			return http.Serve(httpListener, httpHandler)
		}, func(error) {
			logger.Log("transport", "HTTP", "addr", "closing")
			httpListener.Close()
		})
	}
	{
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-c:
				return fmt.Errorf("received signal %s", sig)
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			close(cancelInterrupt)
		})
	}

	logger.Log("exit", g.Run())
}

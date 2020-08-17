package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	pb "github.com/David-solly/gokit-microservice-starter/pkg/api/v1"
	microservice_grpc "github.com/David-solly/gokit-microservice-starter/pkg/api/v1/service"
	"github.com/go-kit/kit/log"
	"google.golang.org/grpc"
)

func main() {

	// Logging domain.
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}

	errChan := make(chan error)

	// parse variable from input command
	var (
		advertiseAddr = flag.String("advertise.addr", "", "advertise address")
		advertisePort = flag.String("advertise.port", "", "advertise port")
	)
	flag.Parse()

	// Blank port will default to run the service on 8080
	// If not containerized - check for port conflict
	//
	if advertisePort == nil {
		port := "8080"
		advertisePort = &port
	}

	// For local Testing or other server endpoints and port - not docker specific
	// here is an example of how that would work
	// var (
	// 	gRPCAddr = flag.String("grpc", "127.0.0.1:"+"8082",
	// 		"gRPC listen address")
	// )
	var (
		gRPCAddr = flag.String("grpc", ""+*advertiseAddr+":"+*advertisePort,
			"gRPC listen address")
	)

	flag.Parse()

	ctx := context.Background()

	// init service
	var svc microservice_grpc.ServiceInterface
	svc = microservice_grpc.MicroService{}

	// creating Endpoints struct
	endpoints := microservice_grpc.ServiceEndpoints{
		ServiceEndpoint: microservice_grpc.MakeServiceFunctionEndpoint(svc),
	}

	//execute grpc server
	go func() {
		listener, err := net.Listen("tcp", *gRPCAddr)
		if err != nil {
			errChan <- err
			return
		}

		handler := microservice_grpc.NewGRPCServer(ctx, endpoints)
		gRPCServer := grpc.NewServer()
		pb.RegisterMicroServiceServer(gRPCServer, handler)
		fmt.Printf("Service info %v", gRPCServer.GetServiceInfo())
		errChan <- gRPCServer.Serve(listener)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	error := <-errChan
	logger.Log("err", error)
	panic(error)

}

// go run . -advertise.addr "127.0.0.1" -advertise.port "8080"
// commandline args to use when running server

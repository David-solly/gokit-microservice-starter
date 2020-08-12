package service

import (
	"context"

	pb "github.com/David-solly/gokit-microservice-starter/pkg/api/v1"
	grpctransport "github.com/go-kit/kit/transport/grpc"
)

type grpcServer struct {
	Service grpctransport.Handler
}

// implement Server Interface in <micro_service.pb.go>
func (s *grpcServer) CallService(ctx context.Context, r *pb.ServiceRequest) (*pb.ServiceResponse, error) {
	_, resp, err := s.Service.ServeGRPC(ctx, r)
	if err != nil {
		return nil, err
	}
	return resp.(*pb.ServiceResponse), nil
}

// NewGRPCServer :
// create new grpc server
func NewGRPCServer(ctx context.Context, endpoint ServiceEndpoints) pb.MicroServiceServer {
	return &grpcServer{
		Service: grpctransport.NewServer(
			endpoint.ServiceEndpoint,
			DecodeGRPCServiceRequest,
			EncodeGRPCServiceResponse,
		),
	}
}

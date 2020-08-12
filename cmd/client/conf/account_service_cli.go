package conf

import (
	pb "github.com/David-solly/gokit-microservice-starter/pkg/api/v1"
	account_grpc "github.com/David-solly/gokit-microservice-starter/pkg/api/v1/service"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

// NewService :
// Call the service method through the onion layer of a service
func NewService(conn *grpc.ClientConn) account_grpc.ServiceInterface {
	var ServiceEndpoint = grpctransport.NewClient(
		conn, "v1.MicroService", "CallService",
		account_grpc.EncodeGRPCServiceRequest,
		account_grpc.DecodeGRPCServiceResponse,
		pb.ServiceResponse{},
	).Endpoint()

	return account_grpc.ServiceEndpoints{
		ServiceEndpoint: ServiceEndpoint,
	}
}

package conf

import (
	pb "github.com/David-solly/gokit-microservice-starter/pkg/api/v1"
	service "github.com/David-solly/gokit-microservice-starter/pkg/api/v1/service"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"google.golang.org/grpc"
)

// NewService :
// Call the service method through the onion layer of a service
func NewService(conn *grpc.ClientConn) service.ServiceInterface {
	var ServiceEndpoint = grpctransport.NewClient(
		conn, "v1.MicroService", "CallService",
		service.EncodeGRPCServiceRequest,
		service.DecodeGRPCServiceResponse,
		pb.ServiceResponse{},
	).Endpoint()

	return service.ServiceEndpoints{
		ServiceEndpoint: ServiceEndpoint,
	}
}

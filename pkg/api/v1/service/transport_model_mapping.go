package service

import (
	"context"

	pb "github.com/David-solly/gokit-microservice-starter/pkg/api/v1"
	models "github.com/David-solly/gokit-microservice-starter/pkg/api/v1/models"
)

//Encode and Decode Service Request
func DecodeGRPCServiceRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.ServiceRequest)
	return models.ServiceRequest{Request: req.Request}, nil
}

//Encode and Decode Service Request
func EncodeGRPCServiceRequest(_ context.Context, r interface{}) (interface{}, error) {
	n := r.(models.ServiceRequest)

	return &pb.ServiceRequest{Request: n.Request}, nil
}

// Encode and Decode Request Response
func EncodeGRPCServiceResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(models.ServiceResponse)
	err := pb.ServiceError{}
	if resp.Error != nil {
		err.Code = int32(resp.Error.Code)
		err.Error = resp.Error.Error
	}

	return &pb.ServiceResponse{
		Error: &err,
		Info: &pb.BusinessResponse{
			Data: resp.Payload.Data,
		}}, nil
}

func DecodeGRPCServiceResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(*pb.ServiceResponse)
	rs := models.ServiceResponse{
		Payload: models.BusinessLogicDataPayload{
			Data: resp.Info.Data,
		},
	}

	return rs, nil

}

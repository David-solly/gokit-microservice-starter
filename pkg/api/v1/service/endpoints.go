package service

import (
	"context"

	models "github.com/David-solly/gokit-microservice-starter/pkg/api/v1/models"
	"github.com/go-kit/kit/endpoint"
)

type ServiceEndpoints struct {
	ServiceEndpoint endpoint.Endpoint
}

func MakeServiceFunctionEndpoint(svc ServiceInterface) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(models.ServiceRequest)

		data, err := svc.CallService(ctx, req)

		if err != nil {
			return nil, err
		}
		return models.ServiceResponse{Payload: *data}, nil
	}
}

func (te ServiceEndpoints) CallService(ctx context.Context, req models.ServiceRequest) (*models.BusinessLogicDataPayload, error) {
	resp, err := te.ServiceEndpoint(ctx, req)
	if err != nil {
		return nil, err
	}

	businessLogicResponse := resp.(models.ServiceResponse)
	return &businessLogicResponse.Payload, nil

}

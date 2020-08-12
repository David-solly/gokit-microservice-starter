package service

import (
	"context"

	models "github.com/David-solly/gokit-microservice-starter/pkg/api/v1/models"
)

// ServiceInterface :
// Service deifinition to be implemented by the server and transport layers
type ServiceInterface interface {
	CallService(ctx context.Context, request models.ServiceRequest) (*models.BusinessLogicDataPayload, error)
}

// MicroService :
// This implements the `ServiceInterface` above
// Used as a handle for the microservice's Business logic
//
type MicroService struct {
}

// CallService :
// FRAGILE Signature WARNING - changing this Function name will require multiple edits.
// This is the actual business logic call that the server will execute
// when the rpc method is invoked
// Any business logic should ideally be handled by this method
// This should be considered the entry point to your server business logic
//
func (accs MicroService) CallService(ctx context.Context, request models.ServiceRequest) (*models.BusinessLogicDataPayload, error) {
	return performBusinessOperation(request.Request) //Do business logic
}

//Business logic 1
func performBusinessOperation(data string) (a *models.BusinessLogicDataPayload, err1 error) {
	//classic String reversal
	//
	rv := []byte(data)
	acci, err := callBusinessFunction(rv)
	return acci, err
}

//Business logic 2
func callBusinessFunction(data []byte) (account *models.BusinessLogicDataPayload, err error) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return &models.BusinessLogicDataPayload{Data: string(data)}, nil
}

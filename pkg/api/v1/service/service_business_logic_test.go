package service

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	pb "github.com/David-solly/gokit-microservice-starter/pkg/api/v1"
	models "github.com/David-solly/gokit-microservice-starter/pkg/api/v1/models"
	"github.com/docker/docker/pkg/testutil/assert"
)

func TestDataTransferIntegration(t *testing.T) {
	microSvc := MicroService{}
	validInput := models.ServiceRequest{}

	t.Run("Ensure Correct object Type assertion ", func(t *testing.T) {
		businessResponse, err := microSvc.CallService(context.Background(), validInput)
		assert.NilError(t, err)
		assert.Equal(t, reflect.TypeOf(businessResponse), reflect.TypeOf(&models.BusinessLogicDataPayload{}))
	})

}

// ############  ----- User Defined Tests -------------   #########
// Basic testing
func TestMyService(t *testing.T) {
	microSvc := MicroService{}

	for _, v := range []struct {
		input, output string
	}{
		{"hello world", "dlrow olleh"},
		{"dlrow olleh", "hello world"},
		{"hi", "ih"},
		{"i", "i"},
	} {
		validInput := models.ServiceRequest{Request: v.input}

		t.Run("Test Data output", func(t *testing.T) {
			businessResponse, err := microSvc.CallService(context.Background(), validInput)
			assert.NilError(t, err)
			assert.NotNil(t, businessResponse)
			assert.Equal(t, reflect.TypeOf(businessResponse), reflect.TypeOf(&models.BusinessLogicDataPayload{}))
			js, err := json.Marshal(businessResponse)
			rsp := string(js)
			assert.NilError(t, err)
			assert.Contains(t, rsp, v.output)
			fmt.Printf("%v", rsp)
		})
	}

}

//Test Domain data cohesion when passing through the service
func TestMyDataMapping(t *testing.T) {
	incomingRequest := models.ServiceRequest{Request: "Data to test"}
	outgoingResponse := models.ServiceResponse{Payload: models.BusinessLogicDataPayload{"Test from function"}}

	t.Run("Ensure Correct object Type assertion transfer TO domain ", func(t *testing.T) {
		requestAsGRPCRequest, err := EncodeGRPCServiceRequest(context.Background(), incomingRequest)
		assert.NilError(t, err)
		assert.Equal(t, reflect.TypeOf(requestAsGRPCRequest), reflect.TypeOf(&pb.ServiceRequest{Request: incomingRequest.Request}))

		requestAsDomainRequest, err := DecodeGRPCServiceRequest(context.Background(), requestAsGRPCRequest)
		assert.NilError(t, err)
		assert.Equal(t, reflect.TypeOf(requestAsDomainRequest), reflect.TypeOf(models.ServiceRequest{Request: incomingRequest.Request}))

	})

	t.Run("Ensure Correct object Type assertion transfer FROM domain ", func(t *testing.T) {
		responseAsGRPCResponse, err := EncodeGRPCServiceResponse(context.Background(), outgoingResponse)
		assert.NilError(t, err)
		assert.Equal(t, reflect.TypeOf(responseAsGRPCResponse), reflect.TypeOf(&pb.ServiceResponse{Info: &pb.BusinessResponse{Data: outgoingResponse.Payload.Data}}))

		responseAsDomainResponse, err := DecodeGRPCServiceResponse(context.Background(), responseAsGRPCResponse)
		assert.NilError(t, err)
		assert.Equal(t, reflect.TypeOf(responseAsDomainResponse), reflect.TypeOf(models.ServiceResponse{Payload: outgoingResponse.Payload}))

	})

}

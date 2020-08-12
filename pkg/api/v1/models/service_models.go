package models

// BusinessLogicDataPayload :
// The business logic function response payload
type BusinessLogicDataPayload struct {
	Data string `json:"data,omitempty"` // edit fields
}

// ServiceRequest :
// The business logic function request
type ServiceRequest struct {
	Request string `json:"request"` //edit fields
}

// ServiceResponse :
// The response from the business logic layer that gets returned to the
// calling client
type ServiceResponse struct {
	Payload BusinessLogicDataPayload `json:"payload,omitempty"`
	Error   *ServiceError            `json:"error,omitempty"`
}

type ServiceError struct {
	Error string `json:"error,omitempty"`
	Code  int    `json:"code,omitempty"`
}

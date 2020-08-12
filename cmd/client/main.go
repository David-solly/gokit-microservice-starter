package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	conf "github.com/David-solly/gokit-microservice-starter/cmd/client/conf"
	models "github.com/David-solly/gokit-microservice-starter/pkg/api/v1/models"
	service_grpc "github.com/David-solly/gokit-microservice-starter/pkg/api/v1/service"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

// Sample naive main method - Rest
// This client can be wrapped beind a http rest call and
// invoke the method
func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", greetHandler).Methods("Get")
	router.HandleFunc("/reverse", serviceHandler).Methods("POST")
	log.Fatal(http.ListenAndServe(":8082", router))

}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(models.ServiceResponse{Payload: models.BusinessLogicDataPayload{Data: "API is running"}})

}

// Rest endpoint Handler
//
func serviceHandler(w http.ResponseWriter, r *http.Request) {
	req := models.ServiceRequest{}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.ServiceError{Error: "Error parsing request body"})
		return
	}

	err = json.Unmarshal(reqBody, &req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.ServiceError{Error: "Error parsing request data"})
		return
	}

	// set the grpc server address and port
	// Can be passed via os variables too or cmd line args
	// var (
	// 	grpcAddr = flag.String("addr.server", "127.0.0.1:8080",
	// 		"gRPC address")
	// )
	// flag.Parse()

	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure(), //Should not use in production
		grpc.WithTimeout(5*time.Second)) // timeout of service

	if err != nil {
		log.Fatalln("gRPC dial:", err)
	}
	// Close gRPC connection with server after single call or timeout
	defer conn.Close()

	//Create grpc connection
	microService := conf.NewService(conn)

	// Call gRPC method
	data, err := callMicroServiceEndpoint(context.Background(), microService, req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(models.ServiceError{Error: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(data)
}

func callMicroServiceEndpoint(ctx context.Context, svc service_grpc.ServiceInterface, req models.ServiceRequest) (interface{}, error) {
	resp, err := svc.CallService(ctx, req)
	return resp, err
}

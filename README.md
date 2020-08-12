# gokit-microservice-starter
A REST to gRPC client server starter for a single rpc method. MVP boiler-plate code. includes Docker build file and pulls a minimal alpine build without a shell within the Docker container.

This was an attempt to make life slightly easier for micro service development using go-kit and gGPC server-client ipc. Unfortunately the only options for using go-kit were to use generators or dig into the code base.
I hope this will help lighten the load somewhat for those looking to start using go-kit without being put-off golang because of the shear number of lines of code required for a simple function.

The Idea was to make a template that could be expanded on while making the minimum amount of changes.

The template is a working example that reverses a string

This microservice can be edited by changing just one file and implementing your business logic there
  - the ```service_business_logic.go``` file
          located in ```pkg/api/v1/service/service_business_logic.go```
  - TDD is a great practice and advisable in genreal but especially when dealing with go-kit or any gRPC communication, so it is strongly encouraged especially here so be sure to edit the corresponding test file

# Running the Code
In order to run the code, both server and client are needed though can be run independantly
the server code and run command can be found at 
  - ```cmd/server/main.go``` - runs on port 8080

and the client - naive basic implementation of a REST endpoint interfacing with a gRPC server
  - ```cmd/client/main.go``` - runs on port 8082

The Rest endpoint has 2 routes
  - GET: /
  - POST: /reverse 
  
The latter calls the gRPC server and passes the JSON request body ```{"request":"String data to process"}``` through the server and reverses it.

  The entire project is customisable just beware that the more files you edit, the more changes you will need to propagate to the corresponding files.
  
  Please feel free to post pull requests and improve as needed.
  
  Happier Coding.

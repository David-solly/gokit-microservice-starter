# gokit-microservice-starter
A REST to gRPC client server starter for a single rpc method. Minimal boiler-plate code based, structured using the popular [starter-kit][https://github.com/golang-standards/project-layout].
### Container Size
This starter kit includes a Docker multi-stage build file for the server and pulls a minimal alpine build without a shell to run within the Docker container. The resulting container is tiny and fully deployable. Your docker environment will cache the larger stages though so you may easily get docker bloat of a few hundred MB while building this image.
### debugging
It means you will get logs printed to the terminal but attempts to run commands inside the shell will fail with an error. If running inside kubernetes, 
you will see the logs printed but you will be unable to execute any commands to the container with exec, good for security - not so good for container level debugging.

### Inspiration for kit
This was an attempt to make life slightly easier for micro service development using go-kit and gRPC server-client ipc.
I hope this will help lighten the load somewhat for those looking to start using go-kit without being put-off golang because of the sheer number of lines of code required for a simple function, also reading some negative experiences with what is otherwise a fantastic piece of kit with a steep barrier to entry for the users.

The Idea was to make a template that could be expanded on while making the minimum amount of changes.

The template is a work in progress and provides a working example that reverses a string

This microservice can be edited by changing just one file and implementing your business logic there
  - the ```service_business_logic.go``` file
          located in ```pkg/api/v1/service/service_business_logic.go```
  - TDD is a great practice and advisable in general but especially when dealing with go-kit or any gRPC communication, so it is strongly encouraged especially here so be sure to edit the corresponding test file. Basic Tests are provided to ensure the data mapping between gokit and your business logic are met, but are by no means a complete set of tests. The unit test for the exported Service function is provided as an example.

# Running the Code
To run the code, both server and client are needed, though can be run independently and the client doesn't have to be the one provided nor does it have to be in golang.
the server code and run command can be found at 
  - ```cmd/server/main.go``` - runs on port 8080
  - eg from the working directory containing the server's `main.go` file. To run on port 8080 and bind to localhost ```go run . -advertise.addr "127.0.0.1" -advertise.port "8080"```

and the client - naive basic implementation of a REST endpoint interfacing with a gRPC server
  - ```cmd/client/main.go``` - runs on port 8082
The method endpoint handlers for the REST client are in this file too.

The Rest endpoint has 2 routes
  - GET: /
  - POST: /reverse 
  
The latter calls the gRPC server and passes the JSON request body ```{"request":"String data to process"}``` through the server and reverses it.

  The entire project is customisable just beware that the more files you edit, the more changes you will need to propagate to the corresponding boilerplate files.
  
  Please feel free to post pull requests and improve as needed.
  
  Happier Coding.

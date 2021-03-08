genG:
	  protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative greet/greetpb/*.proto
clean:
	rm greet/greetpb/*.go
runs:
	go run greet/greet_server/server.go
runc:
	go run greet/greet_client/client.go
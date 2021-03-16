genG:
	  protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative api/grpc/protos/*.proto
clean:
	rm api/grpc/protobuffs/*.go
runs:
	go run cmd/grpc-server/server.go
runc:
	go run cmd/grpc-client/client.go
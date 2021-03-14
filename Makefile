genG:
	  protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative grpc/news_pb/*.proto
clean:
	rm grpc/news_pb/*.go
runs:
	go run grpc/news_server/server.go
runc:
	go run grpc/news_client/client.go
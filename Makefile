compile_proto:
	cd ./proto; protoc -I=. --go_out=../pb --go_opt=paths=source_relative --go-grpc_out=../pb --go-grpc_opt=paths=source_relative *.proto;

start_server:
	go run ./cmd/server/server.go -port 8080

start_client:
	go run ./cmd/client/client.go -port 8080 -grpcType sStream
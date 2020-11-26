compile_proto:
	cd ./proto; protoc -I=. --go_out=../pb --go_opt=paths=source_relative *.proto;

start_unary_server:
	go run ./cmd/server/unary_server.go -port 8080

start_unary_client:
	go run ./cmd/client/unary_client.go -port 8080
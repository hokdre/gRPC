compile_proto:
	cd ./proto; protoc -I=. --go_out=../pb --go_opt=paths=source_relative *.proto;
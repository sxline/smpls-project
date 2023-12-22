run:
	go run cmd/api/main.go

test:
	go test -v ./...

proto:
	protoc --proto_path=internal/pb/ internal/pb/*.proto --go_out=../ --go-grpc_out=require_unimplemented_servers=false:../ --experimental_allow_proto3_optional
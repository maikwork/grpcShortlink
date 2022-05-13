.PHONY: run, build 

go_grpc: 
	protoc --go_out=./internal/api/ --go-grpc_out=./internal/api/ --go-grpc_opt=require_unimplemented_servers=false api.proto

run: go_grpc
	@go run ./cmd/run/main.go

build: go_grpc
	@go build ./cmd/run/main.go

.PHONY: run, build 

DB=postgresql

go_grpc: 
	protoc --go_out=/internal/api/grpc --go-grpc_out=/internal/api/grpc api.proto

run: go_grpc
	@go run ./cmd/run/main.go $(DB)

build: go_grpc
	@go build ./cmd/run/main.go
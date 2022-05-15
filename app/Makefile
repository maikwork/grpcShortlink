.PHONY: run, build 

NAME_PROTO = api.proto
PATH_APP = ./app
PATH_PB = ./pb
PATH_SERVER = ./cmd/server/main.go

go_grpc: 
	protoc --go_out=$(PATH_PB) \
			--go-grpc_out=$(PATH_PB) \
			--go-grpc_opt require_unimplemented_servers=false \
			--grpc-gateway_out=$(PATH_PB) \
			--grpc-gateway_opt generate_unbound_methods=true \
			--openapiv2_out=. \
			$(NAME_PROTO)

run: go_grpc
	@go run $(PATH_SERVER)

build: go_grpc
	@go build $(PATH_SERVER)
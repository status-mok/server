generate: buf-generate go-generate

go-generate:
	@go generate ./...

buf-format:
	@buf format -w

buf-lint:
	@buf lint

buf-generate: buf-lint buf-format
	@buf generate

install-deps: buf-deps go-deps

go-deps:
	@go install github.com/vektra/mockery/v2
	@go install github.com/go-swagger/go-swagger/cmd/swagger

buf-deps:
	@go install github.com/bufbuild/buf/cmd/buf
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2

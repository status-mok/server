generate: buf-generate go-generate

go-generate:
	@go generate ./...

buf-format:
	@buf format -w

buf-lint:
	@buf lint

buf-generate: buf-lint buf-format
	@buf generate

install-deps: buf-deps

buf-deps:
	@go install github.com/bufbuild/buf/cmd/buf@v1.10.0
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.14.0
	@go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.14.0

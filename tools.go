//go:build tools

package main

import (
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "github.com/vektra/mockery/v2"
)

//go:build tools

package tools

import (
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/google/gnostic"
	_ "github.com/googleapis/gnostic-grpc"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "github.com/mgechev/revive"
	_ "google.golang.org/grpc"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)

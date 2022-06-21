//go:build tools
// +build tools

package tools

import (
	_ "ariga.io/ogent"
	_ "entgo.io/contrib/entproto/cmd/protoc-gen-entgrpc"
	_ "entgo.io/ent/cmd/ent"
	_ "github.com/hedwigz/entviz"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)

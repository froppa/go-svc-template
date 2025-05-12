// Code generated for boilerplate. DO NOT EDIT.
package proto

import (
  "context"
  "google.golang.org/grpc"
)

// HelloRequest is a dummy request.
type HelloRequest struct {
  Name string
}

// HelloReply is a dummy reply.
type HelloReply struct {
  Message string
}

// GreeterServer is the server API.
type GreeterServer interface {
  Say(context.Context, *HelloRequest) (*HelloReply, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct{}

func (UnimplementedGreeterServer) Say(context.Context, *HelloRequest) (*HelloReply, error) {
  return nil, nil
}

// RegisterGreeterServer registers the server.
func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
  // no-op; real code comes from generated file
}

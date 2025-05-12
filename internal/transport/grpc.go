package transport

import (
	"context"

	"github.com/froppa/go-svc-template/internal/service"
	pb "github.com/froppa/go-svc-template/proto"
	"google.golang.org/grpc"
)

// greeterServer implements the pb.GreeterServer interface.
type greeterServer struct {
	pb.UnimplementedGreeterServer
	svc service.Service
}

// NewGRPCServer constructs a GreeterServer using the Service.
func NewGRPCServer(svc service.Service) pb.GreeterServer {
	return &greeterServer{svc: svc}
}

// Say implements the Greeter.Say RPC.
func (g *greeterServer) Say(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	msg, err := g.svc.Hello(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	return &pb.HelloReply{Message: msg}, nil
}

// RegisterGRPCServices registers all your gRPC services on the provided server.
func RegisterGRPCServices(server *grpc.Server, svc service.Service) {
	pb.RegisterGreeterServer(server, NewGRPCServer(svc))
}

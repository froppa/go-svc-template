package grpcfx

import (
	"context"
	"fmt"
	"net"

	"github.com/froppa/go-svc-template/internal/service"
	"github.com/froppa/go-svc-template/internal/transport"
	"github.com/froppa/go-svc-template/pkg/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func GRPCPort(cfg *config.Config) int { return cfg.Server.GRPCPort }

var Module = fx.Options(
	fx.Provide(GRPCPort),
	fx.Invoke(NewGRPCServer),
)

func NewGRPCServer(
	lc fx.Lifecycle,
	port int,
	svc service.Service,
	log *zap.Logger,
) {
	addr := fmt.Sprintf(":%d", port)
	lis, _ := net.Listen("tcp", addr)
	server := grpc.NewServer()
	transport.RegisterGRPCServices(server, svc)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go server.Serve(lis)
			log.Info("gRPC listening", zap.Int("port", port))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("gRPC shutdown", zap.Int("port", port))
			server.GracefulStop()
			return nil
		},
	})
}

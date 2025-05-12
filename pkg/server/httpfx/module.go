package httpfx

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/froppa/go-svc-template/internal/service"
	"github.com/froppa/go-svc-template/pkg/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewMux(svc service.Service) *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}

func HTTPPort(cfg *config.Config) int { return cfg.Server.HTTPPort }

var Module = fx.Options(
	fx.Provide(HTTPPort),
	fx.Provide(NewMux),
	fx.Invoke(NewHTTPServer), // binds mux + port + logger → StartHTTPServer
)

func NewHTTPServer(lc fx.Lifecycle, port int, mux *http.ServeMux, log *zap.Logger) {
	addr := fmt.Sprintf(":%d", port)
	srv := &http.Server{Addr: addr, Handler: mux}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go srv.ListenAndServe()
			log.Info("HTTP listening", zap.Int("port", port))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("HTTP shutdown", zap.Int("port", port))
			c, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			return srv.Shutdown(c)
		},
	})
}

package http

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"time"

	"github.com/oluwatunmise/janus-proxy/services/dataplane/internal/config"
)

const (
	shutdownTimeout = 10 * time.Second
)

type Server struct {
	logger     *slog.Logger
	httpServer *http.Server
}

func New(cfg config.Config, logger *slog.Logger, handler http.Handler) *Server {
	return &Server{
		logger: logger,
		httpServer: &http.Server{
			Addr:    cfg.HTTPAddr,
			Handler: handler,
		},
	}
}

func (s *Server) Run(ctx context.Context) error {
	go func() {
		<-ctx.Done()

		s.logger.Info("Shutdown signal received")

		shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)

		defer cancel()

		_ = s.httpServer.Shutdown(shutdownCtx)
	}()

	err := s.httpServer.ListenAndServe()

	if errors.Is(err, http.ErrServerClosed) {
		return nil
	}

	return err
}

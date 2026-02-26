package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/oluwatunmise/janus-proxy/services/ops-console/internal/config"
	"github.com/oluwatunmise/janus-proxy/services/ops-console/internal/log"
	httpserver "github.com/oluwatunmise/janus-proxy/services/ops-console/internal/server/http"
)

type App struct {
	cfg    config.Config
	logger *slog.Logger
	server *httpserver.Server
}

func New() (*App, error) {
	cfg, err := config.Load()

	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	logger := log.New()

	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	mux.HandleFunc("/readyz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ready"))
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("janus ops console scaffold"))
	})

	server := httpserver.New(cfg, logger, mux)

	return &App{cfg: cfg, logger: logger, server: server}, nil
}

func (a *App) Run(ctx context.Context) error {
	a.logger.Info("Ops Console Starting", "addr", a.cfg.HTTPAddr)

	return a.server.Run(ctx)
}

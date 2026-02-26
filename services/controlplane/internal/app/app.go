package app

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/oluwatunmise/janus-proxy/services/controlplane/internal/config"
	"github.com/oluwatunmise/janus-proxy/services/controlplane/internal/log"
	httpserver "github.com/oluwatunmise/janus-proxy/services/controlplane/internal/server/http"
	"github.com/oluwatunmise/janus-proxy/services/controlplane/internal/store"
)

type App struct {
	cfg    config.Config
	logger *slog.Logger
	db     *sql.DB
	server *httpserver.Server
}

func New() (*App, error) {
	cfg, err := config.Load()

	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	logger := log.New()
	db, err := store.NewDB(cfg)

	if err != nil {
		return nil, fmt.Errorf("connect db: %w", err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	mux.HandleFunc("/readyz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ready"))
	})

	server := httpserver.New(cfg, logger, mux)

	return &App{cfg: cfg, logger: logger, db: db, server: server}, nil
}

func (a *App) Run(ctx context.Context) error {
	a.logger.Info("ControlPlane Starting", "addr", a.cfg.HTTPAddr)

	defer func() { _ = a.db.Close() }()

	return a.server.Run(ctx)
}

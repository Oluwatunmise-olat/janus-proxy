package app

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/oluwatunmise/janus-proxy/services/diff-worker/internal/config"
	"github.com/oluwatunmise/janus-proxy/services/diff-worker/internal/log"
)

type App struct {
	cfg    config.Config
	logger *slog.Logger
}

func New() (*App, error) {
	cfg, err := config.Load()

	if err != nil {
		return nil, fmt.Errorf("load config: %w", err)
	}

	logger := log.New()

	return &App{cfg: cfg, logger: logger}, nil
}

func (a *App) Run(ctx context.Context) error {
	a.logger.Info("Diff Worker Starting")

	<-ctx.Done()

	a.logger.Info("Shutdown signal received")

	return nil
}

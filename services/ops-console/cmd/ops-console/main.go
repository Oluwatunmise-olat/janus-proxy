package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/oluwatunmise/janus-proxy/services/ops-console/internal/app"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	a, err := app.New()

	if err != nil {
		log.Fatalf("Ops Console Bootstrap failed: %v", err)
	}

	if err := a.Run(ctx); err != nil {
		log.Fatalf("Ops Console stopped with error: %v", err)
	}

	os.Exit(0)
}

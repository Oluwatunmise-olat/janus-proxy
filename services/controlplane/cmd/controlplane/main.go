package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/oluwatunmise/janus-proxy/services/controlplane/internal/app"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	a, err := app.New()

	if err != nil {
		log.Fatalf("ControlPlane Bootstrap Failed: %v", err)
	}

	if err := a.Run(ctx); err != nil {
		log.Fatalf("ControlPlane Stopped with error: %v", err)
	}

	os.Exit(0)
}

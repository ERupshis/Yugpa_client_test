package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/erupshis/yugpa_test/internal/agent"
	dialer2 "github.com/erupshis/yugpa_test/internal/agent/dialer"
	"github.com/erupshis/yugpa_test/internal/config"
	"github.com/erupshis/yugpa_test/internal/logger"
)

func main() {
	//log.
	log, err := logger.CreateZapLogger("info")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to create logger: %v", err)
	}

	//config.
	cfg := config.Parse()

	//dialer.
	dialer := dialer2.CreateDefaultTCP(cfg.ServerAddr, log)

	//agent
	client := agent.Create(dialer, log)

	ctxWithCancel, cancel := context.WithCancel(context.Background())
	defer cancel()

	//client's goroutines init.
	client.Serve(ctxWithCancel, cfg.ConnectionsCount)

	// Create a channel to wait for signals (e.g., Ctrl+C) to gracefully exit.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh
}

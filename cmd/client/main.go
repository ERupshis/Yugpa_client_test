package main

import (
	"context"
	"fmt"
	"os"
	"yugpa_test/internal/agent"
	"yugpa_test/internal/config"
	"yugpa_test/internal/logger"
)

func main() {
	//log.
	log, err := logger.CreateZapLogger("info")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to create logger: %v", err)
	}

	cfg := config.Parse()

	client := agent.Create(cfg.ServerAddr)

	ctxWithCancel, cancel := context.WithCancel(context.Background())
	defer cancel()

	client.Serve(ctxWithCancel, 4)

}

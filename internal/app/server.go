package app

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
)

func StartServer(ctx context.Context, port int) error {
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}

	go func(ctx context.Context) {
		<-ctx.Done()
		if err := srv.Shutdown(context.Background()); err != nil {
			fmt.Printf("Server shutdown failed: %v\n", err)
		}
	}(ctx)

	slog.Info("Starting server", "port", port)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("Server failed to start", "error", err)
		return fmt.Errorf("listen: %w", err)
	}

	return nil
}

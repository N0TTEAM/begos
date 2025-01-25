package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/N0TTEAM/begos/internal/config"
	"github.com/N0TTEAM/begos/internal/db"
)

func main() {
	cfg := config.LoadConf()

	router := http.NewServeMux()
	database := db.NewConnection(&cfg.Postgres)

	sqlDB, err := database.DB()

	if err != nil {
		slog.Error("Failed to get database instance", slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer func() {
		if err := sqlDB.Close(); err != nil {
			slog.Error("Failed to close database connection", slog.String("error", err.Error()))
		}
	}()

	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcomee"))
	})

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("server started", slog.String("address", cfg.Addr))

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start")
		}
	}()

	<-done

	slog.Info("Shutting down")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("failed to shutdown server", slog.String("error", err.Error()))
	}

	slog.Info("server shutdown success")
}

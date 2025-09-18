package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/singhJasvinder101/basic-api/internal/config"
	"github.com/singhJasvinder101/basic-api/internal/http/handlers/student"
	"github.com/singhJasvinder101/basic-api/storage/sqlite"
)

func main() {
	// load config
	cfg := config.MustLoad()

	// database setup
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	slog.Info("Storage initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	// setup router
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetStudentById(storage))

	// setups server

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.HttpServer.Host, cfg.HttpServer.Port),
		Handler: router,
	}

	println("server running....")

	err = server.ListenAndServe()

	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}

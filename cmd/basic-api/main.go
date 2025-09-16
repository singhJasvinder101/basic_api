package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/singhJasvinder101/basic-api/internal/config"
)

func main() {
	// load config
	// cfg := config.MustLoad()
	cfgServer := config.MustLoad().HttpServer

	// database setup
	// setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("running...."))
	})

	// setups server

	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfgServer.Host, cfgServer.Port),
		Handler: router,
	}

	println("server running....")
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}

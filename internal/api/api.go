package api

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/farbautie/notex/config"
	"github.com/farbautie/notex/pkg/server"
)

func Run(config *config.Config) {
	r := NewRouter()
	log.Printf("Starting server on port %s", config.Http.Port)
	srv := server.New(r, server.Port(config.Http.Port))
	srv.Start()
	
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		log.Printf("Received signal %s, shutting down...", s)
	case err := <-srv.Notify():
		log.Printf("Server stopped with error: %s", err)
	}

	if err := srv.Shutdown(); err != nil {
		log.Printf("Error shutting down server: %s", err)
	}
}

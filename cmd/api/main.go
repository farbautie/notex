package main

import (
	"log"

	"github.com/farbautie/notex/config"
	"github.com/farbautie/notex/internal/api"
)

func main() {
	cfg, err := config.DefineConfig()
	if err != nil {
		log.Fatalf("Error loading config: %s", err)
	}
	api.Run(cfg)
}

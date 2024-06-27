package main

import (
	"github/cesardev31/proxy-go/config"
	"github/cesardev31/proxy-go/internal/server"
	"log"
)

func main() {
	// Carga la configuración desde el archivo config.toml
	cfg, err := config.LoadConfig("config.toml")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	// Crea una nueva instancia del servidor con la configuración cargada
	srv := server.NewServer(cfg)
	log.Println("Starting proxy server...")
	// Inicia el servidor y maneja cualquier error que ocurra durante la ejecución
	if err := srv.Run(); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}

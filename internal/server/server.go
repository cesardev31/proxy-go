package server

import (
	"github/cesardev31/proxy-go/config"
	"github/cesardev31/proxy-go/pkg/proxy"
	"log"
	"net/http"
)

// Server representa el servidor proxy
type Server struct {
	config *config.Config
}

// NewServer crea una nueva instancia del servidor con la configuración proporcionada
func NewServer(cfg *config.Config) *Server {
	return &Server{config: cfg}
}

// Run inicia el servidor proxy
func (s *Server) Run() error {
	var backendURLs []string
	for _, backend := range s.config.Backends {
		backendURLs = append(backendURLs, backend.URL)
	}
	http.HandleFunc("/", proxy.ProxyHandler(backendURLs))
	address := s.config.EntryPoints.Web.Address
	log.Printf("Listening on port %s...", address)
	return http.ListenAndServe(address, nil)
}

package server

import (
	"github/cesardev31/proxy-go/config"
	"github/cesardev31/proxy-go/pkg/proxy"
	"log"
	"net/http"
)

type Server struct {
	config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{config: cfg}
}

func (s *Server) Run() error {
	var backendURLs []string
	for _, backend := range s.config.Backends {
		backendURLs = append(backendURLs, backend.URL)
	}

	proxyHandler := proxy.ProxyHandler(backendURLs)
	loggerMiddleware, err := proxy.NewLoggerMiddleware(s.config.AccessLog.FilePath)
	if err != nil {
		return err
	}

	http.Handle("/", loggerMiddleware(http.HandlerFunc(proxyHandler)))
	address := s.config.EntryPoints.Web.Address
	log.Printf("Listening on port %s...", address)
	return http.ListenAndServe(address, nil)
}

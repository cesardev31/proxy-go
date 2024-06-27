package server

import (
	"github/cesardev31/proxy-go/config"
	"github/cesardev31/proxy-go/pkg/proxy"
	"log"
	"net"
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

	go func() {
		if err := http.ListenAndServe(address, nil); err != nil {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}()

	// Configuración para manejar conexiones TCP
	for _, tcpBackend := range s.config.TCPBackends {
		go s.handleTCP(tcpBackend)
	}

	select {}
}

func (s *Server) handleTCP(backend config.TCPBackendConfig) {
	listener, err := net.Listen("tcp", ":3333") // Escuchar en el puerto configurado
	if err != nil {
		log.Fatalf("Error starting TCP listener: %v", err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting TCP connection: %v", err)
			continue
		}

		go proxy.HandleTCPTunnel(conn, backend.Host, backend.Port)
	}
}

package proxy

import (
	"io"
	"log"
	"net"
)

func HandleTCPTunnel(clientConn net.Conn, backendHost, backendPort string) {
	// Conectar al backend
	backendConn, err := net.Dial("tcp", net.JoinHostPort(backendHost, backendPort))
	if err != nil {
		log.Printf("Error connecting to backend: %v", err)
		clientConn.Close()
		return
	}
	defer backendConn.Close()

	// Canal para terminar la transferencia de datos
	done := make(chan struct{})

	// Transferir datos del cliente al backend
	go func() {
		io.Copy(backendConn, clientConn)
		done <- struct{}{}
	}()

	// Transferir datos del backend al cliente
	go func() {
		io.Copy(clientConn, backendConn)
		done <- struct{}{}
	}()

	<-done
	clientConn.Close()
}

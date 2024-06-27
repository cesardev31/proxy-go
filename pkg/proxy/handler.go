package proxy

import (
	"io"
	"log"
	"net/http"
)

// ProxyHandler crea un manejador HTTP que reenvía las solicitudes a uno de los servidores backend especificados
func ProxyHandler(backendURLs []string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Utiliza la primera URL del backend para este ejemplo simple
		backendURL := backendURLs[0]

		// Crea una nueva solicitud HTTP para el backend
		req, err := http.NewRequest(r.Method, backendURL+r.RequestURI, r.Body)
		if err != nil {
			// Si hay un error al crear la solicitud, responde con un error HTTP 500 (Internal Server Error)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Copia las cabeceras de la solicitud original a la nueva solicitud
		for k, v := range r.Header {
			req.Header[k] = v
		}

		// Envía la solicitud al backend
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			// Si hay un error al realizar la solicitud, responde con un error HTTP 502 (Bad Gateway)
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
		defer resp.Body.Close()

		// Copia las cabeceras de la respuesta del backend a la respuesta original
		for k, v := range resp.Header {
			w.Header()[k] = v
		}

		// Establece el código de estado de la respuesta original al código de estado de la respuesta del backend
		w.WriteHeader(resp.StatusCode)

		// Copia el cuerpo de la respuesta del backend al cuerpo de la respuesta original
		_, err = io.Copy(w, resp.Body)
		if err != nil {
			// Si hay un error al copiar el cuerpo de la respuesta, lo registra en los logs
			log.Printf("Error copying response body: %v", err)
		}
	}
}

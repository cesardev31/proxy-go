package proxy

import (
	"log"
	"net/http"
	"os"
	"time"
)

func NewLoggerMiddleware(filePath string) (func(http.Handler) http.Handler, error) {
	var logger *log.Logger

	if filePath != "" {
		file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			return nil, err
		}
		logger = log.New(file, "", log.LstdFlags)
	} else {
		logger = log.New(os.Stdout, "", log.LstdFlags)
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			rw := &responseWriter{w, http.StatusOK}
			next.ServeHTTP(rw, r)
			duration := time.Since(start)

			logger.Printf("%s %s %s %d %s", r.RemoteAddr, r.Method, r.URL, rw.statusCode, duration)
		})
	}, nil
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/handlers"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {

	s.httpServer = &http.Server{
		Addr:           port,
		Handler:        s.handler(handler),
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func (s *Server) handler(handler http.Handler) http.Handler {
	handler = s.cors(handler)
	handler = s.gzip(handler)
	return handler
}

func (s *Server) cors(handler http.Handler) http.Handler {
	options := []handlers.CORSOption{
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowCredentials(),
		handlers.AllowedHeaders([]string{"Content-Type"}),
		handlers.AllowedMethods([]string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPost,
			http.MethodPut,
			http.MethodOptions,
			http.MethodPatch,
		}),
	}
	return handlers.CORS(options...)(handler)
}

func (s *Server) gzip(handler http.Handler) http.Handler {
	return handlers.CompressHandler(handler)
}

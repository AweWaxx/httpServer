package server

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	logger     *log.Logger
	httpServer *http.Server
}

func New(logger *log.Logger) *Server {
	router := createRouter(logger)

	httpServer := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		logger:     logger,
		httpServer: httpServer,
	}
}

func createRouter(logger *log.Logger) *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", handlers.RootHandler)
	router.HandleFunc("/upload", handlers.UploadHandler)
	logger.Println("Роутер создан, хендлеры зарегистрированы")
	return router
}

func (s *Server) Start() error {
	s.logger.Printf("Сервер запущен на %s\n", s.httpServer.Addr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

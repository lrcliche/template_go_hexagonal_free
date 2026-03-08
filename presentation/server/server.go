package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"template-go-hexagonal/application/config"
	"template-go-hexagonal/presentation/container"
	"template-go-hexagonal/presentation/routes"
)

const shutdownTimeout = 10 * time.Second

type Server struct {
	httpServer *http.Server
}

func NewServer(_ context.Context, cfg config.Config) (*Server, error) {
	appContainer := container.New()
	engine := routes.New(appContainer)

	httpServer := &http.Server{
		Addr:              ":" + cfg.HTTPPort,
		Handler:           engine,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       time.Duration(cfg.ReadTimeoutSeconds) * time.Second,
		WriteTimeout:      time.Duration(cfg.WriteTimeoutSeconds) * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	return &Server{httpServer: httpServer}, nil
}

func (s *Server) Run() error {
	log.Printf("[BOOT] operation=start-server message=server running on http://localhost%s", s.httpServer.Addr)

	listenErr := make(chan error, 1)
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			listenErr <- err
			return
		}
		listenErr <- nil
	}()

	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(stopSignal)

	select {
	case err := <-listenErr:
		return err
	case <-stopSignal:
	}

	shutdownContext, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := s.httpServer.Shutdown(shutdownContext); err != nil {
		return err
	}

	if err := <-listenErr; err != nil {
		return err
	}

	log.Printf("[BOOT] operation=shutdown message=server stopped")
	return nil
}

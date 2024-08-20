package protocol

import (
	"backend/pkg/config"
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"

	"go.uber.org/zap"
)

type Server struct {
	Dependencies    *Dependencies
	Router          *fiber.App
	log             *zap.Logger
	ShutdownTimeout time.Duration
}

type Dependencies struct {
	Cfg *config.Config
}

func NewServer(deps *Dependencies, cfg *config.Config) *Server {
	r := fiber.New()
	return &Server{
		Dependencies:    deps,
		Router:          r,
		log:             zap.L().Named("server"),
		ShutdownTimeout: time.Duration(30) * time.Second,
	}
}

func (s *Server) registerRoutes() {
	// r := s.Router.Group("/api")

}

func (s *Server) Run(ctx context.Context) error {
	stopCh := ctx.Done()

	s.registerRoutes()

	addr := fmt.Sprintf(":%s", s.Dependencies.Cfg.Server.HTTPPort)

	s.log.Info("Starting HTTP server", zap.String("port", s.Dependencies.Cfg.Server.HTTPPort))
	err := s.Router.Listen(addr)
	if err != nil {
		return err
	}

	go func() {
		<-stopCh
		fmt.Println("Shutting down server...")
		s.Router.ShutdownWithContext(ctx)
	}()

	return nil
}

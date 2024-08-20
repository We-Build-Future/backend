package app

import (
	"backend/pkg/config"
	"backend/pkg/infra/storage/db"
	"backend/pkg/protocol"
	"backend/registry"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sync"

	"go.uber.org/zap"
)

type Server struct {
	postgresDB *sql.DB
	services   []registry.RunFunc
	log        *zap.Logger
}

func NewServer(isStandaloneMode bool) (*Server, error) {
	cfg, err := config.FromEnv()
	if err != nil {
		return nil, err
	}

	postgresDB, err := db.New(cfg.Postgres.GetConnectionString())
	if err != nil {
		return nil, err
	}

	restServer := protocol.NewServer(&protocol.Dependencies{
		Cfg: cfg,
	}, cfg)

	services := registry.NewServiceRegistry(
		restServer.Run,
	)

	if isStandaloneMode {
		services = registry.NewServiceRegistry(
			restServer.Run,
		)
	}

	return &Server{
		postgresDB: postgresDB,
		services:   services.GetServices(),
		log:        zap.L().Named("apiserver"),
	}, nil
}

func (s *Server) Run(ctx context.Context) {
	defer func() {
		s.postgresDB.Close()
	}()

	var wg sync.WaitGroup
	wg.Add(len(s.services))

	for _, svc := range s.services {
		go func(svc registry.RunFunc) error {
			defer wg.Done()
			err := svc(ctx)
			if err != nil && !errors.Is(err, context.Canceled) {
				s.log.Error("stopped server", zap.String("service", serviceName), zap.Error(err))
				return fmt.Errorf("%s run error: %w", serviceName, err)
			}

			return nil
		}(svc)
	}

	wg.Wait()
}

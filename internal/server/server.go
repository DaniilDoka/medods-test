package server

import (
	"fmt"
	"medods-test/config"
	"medods-test/internal/auth"
	auth_repository "medods-test/internal/auth/repository"
	auth_usercase "medods-test/internal/auth/usecase"
	"medods-test/pkg/logger"
	"medods-test/pkg/pg"

	"github.com/gofiber/fiber/v3"
)

type server struct {
	app    *fiber.App
	logger *logger.Logger
	cfg    *config.Config
}

func NewServer(logger *logger.Logger, cfg *config.Config) *server {
	return &server{
		app:    fiber.New(),
		logger: logger,
		cfg:    cfg,
	}
}

func (s *server) MapRoutes(pgConn *pg.Pg) {
	authRepo := auth_repository.NewRepository(pgConn)
	authUsecase := auth_usercase.NewUsecase(authRepo, s.cfg, s.logger)
	authHandlers := auth.NewHandlers(authUsecase, s.logger)

	authHandlers.MapRoutes(s.app.Group("/user"))
}

func (s *server) Run() error {
	return s.app.Listen(fmt.Sprintf(":%d", s.cfg.Server.Port))
}

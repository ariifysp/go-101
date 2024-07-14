package server

import (
	"errors"
	"fmt"
	"github/ariifysp/go-101/config"
	"github/ariifysp/go-101/pkg/cache"
	"github/ariifysp/go-101/pkg/custom"
	"github/ariifysp/go-101/pkg/database"
	"sync"

	"github.com/gofiber/fiber/v3"
)

type Server struct {
	App    *fiber.App
	Config *config.Config
	DB     database.DatabaseInterface
	Cache  cache.RedisInterface
}

var (
	once   sync.Once
	server *Server
)

func NewServer(cfg *config.Config, db database.DatabaseInterface, cache cache.RedisInterface) *Server {
	once.Do(func() {
		server = &Server{
			App:    fiber.New(),
			Config: cfg,
			DB:     db,
			Cache:  cache,
		}
	})
	return server
}

func (s *Server) Start() {
	s.App.Get("/v1/health", s.healthCheck)
	s.ItemRouter()

	s.App.Use(s.handleEndpointNotFound)

	s.fiberListening()
}

func (s *Server) handleEndpointNotFound(ctx fiber.Ctx) error {
	return custom.ErrorResponse(ctx, fiber.ErrInternalServerError, errors.New("Endpoint not found"))
}

func (s *Server) healthCheck(ctx fiber.Ctx) error {
	return custom.SuccessResponse(ctx, fiber.StatusOK, "OK")
}

func (s *Server) fiberListening() {
	host := s.Config.App.Host
	port := s.Config.App.Port
	fiberConnectURL := fmt.Sprintf("%s:%s", host, port)

	if err := s.App.Listen(fiberConnectURL); err != nil {
		panic(err)
	}
}

package server

import (
	"fmt"

	"github.com/acm-uiuc/core/controller"
	"github.com/acm-uiuc/core/service"
)

type Server struct {
	cntlr *controller.Controller
	svc   *service.Service
}

func (server *Server) Start(port string) error {
	return server.cntlr.Start(port)
}

func New() (*Server, error) {
	svc, err := service.New()
	if err != nil {
		return nil, fmt.Errorf("failed to create services: %w", err)
	}

	cntlr, err := controller.New(svc)
	if err != nil {
		return nil, fmt.Errorf("failed to create controllers: %w", err)
	}

	return &Server{
		cntlr: cntlr,
		svc:   svc,
	}, nil
}

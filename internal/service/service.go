package service

import (
	"github.com/begenov/api-gateway/internal/repository"
	"github.com/begenov/api-gateway/pkg/logger"
)

type Service struct {
}

func NewService(repo *repository.Repository, logger *logger.Logger) *Service {
	return &Service{}
}

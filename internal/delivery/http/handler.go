package http

import (
	"github.com/begenov/api-gateway/internal/client"
	"github.com/begenov/api-gateway/internal/service"
	"github.com/begenov/api-gateway/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler(service *service.Service, logger *logger.Logger, registerClient *client.RegisterServiceClient) *Handler {
	return &Handler{}
}

func (h *Handler) Init() *gin.Engine {
	return nil
}

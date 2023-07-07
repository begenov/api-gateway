package v1

import (
	"github.com/begenov/api-gateway/internal/client"
	"github.com/begenov/api-gateway/internal/service"
	"github.com/begenov/api-gateway/pkg/logger"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service        *service.Service
	logger         *logger.Logger
	registerClient *client.RegisterServiceClient
}

func NewHandler(service *service.Service, logger *logger.Logger, registerClient *client.RegisterServiceClient) *Handler {
	return &Handler{
		service:        service,
		logger:         logger,
		registerClient: registerClient,
	}
}

func (h *Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.registerLoadRouter(v1)
	}
}

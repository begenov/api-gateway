package http

import (
	"github.com/begenov/api-gateway/internal/client"
	v1 "github.com/begenov/api-gateway/internal/delivery/http/v1"
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

func (h *Handler) Init() *gin.Engine {
	router := gin.Default()

	h.init(router)

	return router
}

func (h *Handler) init(router *gin.Engine) {
	getewayRestHandler := v1.NewHandler(h.service, h.logger, h.registerClient)
	api := router.Group("/api")

	getewayRestHandler.Init(api)
}

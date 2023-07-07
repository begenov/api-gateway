package v1

import (
	"github.com/gin-gonic/gin"
)

type ResponseMessage struct {
	Message any `json:"message"`
}

type ResponseError struct {
	ErrorMessage any `json:"error_message"`
}

func (h *Handler) newResponseError(c *gin.Context, statusCode int, message string) {
	h.logger.Sugar().Info(message)
	c.AbortWithStatusJSON(statusCode, ResponseError{ErrorMessage: message})
}

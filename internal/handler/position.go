package handler

import (
	"github.com/gin-gonic/gin"
	"core_service/internal/service"
)

type PositionHandler struct {
	*Handler
	positionService service.PositionService
}

func NewPositionHandler(
    handler *Handler,
    positionService service.PositionService,
) *PositionHandler {
	return &PositionHandler{
		Handler:      handler,
		positionService: positionService,
	}
}

func (h *PositionHandler) GetPosition(ctx *gin.Context) {

}

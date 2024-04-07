package handler

import (
	"github.com/gin-gonic/gin"
	"core_service/internal/service"
)

type ActionEventHandler struct {
	*Handler
	actionEventService service.ActionEventService
}

func NewActionEventHandler(
    handler *Handler,
    actionEventService service.ActionEventService,
) *ActionEventHandler {
	return &ActionEventHandler{
		Handler:      handler,
		actionEventService: actionEventService,
	}
}

func (h *ActionEventHandler) GetActionEvent(ctx *gin.Context) {

}

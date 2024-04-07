package handler

import (
	"github.com/gin-gonic/gin"
	"core_service/internal/service"
)

type WorkspaceHandler struct {
	*Handler
	workspaceService service.WorkspaceService
}

func NewWorkspaceHandler(
    handler *Handler,
    workspaceService service.WorkspaceService,
) *WorkspaceHandler {
	return &WorkspaceHandler{
		Handler:      handler,
		workspaceService: workspaceService,
	}
}

func (h *WorkspaceHandler) GetWorkspace(ctx *gin.Context) {

}

package handler

import (
	"github.com/gin-gonic/gin"
	"core_service/internal/service"
)

type DocumentHandler struct {
	*Handler
	documentService service.DocumentService
}

func NewDocumentHandler(
    handler *Handler,
    documentService service.DocumentService,
) *DocumentHandler {
	return &DocumentHandler{
		Handler:      handler,
		documentService: documentService,
	}
}

func (h *DocumentHandler) GetDocument(ctx *gin.Context) {

}

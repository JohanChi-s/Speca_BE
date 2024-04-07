package handler

import (
	"github.com/gin-gonic/gin"
	"core_service/internal/service"
)

type CollectionHandler struct {
	*Handler
	collectionService service.CollectionService
}

func NewCollectionHandler(
    handler *Handler,
    collectionService service.CollectionService,
) *CollectionHandler {
	return &CollectionHandler{
		Handler:      handler,
		collectionService: collectionService,
	}
}

func (h *CollectionHandler) GetCollection(ctx *gin.Context) {

}

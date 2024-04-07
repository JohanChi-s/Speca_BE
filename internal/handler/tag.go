package handler

import (
	"github.com/gin-gonic/gin"
	"core_service/internal/service"
)

type TagHandler struct {
	*Handler
	tagService service.TagService
}

func NewTagHandler(
    handler *Handler,
    tagService service.TagService,
) *TagHandler {
	return &TagHandler{
		Handler:      handler,
		tagService: tagService,
	}
}

func (h *TagHandler) GetTag(ctx *gin.Context) {

}

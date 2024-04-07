package handler

import (
	"github.com/gin-gonic/gin"
	"core_service/internal/service"
)

type CommentHandler struct {
	*Handler
	commentService service.CommentService
}

func NewCommentHandler(
    handler *Handler,
    commentService service.CommentService,
) *CommentHandler {
	return &CommentHandler{
		Handler:      handler,
		commentService: commentService,
	}
}

func (h *CommentHandler) GetComment(ctx *gin.Context) {

}

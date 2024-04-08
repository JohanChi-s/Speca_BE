package handler

import (
	"github.com/gin-gonic/gin"
	"core_service/internal/service"
)

type MemberHandler struct {
	*Handler
	memberService service.MemberService
}

func NewMemberHandler(
    handler *Handler,
    memberService service.MemberService,
) *MemberHandler {
	return &MemberHandler{
		Handler:      handler,
		memberService: memberService,
	}
}

func (h *MemberHandler) GetMember(ctx *gin.Context) {

}

package handler

import (
	"core_service/internal/service"

	"github.com/gin-gonic/gin"
)

type ProfileHandler struct {
	*Handler
	profileService service.ProfileService
}

func NewProfileHandler(
	handler *Handler,
	profileService service.ProfileService,
) *ProfileHandler {
	return &ProfileHandler{
		Handler:        handler,
		profileService: profileService,
	}
}

func (h *ProfileHandler) GetProfile(ctx *gin.Context) {

}

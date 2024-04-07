package handler

import (
	"github.com/gin-gonic/gin"
	"core_service/internal/service"
)

type TeamHandler struct {
	*Handler
	teamService service.TeamService
}

func NewTeamHandler(
    handler *Handler,
    teamService service.TeamService,
) *TeamHandler {
	return &TeamHandler{
		Handler:      handler,
		teamService: teamService,
	}
}

func (h *TeamHandler) GetTeam(ctx *gin.Context) {

}

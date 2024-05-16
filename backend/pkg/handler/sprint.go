package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createSprints(c *gin.Context, eventId int) int {
	id, err := h.services.Sprint.Create(eventId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return 0
	}

	return id
}

func (h *Handler) getAllSprints(c *gin.Context) {

}

func (h *Handler) getSprintById(c *gin.Context) {

}

func (h *Handler) updateSprint(c *gin.Context) {}

func (h *Handler) deleteSprint(c *gin.Context) {

}

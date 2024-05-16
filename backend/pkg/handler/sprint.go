package handler

import (
	"github.com/gin-gonic/gin"
	augventure "github.com/klausfun/Augventure"
	"net/http"
)

func (h *Handler) createSprints(c *gin.Context) {
	var input augventure.Sprint
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Sprint.Create(-1, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) getAllSprints(c *gin.Context) {

}

func (h *Handler) getSprintById(c *gin.Context) {

}

func (h *Handler) updateSprint(c *gin.Context) {

}

func (h *Handler) deleteSprint(c *gin.Context) {

}

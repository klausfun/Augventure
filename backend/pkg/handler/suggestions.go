package handler

import (
	"github.com/gin-gonic/gin"
	augventure "github.com/klausfun/Augventure"
	"net/http"
)

func (h *Handler) createSuggestions(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input augventure.Suggestion
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	suggestionId, err := h.services.Suggestion.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"suggestionId": suggestionId,
	})
}

func (h *Handler) getSuggestionsBySprintId(c *gin.Context) {

}

func (h *Handler) voteSuggestions(c *gin.Context) {

}

func (h *Handler) deleteSuggestions(c *gin.Context) {

}

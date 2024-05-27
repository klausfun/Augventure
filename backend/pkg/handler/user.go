package handler

import (
	"github.com/gin-gonic/gin"
	augventure "github.com/klausfun/Augventure"
	"net/http"
)

type getAuthorResponse struct {
	User augventure.Author `json:"user"`
}

func (h *Handler) getUser(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	user, err := h.services.Profile.GetById(userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAuthorResponse{
		User: user,
	})

}

func (h *Handler) updatePFP(c *gin.Context) {

}

func (h *Handler) updatePassword(c *gin.Context) {

}

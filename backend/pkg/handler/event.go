package handler

import (
	"github.com/gin-gonic/gin"
	augventure "github.com/klausfun/Augventure"
	"net/http"
	"strconv"
)

const (
	// For sprints
	votingStateSprint       = "voting" // default
	implementingStateSprint = "implementing"
	endedStateSprint        = "ended"

	// For events
	scheduledStateEvent  = "scheduled"
	inProgressStateEvent = "in_progress"
	endedStateEvent      = "ended"
	empty                = ""
)

func (h *Handler) createEvents(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input augventure.Event
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	eventId, err := h.services.Event.Create(userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	sprintId, err := h.services.Sprint.Create(eventId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"eventId":  eventId,
		"sprintId": sprintId,
	})
}

type getAllEventsResponse struct {
	Data []augventure.Event `json:"data"`
}

func (h *Handler) getAllEvents(c *gin.Context) {
	events, err := h.services.Event.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllEventsResponse{
		Data: events,
	})
}

func (h *Handler) getEventById(c *gin.Context) {
	eventId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid is param")
		return
	}

	event, err := h.services.Event.GetById(eventId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, event)
}

func (h *Handler) updateEvent(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	eventId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid is param")
		return
	}

	resp, err := h.services.Event.CheckingTheStatus(eventId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	} else if !resp {
		newErrorResponse(c, http.StatusBadRequest, "The event status is not 'in_progress'")
		return
	}

	var input augventure.UpdateEventInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	input.Status = empty

	if err := h.services.Event.Update(userId, eventId, input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{"ok"})
}

func (h *Handler) deleteEvent(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	eventId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	err = h.services.Event.Delete(userId, eventId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

func (h *Handler) finishVoting(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	eventId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid is param")
		return
	}

	resp, err := h.services.Event.CheckingTheStatus(eventId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	} else if !resp {
		newErrorResponse(c, http.StatusBadRequest, "The event status is not 'in_progress'")
		return
	}

	var input augventure.FinishVoting
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	sprintId, err := h.services.Event.FinishVoting(userId, eventId, *input.SuggestionWinnerId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	finishVotingInput := augventure.UpdateSprintInput{
		SprintId:           sprintId,
		SuggestionWinnerId: input.SuggestionWinnerId,
		Status:             implementingStateSprint,
	}
	if err := h.services.Sprint.Update(finishVotingInput); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"status": "ok",
	})
}

func (h *Handler) finishImplementing(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	eventId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid is param")
		return
	}

	resp, err := h.services.Event.CheckingTheStatus(eventId)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	} else if !resp {
		newErrorResponse(c, http.StatusBadRequest, "The event status is not 'in_progress'")
		return
	}

	var input augventure.FinishImplementing
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	sprintId, err := h.services.Event.FinishImplementing(userId, eventId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	finishImplementingInput := augventure.UpdateSprintInput{
		SprintId:           sprintId,
		SuggestionWinnerId: nil,
		Status:             endedStateSprint,
	}
	if err := h.services.Sprint.Update(finishImplementingInput); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if input.IsLastSprint {
		input := augventure.UpdateEventInput{
			Title:       nil,
			Description: nil,
			Status:      endedStateEvent,
		}

		if err := h.services.Event.Update(userId, eventId, input); err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"status": "ok",
		})
	} else {
		newSprintId, err := h.services.Sprint.Create(eventId)
		if err != nil {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		c.JSON(http.StatusOK, map[string]interface{}{
			"status":        "ok",
			"new_sprint_id": newSprintId,
		})
	}
}

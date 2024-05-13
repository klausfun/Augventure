package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/klausfun/Augventure/pkg/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		users := api.Group("/users/me")
		{
			users.GET("/", h.getUser)
			users.PUT("/upload_pfp", h.updatePFP)
			users.PUT("/password_reset", h.updatePassword)
		}

		events := api.Group("/events")
		{
			events.POST("/", h.createEvents)
			events.DELETE("/:id", h.deleteEvent)
			events.GET("/", h.getAllEvents)
			events.GET("/:id", h.getEventById)
			events.PUT("/:id", h.updateEvent)
			events.PATCH("/:id/finish_voting", h.finishVoting)
			events.PATCH("/:id/finish_implementing", h.finishImplementing)
		}

		sprints := api.Group("/sprints")
		{
			sprints.POST("/", h.createSprints)
			sprints.DELETE("/:id", h.deleteSprint)
			sprints.GET("/", h.getAllSprints)
			sprints.GET("/:id", h.getSprintById)
			sprints.PUT("/:id", h.updateSprint)
		}

		suggestions := api.Group("/suggestions")
		{
			suggestions.POST("/", h.createSuggestions)
			suggestions.DELETE("/:id", h.deleteSuggestions)
			suggestions.GET("/", h.getSuggestionsBySprintId)
			suggestions.PUT("/:id/vote", h.voteSuggestions)
			//suggestions.PUT("/:id/add_media", h.addMedia)
		}
	}

	return router
}

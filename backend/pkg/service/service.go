package service

import (
	augventure "github.com/klausfun/Augventure"
	"github.com/klausfun/Augventure/pkg/infrastructure"
	"github.com/klausfun/Augventure/pkg/repository"
)

type Authorization interface {
	CreateUser(user augventure.User) (int, error)
	GenerateToken(password, email string) (augventure.Author, string, error)
	ParseToken(token string) (int, error)
}

type Event interface {
	Create(userId int, event augventure.Event) (int, error)
	GetAll() ([]augventure.EventAndSprints, error)
	GetById(eventId int) (augventure.EventAndSprints, error)
	Delete(userId, eventId int) error
	Update(userId, eventId int, input augventure.UpdateEventInput) error
	FinishVoting(userId, eventId, suggestionWinnerId int) (int, error)
	FinishImplementing(userId, eventId int) (int, error)
	CheckingTheStatus(eventId int) (bool, error)
	FilterEvents(authorId int) ([]augventure.Event, error)
}

type Sprint interface {
	Create(eventId int) (int, error)
	Update(input augventure.UpdateSprintInput) error
}

type Profile interface {
	GetById(userId int) (augventure.Author, error)
	UpdatePassword(userId int, input augventure.UpdatePasswordInput) error
}

type Suggestion interface {
	Create(userId int, suggestion augventure.Suggestion) (int, error)
	GetBySprintId(sprintId int) ([]augventure.FilterSuggestions, error)
	Vote(voteType bool, suggestionId, userId int) (int, error)
}

type Service struct {
	Authorization
	Event
	Sprint
	Profile
	Suggestion
}

func NewService(repos *repository.Repository, storage *infrastructure.S3Storage) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Event:         NewEventService(repos.Event),
		Sprint:        NewSprintService(repos.Sprint),
		Suggestion:    NewSuggestionService(repos.Suggestion, storage),
		Profile:       NewProfileService(repos.Profile),
	}
}

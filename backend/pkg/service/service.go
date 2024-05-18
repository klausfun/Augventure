package service

import (
	augventure "github.com/klausfun/Augventure"
	"github.com/klausfun/Augventure/pkg/infrastructure"
	"github.com/klausfun/Augventure/pkg/repository"
)

type Authorization interface {
	CreateUser(user augventure.User) (int, error)
	GenerateToken(password, email string) (string, error)
	ParseToken(token string) (int, error)
}

type Event interface {
	Create(userId int, event augventure.Event) (int, error)
	GetAll() ([]augventure.Event, error)
	GetById(eventId int) (augventure.Event, error)
	Delete(userId, eventId int) error
	Update(userId, eventId int, input augventure.UpdateEventInput) error
	FinishVoting(userId, eventId int) (int, error)
	FinishImplementing(userId, eventId int) (int, error)
}

type Sprint interface {
	Create(eventId int) (int, error)
	Update(input augventure.UpdateSprintInput) error
}

type Profile interface{}

type Suggestion interface {
	Create(userId int, suggestion augventure.Suggestion) (int, error)
	GetBySprintId(sprintId int) ([]augventure.FilterSuggestions, error)
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
	}
}

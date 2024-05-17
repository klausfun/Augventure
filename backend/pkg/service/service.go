package service

import (
	augventure "github.com/klausfun/Augventure"
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
}

type Sprint interface {
	Create(eventId int) (int, error)
	Update(input augventure.UpdateSprintInput) error
}

type Profile interface {
}

type Service struct {
	Authorization
	Event
	Sprint
	Profile
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Event:         NewEventService(repos.Event),
		Sprint:        NewSprintService(repos.Sprint),
	}
}

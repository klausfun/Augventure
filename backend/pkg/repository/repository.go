package repository

import (
	"github.com/jmoiron/sqlx"
	augventure "github.com/klausfun/Augventure"
)

type Authorization interface {
	CreateUser(user augventure.User) (int, error)
	GetUser(password, email string) (augventure.User, error)
}

type Event interface {
	Create(userId int, event augventure.Event) (int, error)
	GetAll() ([]augventure.Event, error)
	GetById(eventId int) (augventure.Event, error)
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
}

type Suggestion interface {
	Create(userId int, suggestion augventure.Suggestion) (int, error)
	GetBySprintId(sprintId int) ([]augventure.FilterSuggestions, error)
	Vote(voteType bool, suggestionId, userId int) (int, error)
}

type Repository struct {
	Authorization
	Event
	Sprint
	Profile
	Suggestion
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Event:         NewEventPostgres(db),
		Sprint:        NewSprintPostgres(db),
		Suggestion:    NewSuggestionPostgres(db),
		Profile:       NewUserPostgres(db),
	}
}

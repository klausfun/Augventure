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
}

type Sprint interface {
}

type Profile interface {
}

type Repository struct {
	Authorization
	Event
	Sprint
	Profile
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Event:         NewEventPostgres(db),
	}
}

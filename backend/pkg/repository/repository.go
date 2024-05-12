package repository

import (
	"github.com/jmoiron/sqlx"
	augventure "github.com/klausfun/Augventure"
)

type Authorization interface {
	CreateUser(user augventure.User) (int, error)
}

type Event interface {
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
	}
}

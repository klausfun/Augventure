package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
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
	return &Repository{}
}

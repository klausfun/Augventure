package repository

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

func NewRepository() *Repository {
	return &Repository{}
}

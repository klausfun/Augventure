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
}

type Sprint interface {
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
	}
}

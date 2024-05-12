package service

import "github.com/klausfun/Augventure/pkg/repository"

type Authorization interface {
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
	return &Service{}
}

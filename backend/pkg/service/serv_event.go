package service

import (
	augventure "github.com/klausfun/Augventure"
	"github.com/klausfun/Augventure/pkg/repository"
)

type EventService struct {
	repo repository.Event
}

func NewEventService(repo repository.Event) *EventService {
	return &EventService{repo: repo}
}

func (s *EventService) Create(userId int, event augventure.Event) (int, error) {
	return s.repo.Create(userId, event)
}

func (s *EventService) GetAll() ([]augventure.Event, error) {
	return s.repo.GetAll()
}

func (s *EventService) GetById(eventId int) (augventure.Event, error) {
	return s.repo.GetById(eventId)
}

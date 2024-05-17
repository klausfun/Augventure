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

func (s *EventService) Delete(userId, eventId int) error {
	return s.repo.Delete(userId, eventId)
}

func (s *EventService) Update(userId, eventId int, input augventure.UpdateEventInput) error {
	return s.repo.Update(userId, eventId, input)
}

func (s *EventService) FinishVoting(userId, eventId int) (int, error) {
	return s.repo.FinishVoting(userId, eventId)
}

package service

import (
	augventure "github.com/klausfun/Augventure"
	"github.com/klausfun/Augventure/pkg/repository"
)

type SprintService struct {
	repo repository.Sprint
}

func NewSprintService(repo repository.Sprint) *SprintService {
	return &SprintService{repo: repo}
}

func (s *SprintService) Create(eventId int, sprint augventure.Sprint) (int, error) {
	return s.repo.Create(eventId, sprint)
}

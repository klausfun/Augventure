package service

import (
	augventure "github.com/klausfun/Augventure"
	"github.com/klausfun/Augventure/pkg/repository"
)

type SuggestionService struct {
	repo repository.Suggestion
}

func NewSuggestionService(repo repository.Suggestion) *SuggestionService {
	return &SuggestionService{repo: repo}
}

func (s *SuggestionService) Create(userId int, suggestion augventure.Suggestion) (int, error) {
	return s.repo.Create(userId, suggestion)
}

package service

import (
	"errors"
	"fmt"
	augventure "github.com/klausfun/Augventure"
	"github.com/klausfun/Augventure/pkg/infrastructure"
	"github.com/klausfun/Augventure/pkg/repository"
)

type SuggestionService struct {
	repo    repository.Suggestion
	storage *infrastructure.S3Storage
}

func NewSuggestionService(repo repository.Suggestion, storage *infrastructure.S3Storage) *SuggestionService {
	return &SuggestionService{repo: repo, storage: storage}
}

func (s *SuggestionService) Create(userId int, suggestion augventure.Suggestion) (int, error) {
	// todo нужно генерировать уникальные ссылки!!!
	url, err := s.storage.SaveFile([]byte(suggestion.TextContent),
		fmt.Sprintf("sprintId=%d, userId=%d", suggestion.SprintId, userId))
	if err != nil {
		return 0, errors.New("error writing to the cloud: " + err.Error())
	}
	suggestion.TextContent = url

	return s.repo.Create(userId, suggestion)
}

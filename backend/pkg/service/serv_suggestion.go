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
	url, err := s.storage.SaveFile([]byte(suggestion.TextContent),
		fmt.Sprintf("Id=%d", repository.LastId))
	if err != nil {
		return 0, errors.New("error writing to the cloud: " + err.Error())
	}
	suggestion.TextContent = url

	return s.repo.Create(userId, suggestion)
}

func (s *SuggestionService) GetBySprintId(sprintId int) ([]augventure.FilterSuggestions, error) {
	filterSuggestions, err := s.repo.GetBySprintId(sprintId)
	if err != nil {
		return nil, err
	}

	for i, curSuggestion := range filterSuggestions {
		content, err := s.storage.GetFile(fmt.Sprintf("%s", curSuggestion.Content[43:]))
		if err != nil {
			return nil, errors.New("error receiving a file from the cloud: " + err.Error())
		}

		filterSuggestions[i].Content = string(content)
	}

	return filterSuggestions, nil
}

func (s *SuggestionService) Vote(voteType bool, suggestionId, userId int) error {
	return s.repo.Vote(voteType, suggestionId, userId)
}

package service

import (
	augventure "github.com/klausfun/Augventure"
	"github.com/klausfun/Augventure/pkg/repository"
)

type ProfileService struct {
	repo repository.Profile
}

func NewProfileService(repo repository.Profile) *ProfileService {
	return &ProfileService{repo: repo}
}

func (s *ProfileService) GetById(userId int) (augventure.Author, error) {
	return s.repo.GetById(userId)
}

func (s *ProfileService) UpdatePassword(userId int, input augventure.UpdatePasswordInput) error {
	input.NewPassword = generatePasswordHash(input.NewPassword)
	input.OldPassword = generatePasswordHash(input.OldPassword)

	return s.repo.UpdatePassword(userId, input)
}

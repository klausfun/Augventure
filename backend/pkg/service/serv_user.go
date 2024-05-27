package service

import (
	augventure "github.com/klausfun/Augventure"
	"github.com/klausfun/Augventure/pkg/repository"
)

type UserService struct {
	repo repository.Profile
}

func NewUserService(repo repository.Profile) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetById(userId int) (augventure.Author, error) {
	return s.repo.GetById(userId)
}

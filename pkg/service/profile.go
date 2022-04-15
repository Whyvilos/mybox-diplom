package service

import (
	"github.com/whyvilos/mybox"
	"github.com/whyvilos/mybox/pkg/repository"
)

type UserProfileService struct {
	repo repository.UserProfile
}

func NewUserProfileService(repo repository.UserProfile) *UserProfileService {
	return &UserProfileService{repo: repo}
}

func (s *UserProfileService) GetById(you_id, id_user int) (mybox.User, error) {
	return s.repo.GetById(you_id, id_user)
}

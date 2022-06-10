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

func (s *UserProfileService) Follow(you_id, id_user int) error {
	return s.repo.Follow(you_id, id_user)
}

func (s *UserProfileService) UnFollow(you_id, id_user int) error {
	return s.repo.UnFollow(you_id, id_user)
}

func (s *UserProfileService) LoadLine(you_id int) ([]mybox.Post, error) {
	return s.repo.LoadLine(you_id)
}
func (s *UserProfileService) CheckFollow(you_id, id_user int) (bool, error) {
	return s.repo.CheckFollow(you_id, id_user)
}

func (s *UserProfileService) LoadFavorite(you_id int) ([]mybox.Item, error) {
	return s.repo.LoadFavorite(you_id)
}

func (s *UserProfileService) GetNotices(you_id int) ([]mybox.Notice, error) {
	return s.repo.GetNotices(you_id)
}

func (s *UserProfileService) NoticeCheck(you_id int) error {
	return s.repo.NoticeCheck(you_id)
}

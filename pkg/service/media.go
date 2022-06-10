package service

import (
	"github.com/whyvilos/mybox/pkg/repository"
)

type MediaService struct {
	repo repository.Media
}

func NewMediaService(repo repository.Media) *MediaService {
	return &MediaService{repo: repo}
}

func (r *MediaService) SaveUrlAvatar(userId int, path string) error {
	return r.repo.SaveUrlAvatar(userId, path)
}

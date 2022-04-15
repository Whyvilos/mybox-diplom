package service

import (
	"github.com/whyvilos/mybox"
	"github.com/whyvilos/mybox/pkg/repository"
)

type FeedService struct {
	repo repository.Feed
}

func NewFeedService(repo repository.Feed) *FeedService {
	return &FeedService{repo: repo}
}

func (s *FeedService) CreatePost(id_user int, post mybox.Post) (int, error) {
	return s.repo.CreatePost(id_user, post)
}

func (s *FeedService) GetAll(id_feed int) ([]mybox.Post, error) {
	return s.repo.GetAll(id_feed)
}

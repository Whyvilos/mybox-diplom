package service

import (
	"errors"

	"github.com/whyvilos/mybox"
	"github.com/whyvilos/mybox/pkg/repository"
)

type CharService struct {
	repo repository.Chat
}

func NewCharService(repo repository.Chat) *CharService {
	return &CharService{repo: repo}
}

func (s *CharService) CreateChat(you_id, id_order int, status string) (int, error) {
	return s.repo.CreateChat(you_id, id_order, status)
}

func (s *CharService) SendMassage(you_id, id_chat int, input mybox.Messaage) (int, error) {
	return s.repo.SendMassage(you_id, id_chat, input)
}

func (s *CharService) GetAllMessage(you_id, id_chat int) (mybox.AllMessages, error) {
	var allMessages mybox.AllMessages
	flag, err := s.repo.CheckYouInChat(you_id, id_chat)
	if err != nil {
		return allMessages, err
	}
	if flag {
		users, err := s.repo.GetUserInChat(id_chat)
		if err != nil {
			return allMessages, err
		}
		messages, err := s.repo.GetMessages(id_chat)
		if err != nil {
			return allMessages, err
		}
		allMessages.Users = users
		allMessages.Massages = messages
		return allMessages, err
	}
	return allMessages, errors.New("вы не в этом чате")
}

func (s *CharService) FindChat(you_id, id_order int) (int, error) {
	return s.repo.FindChat(you_id, id_order)
}
func (s *CharService) FindChat2(you_id, id_order int) (int, error) {
	return s.repo.FindChat2(you_id, id_order)
}

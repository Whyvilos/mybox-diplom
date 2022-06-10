package service

import (
	"github.com/whyvilos/mybox"
	"github.com/whyvilos/mybox/pkg/repository"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (r *OrderService) CreateOrder(you_id int, input mybox.Order) (int, error) {
	return r.repo.CreateOrder(you_id, input)
}

func (r *OrderService) GetOrders(you_id int) ([]mybox.OrderResponce, error) {
	return r.repo.GetOrders(you_id)
}

func (r *OrderService) GetOrdersForYou(you_id int) ([]mybox.OrderResponce, error) {
	return r.repo.GetOrdersForYou(you_id)
}

func (r *OrderService) UpdateOrderStatus(you_id, id_order int, status string) error {
	return r.repo.UpdateOrderStatus(you_id, id_order, status)
}

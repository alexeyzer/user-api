package service

import (
	"context"
	"github.com/alexeyzer/user-api/internal/pkg/datastruct"
	"github.com/alexeyzer/user-api/internal/pkg/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type OrderService interface {
	CreateOrder(ctx context.Context, userID int64) (*datastruct.Order, error)
	ListOrders(ctx context.Context, userID int64) ([]*datastruct.Order, error)
}

type orderService struct {
	dao         repository.DAO
	cartService CartService
}

func (s *orderService) CreateOrder(ctx context.Context, userID int64) (*datastruct.Order, error) {
	cartItems, err := s.cartService.ListCartItems(ctx, userID)
	if err != nil {
		return nil, err
	}
	if len(cartItems) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Cart is empty, add products to cart to create order")
	}
	order := datastruct.Order{
		UserID:      userID,
		OrderStatus: datastruct.OrderStatus_CREATED,
		OrderDate:   time.Now().UTC(),
		TotalPrice:  0,
		Items:       make([]*datastruct.FullCartItem, 0, len(cartItems)),
	}
	for _, item := range cartItems {
		order.TotalPrice = order.TotalPrice + (item.FullFinalProduct.Price * float64(item.UserQuantity))
		order.Items = append(order.Items, item)
	}
	resp, err := s.dao.OrderQuery().Create(ctx, order)
	if err != nil {
		return nil, err
	}

	err = s.dao.CartQuery().DeleteByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *orderService) ListOrders(ctx context.Context, userID int64) ([]*datastruct.Order, error) {
	resp, err := s.dao.OrderQuery().List(ctx, userID)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func NewOrderService(dao repository.DAO, cartService CartService) OrderService {
	return &orderService{
		dao:         dao,
		cartService: cartService,
	}
}

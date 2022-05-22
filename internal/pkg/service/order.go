package service

import (
	"context"
	"github.com/alexeyzer/user-api/internal/client"
	"github.com/alexeyzer/user-api/internal/pkg/datastruct"
	"github.com/alexeyzer/user-api/internal/pkg/repository"
	pb "github.com/alexeyzer/user-api/pb/api/product/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

type OrderService interface {
	CreateOrder(ctx context.Context, userID int64) (*datastruct.Order, error)
	ListOrders(ctx context.Context, userID int64) ([]*datastruct.Order, error)
}

type orderService struct {
	dao              repository.DAO
	cartService      CartService
	productAPIClient client.ProductAPIClient
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
	updateReq := &pb.BatchUpdateFinalProductRequest{
		Items: make([]*pb.BatchUpdateFinalProductRequest_Item, 0, len(cartItems)),
	}
	for _, item := range cartItems {
		if item.UserQuantity > item.FullFinalProduct.Amount {
			return nil, status.Errorf(codes.InvalidArgument, "Вы патаетесь оформить больше товара: %s  чем есть в наличии", item.FullFinalProduct.Name)
		}
		updateReq.Items = append(updateReq.Items, &pb.BatchUpdateFinalProductRequest_Item{
			Id:     item.FullFinalProduct.ID,
			Amount: item.FullFinalProduct.Amount - item.UserQuantity,
		})
		order.TotalPrice = order.TotalPrice + (item.FullFinalProduct.Price * float64(item.UserQuantity))
		order.Items = append(order.Items, item)
	}
	err = s.productAPIClient.BatchUpdateFinalProduct(ctx, updateReq)
	if err != nil {
		return nil, err
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

func NewOrderService(dao repository.DAO, cartService CartService, productAPIClient client.ProductAPIClient) OrderService {
	return &orderService{
		dao:              dao,
		cartService:      cartService,
		productAPIClient: productAPIClient,
	}
}

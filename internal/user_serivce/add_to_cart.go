package user_serivce

import (
	"context"
	"github.com/alexeyzer/user-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserApiServiceServer) AddToCart(ctx context.Context, req *desc.AddToCartRequest) (*desc.AddToCartResponse, error) {
	resp, err := s.SessionCheck(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	cartItem, err := s.cartService.AddToCart(ctx, datastruct.CartItem{
		UserID:         resp.UserId,
		FinalProductID: req.FinalProductId,
		Quantity:       req.Quantity,
	})
	if err != nil {
		return nil, err
	}
	return &desc.AddToCartResponse{
		Id:             cartItem.ID,
		UserId:         cartItem.UserID,
		FinalProductId: cartItem.FinalProductID,
		Quantity:       cartItem.Quantity,
	}, nil
}

package user_serivce

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
)

func (s *UserApiServiceServer) UpdateCart(ctx context.Context, req *desc.UpdateCartRequest) (*desc.UpdateCartResponse, error) {
	user, err := s.SessionCheck(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	resp, err := s.cartService.UpdateCartItem(ctx, req.GetFinalProductId(), user.GetUserId(), req.GetQuantity())
	if err != nil {
		return nil, err
	}
	return &desc.UpdateCartResponse{
			FinalProductId: resp.FinalProductID,
			Quantity:       resp.Quantity,
		},
		nil
}

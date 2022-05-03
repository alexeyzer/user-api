package user_serivce

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserApiServiceServer) DeleteAllFromCart(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	resp, err := s.SessionCheck(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	err = s.cartService.DeleteAllFromCart(ctx, resp.GetUserId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

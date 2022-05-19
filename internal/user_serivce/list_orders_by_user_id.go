package user_serivce

import (
	"context"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
)

func (s *UserApiServiceServer) ListOrderByUserId(ctx context.Context, req *desc.ListOrderByUserIdRequest) (*desc.ListOrderResponse, error) {
	resp, err := s.orderService.ListOrders(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	return s.datastructOrdersToProto(resp), nil
}

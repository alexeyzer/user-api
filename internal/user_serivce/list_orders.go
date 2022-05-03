package user_serivce

import (
	"context"
	"github.com/alexeyzer/user-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserApiServiceServer) ListOrder(ctx context.Context, _ *emptypb.Empty) (*desc.ListOrderResponse, error) {
	userInfo, err := s.SessionCheck(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	resp, err := s.orderService.ListOrders(ctx, userInfo.GetUserId())
	if err != nil {
		return nil, err
	}
	return s.datastructOrdersToProto(resp), nil
}

func (s *UserApiServiceServer) datastructOrdersToProto(resp []*datastruct.Order) *desc.ListOrderResponse {
	internalResp := &desc.ListOrderResponse{
		Orders: make([]*desc.CreateOrderResponse, 0, len(resp)),
	}
	for _, order := range resp {
		internalResp.Orders = append(internalResp.Orders, s.datastructOrderToProto(order))
	}
	return internalResp
}

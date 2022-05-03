package user_serivce

import (
	"context"
	"github.com/alexeyzer/user-api/internal/pkg/datastruct"
	"google.golang.org/protobuf/types/known/timestamppb"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
)

func (s *UserApiServiceServer) CreateOrder(ctx context.Context, _ *emptypb.Empty) (*desc.CreateOrderResponse, error) {
	userInfo, err := s.SessionCheck(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	resp, err := s.orderService.CreateOrder(ctx, userInfo.GetUserId())
	if err != nil {
		return nil, err
	}
	return s.datastructOrderToProto(resp), nil
}

func (s *UserApiServiceServer) datastructOrderToProto(resp *datastruct.Order) *desc.CreateOrderResponse {
	internalResp := &desc.CreateOrderResponse{
		Id:          resp.ID,
		UserId:      resp.UserID,
		OrderStatus: desc.OrderStatus(desc.OrderStatus_value[string(resp.OrderStatus)]),
		OrderDate:   timestamppb.New(resp.OrderDate),
		TotalPrice:  resp.TotalPrice,
		Products:    make([]*desc.FullCartItem, 0, len(resp.Items)),
	}
	for _, item := range resp.Items {
		internalResp.Products = append(internalResp.Products, &desc.FullCartItem{
			FullProductId: item.FullFinalProduct.ID,
			Name:          item.FullFinalProduct.Name,
			Description:   item.FullFinalProduct.Description,
			Url:           item.FullFinalProduct.Url,
			BrandName:     item.FullFinalProduct.BrandName,
			CategoryName:  item.FullFinalProduct.CategoryName,
			Color:         item.FullFinalProduct.Color,
			Price:         item.FullFinalProduct.Price,
			Size:          item.FullFinalProduct.Size,
			Amount:        item.FullFinalProduct.Amount,
			UserQuantity:  item.UserQuantity,
			Id:            item.ID,
		})
	}
	return internalResp
}

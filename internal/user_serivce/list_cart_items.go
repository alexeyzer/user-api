package user_serivce

import (
	"context"
	"github.com/alexeyzer/user-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserApiServiceServer) ListCartItems(ctx context.Context, _ *emptypb.Empty) (*desc.ListCartItemsResponse, error) {
	resp, err := s.SessionCheck(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	cartItems, err := s.cartService.ListCartItems(ctx, resp.GetUserId())
	if err != nil {
		return nil, err
	}
	return s.datastructFullCartItemMapToProto(cartItems), nil
}

func (s *UserApiServiceServer) datastructFullCartItemMapToProto(resp map[int64]*datastruct.FullCartItem) *desc.ListCartItemsResponse {
	internalResp := &desc.ListCartItemsResponse{
		Products: make([]*desc.FullCartItem, 0, len(resp)),
	}
	totalPrice := float64(0)
	totalCount := int64(0)
	for _, fullCartItem := range resp {
		internalResp.Products = append(internalResp.Products, &desc.FullCartItem{
			FullProductId: fullCartItem.FullFinalProduct.ID,
			Name:          fullCartItem.FullFinalProduct.Name,
			Description:   fullCartItem.FullFinalProduct.Description,
			Url:           fullCartItem.FullFinalProduct.Url,
			BrandName:     fullCartItem.FullFinalProduct.BrandName,
			CategoryName:  fullCartItem.FullFinalProduct.CategoryName,
			Color:         fullCartItem.FullFinalProduct.Color,
			Price:         fullCartItem.FullFinalProduct.Price,
			Size:          fullCartItem.FullFinalProduct.Size,
			Amount:        fullCartItem.FullFinalProduct.Amount,
			UserQuantity:  fullCartItem.UserQuantity,
			Id:            fullCartItem.ID,
		})
		totalPrice = totalPrice + fullCartItem.FullFinalProduct.Price*float64(fullCartItem.UserQuantity)
		totalCount = totalCount + fullCartItem.UserQuantity
	}
	internalResp.TotalPrice = totalPrice
	internalResp.TotalCountProducts = totalCount

	return internalResp
}

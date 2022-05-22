package user_serivce

import (
	"context"
	"github.com/alexeyzer/user-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserApiServiceServer) ListFavorite(ctx context.Context, _ *emptypb.Empty) (*desc.ListFavoriteResponse, error) {
	user, err := s.SessionCheck(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	resp, err := s.favoriteService.ListFavorite(ctx, user.GetUserId())
	if err != nil {
		return nil, err
	}
	return s.productsToListFavoriteResponse(resp), nil
}

func (s *UserApiServiceServer) productsToListFavoriteResponse(resp map[int64]*datastruct.FavoriteProduct) *desc.ListFavoriteResponse {
	internalResp := &desc.ListFavoriteResponse{
		Products: make([]*desc.ProductResponse, 0, len(resp)),
	}
	for id, item := range resp {
		internalResp.Products = append(internalResp.Products, &desc.ProductResponse{
			Id:          id,
			FavoriteId:  item.FavoriteID,
			Name:        item.Name,
			Description: item.Description,
			Url:         item.Url,
			BrandId:     item.BrandID,
			CategoryId:  item.CategoryID,
			Color:       item.Color,
			Price:       item.Price,
		})
	}
	return internalResp
}

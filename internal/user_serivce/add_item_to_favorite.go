package user_serivce

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
)

func (s *UserApiServiceServer) AddItemToFavorite(ctx context.Context, req *desc.AddItemToFavoriteRequest) (*desc.AddItemToFavoriteResponse, error) {
	user, err := s.SessionCheck(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	resp, err := s.favoriteService.AddFavoriteItem(ctx, user.GetUserId(), req.GetProductId())
	if err != nil {
		return nil, err
	}
	return &desc.AddItemToFavoriteResponse{
		Id:        resp.ID,
		ProductId: resp.ProductID,
		UserId:    resp.UserID,
	}, nil
}

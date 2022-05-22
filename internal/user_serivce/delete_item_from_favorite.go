package user_serivce

import (
	"context"

	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserApiServiceServer) DeleteItemFromFavorite(ctx context.Context, req *desc.DeleteItemFromFavoriteRequest) (*emptypb.Empty, error) {
	err := s.favoriteService.DeleteFavoriteItem(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

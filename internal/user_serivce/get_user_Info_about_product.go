package user_serivce

import (
	"context"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserApiServiceServer) GetUserInfoAboutProduct(ctx context.Context, req *desc.GetUserInfoAboutProductRequest) (*desc.GetUserInfoAboutProductResponse, error) {
	isFavorite := false
	favoriteID := int64(0)
	userQuantity := int64(0)

	user, err := s.SessionCheck(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}
	resp, err := s.cartService.ListCartItems(ctx, user.GetUserId())
	if err != nil {
		return nil, err
	}
	if product, ok := resp[req.GetProductId()]; ok {
		userQuantity = product.UserQuantity
	}
	favorites, err := s.favoriteService.ListFavorite(ctx, user.GetUserId())
	if err != nil {
		return nil, err
	}
	if favorite, ok := favorites[req.GetProductId()]; ok {
		isFavorite = true
		favoriteID = favorite.FavoriteID
	}

	return &desc.GetUserInfoAboutProductResponse{
		UserQuantity: userQuantity,
		IsFavorite:   isFavorite,
		FavoriteId:   favoriteID,
	}, nil
}

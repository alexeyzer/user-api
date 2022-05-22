package service

import (
	"context"
	"github.com/alexeyzer/user-api/internal/client"
	"github.com/alexeyzer/user-api/internal/pkg/datastruct"
	"github.com/alexeyzer/user-api/internal/pkg/repository"
)

type FavoriteService interface {
	DeleteFavoriteItem(ctx context.Context, id int64) error
	AddFavoriteItem(ctx context.Context, userID, productID int64) (*datastruct.FavoriteItem, error)
	ListFavorite(ctx context.Context, userID int64) (map[int64]*datastruct.FavoriteProduct, error)
}

type favoriteService struct {
	dao              repository.DAO
	productAPIClient client.ProductAPIClient
}

func (s *favoriteService) DeleteFavoriteItem(ctx context.Context, id int64) error {
	err := s.dao.FavoriteQuery().Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *favoriteService) AddFavoriteItem(ctx context.Context, userID, productID int64) (*datastruct.FavoriteItem, error) {
	resp, err := s.dao.FavoriteQuery().Create(ctx, datastruct.FavoriteItem{
		UserID:    userID,
		ProductID: productID,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *favoriteService) ListFavorite(ctx context.Context, userID int64) (map[int64]*datastruct.FavoriteProduct, error) {
	resp, err := s.dao.FavoriteQuery().List(ctx, userID)
	if err != nil {
		return nil, err
	}
	ids := make([]int64, 0, len(resp))
	result := make(map[int64]*datastruct.FavoriteProduct, len(resp))
	for _, item := range resp {
		result[item.ProductID] = &datastruct.FavoriteProduct{
			FavoriteID: item.ID,
		}
		ids = append(ids, item.ProductID)
	}

	products, err := s.productAPIClient.ListProductsById(ctx, ids)
	if err != nil {
		return nil, err
	}
	for _, product := range products.GetProducts() {
		result[product.Id].Price = product.Price
		result[product.Id].Name = product.Name
		result[product.Id].Url = product.Url
		result[product.Id].Description = product.Url
		result[product.Id].CategoryID = product.CategoryId
		result[product.Id].BrandID = product.BrandId
		result[product.Id].Color = product.Color

	}

	return result, nil
}

func NewFavoriteService(dao repository.DAO, productAPIClient client.ProductAPIClient) FavoriteService {
	return &favoriteService{
		dao:              dao,
		productAPIClient: productAPIClient,
	}
}

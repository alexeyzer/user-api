package service

import (
	"context"
	"database/sql"
	"github.com/alexeyzer/user-api/internal/client"
	pb "github.com/alexeyzer/user-api/pb/api/product/v1"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/alexeyzer/user-api/internal/pkg/datastruct"
	"github.com/alexeyzer/user-api/internal/pkg/repository"
)

type CartService interface {
	AddToCart(ctx context.Context, req datastruct.CartItem) (*datastruct.CartItem, error)
	UpdateCartItem(ctx context.Context, finalProductID, userID, quantity int64) (*datastruct.CartItem, error)
	GetCartItem(ctx context.Context, ID int64) (*datastruct.CartItem, error)
	DeleteItemFromCart(ctx context.Context, ID, UserID int64) error
	DeleteAllFromCart(ctx context.Context, userID int64) error
	ListCartItems(ctx context.Context, userID int64) (map[int64]*datastruct.FullCartItem, error)
}

type cartService struct {
	dao              repository.DAO
	productAPIClient client.ProductAPIClient
}

func (s *cartService) UpdateCartItem(ctx context.Context, finalProductID, userID, quantity int64) (*datastruct.CartItem, error) {
	finalProduct, err := s.productAPIClient.GetFinalProduct(ctx, &pb.GetFinalProductRequest{Id: finalProductID})
	if err != nil {
		return nil, err
	}
	if finalProduct.GetAmount() < quantity {
		return nil, status.Errorf(codes.InvalidArgument, "Невозможно обновить, введено больше чем доступно")
	}

	resp, err := s.dao.CartQuery().Update(ctx, userID, finalProductID, quantity)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *cartService) DeleteAllFromCart(ctx context.Context, userID int64) error {
	err := s.dao.CartQuery().DeleteByUserID(ctx, userID)
	if err != nil {
		return err
	}
	return nil
}

func (s *cartService) ListCartItems(ctx context.Context, userID int64) (map[int64]*datastruct.FullCartItem, error) {
	res, err := s.dao.CartQuery().List(ctx, userID)
	if err != nil {
		return nil, err
	}
	if len(res) == 0 {
		return nil, nil
	}
	ids := make([]int64, 0, len(res))
	response := make(map[int64]*datastruct.FullCartItem, len(res))
	for _, cartItem := range res {
		ids = append(ids, cartItem.FinalProductID)
		response[cartItem.FinalProductID] = &datastruct.FullCartItem{
			UserQuantity: cartItem.Quantity,
			ID:           cartItem.ID,
		}
	}
	fullFinalProducts, err := s.productAPIClient.ListFullFinalProducts(ctx, &pb.ListFullFinalProductsRequest{
		Ids: ids,
	})
	if err != nil {
		return nil, err
	}
	for _, fullFinalProduct := range fullFinalProducts.Products {
		response[fullFinalProduct.Id].FullFinalProduct = datastruct.FullFinalProduct{
			ID:           fullFinalProduct.GetId(),
			Amount:       fullFinalProduct.GetAmount(),
			Sku:          fullFinalProduct.GetSku(),
			Name:         fullFinalProduct.GetName(),
			Description:  fullFinalProduct.GetDescription(),
			Url:          fullFinalProduct.GetUrl(),
			BrandName:    fullFinalProduct.GetBrandName(),
			CategoryName: fullFinalProduct.GetCategoryName(),
			Price:        fullFinalProduct.GetPrice(),
			Color:        fullFinalProduct.GetColor(),
			Size:         fullFinalProduct.GetSize(),
		}
	}

	return response, nil
}

func (s *cartService) AddToCart(ctx context.Context, req datastruct.CartItem) (*datastruct.CartItem, error) {
	_, err := s.productAPIClient.GetFinalProduct(ctx, &pb.GetFinalProductRequest{Id: req.FinalProductID})
	if err != nil {
		return nil, err
	}

	exists, err := s.dao.CartQuery().Exists(ctx, req.UserID, req.FinalProductID)
	if err != nil {
		return nil, err
	}
	if exists == true {
		return nil, status.Errorf(codes.InvalidArgument, "cartItem with user_id = %d and final_product_id = %d already exists", req.UserID, req.FinalProductID)
	}

	res, err := s.dao.CartQuery().Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *cartService) GetCartItem(ctx context.Context, ID int64) (*datastruct.CartItem, error) {
	res, err := s.dao.CartQuery().Get(ctx, ID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *cartService) DeleteItemFromCart(ctx context.Context, ID, userID int64) error {
	cartItem, err := s.dao.CartQuery().Get(ctx, ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.InvalidArgument, "cartItem with id = %d doesn't exist", ID)
		}
		return err
	}
	if cartItem.UserID != userID {
		return status.Error(codes.PermissionDenied, "failed to delete cartItem")
	}

	err = s.dao.CartQuery().Delete(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}

func NewCartService(dao repository.DAO, productAPIClient client.ProductAPIClient) CartService {
	return &cartService{
		dao:              dao,
		productAPIClient: productAPIClient,
	}
}

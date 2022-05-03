package client

import (
	"context"
	pb "github.com/alexeyzer/user-api/pb/api/product/v1"
	"google.golang.org/grpc"
)

type ProductAPIClient interface {
	GetFinalProduct(ctx context.Context, req *pb.GetFinalProductRequest) (*pb.GetFinalProductResponse, error)
	ListFullFinalProducts(ctx context.Context, req *pb.ListFullFinalProductsRequest) (*pb.ListFullFinalProductsResponse, error)
}

type productAPIClient struct {
	conn   *grpc.ClientConn
	client pb.ProductApiServiceClient
}

func (c *productAPIClient) GetFinalProduct(ctx context.Context, req *pb.GetFinalProductRequest) (*pb.GetFinalProductResponse, error) {
	res, err := c.client.GetFinalProduct(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (c *productAPIClient) ListFullFinalProducts(ctx context.Context, req *pb.ListFullFinalProductsRequest) (*pb.ListFullFinalProductsResponse, error) {
	res, err := c.client.ListFullFinalProducts(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewProductApiClient(address string) (ProductAPIClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := pb.NewProductApiServiceClient(conn)

	client := &productAPIClient{
		conn:   conn,
		client: c,
	}
	return client, nil
}

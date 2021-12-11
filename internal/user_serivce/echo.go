package user_serivce

import (
	"context"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
)

func (UserApiServiceServer) Echo(ctx context.Context, req *desc.EchoRequest) (*desc.EchoResponse, error) {
	return &desc.EchoResponse{
		Message: req.Message + " works!",
	}, nil
}

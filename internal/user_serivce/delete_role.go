package user_serivce

import (
	"context"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserApiServiceServer) DeleteRole(ctx context.Context, req *desc.DeleteRoleRequest) (*emptypb.Empty, error) {
	err := s.roleService.DeleteRole(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

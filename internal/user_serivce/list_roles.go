package user_serivce

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/alexeyzer/user-api/internal/pkg/datastruct"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
)

func (s *UserApiServiceServer) ListRoles(ctx context.Context, _ *emptypb.Empty) (*desc.ListRolesResponse, error) {
	res, err := s.roleService.ListRoles(ctx)
	if err != nil {
		return nil, err
	}
	return s.rolesToProtoListRolesResponse(res), nil
}

func (s *UserApiServiceServer) rolesToProtoListRolesResponse(resp []*datastruct.Role) *desc.ListRolesResponse {
	internalResp := &desc.ListRolesResponse{
		Roles: make([]*desc.GetRoleResponse, 0, len(resp)),
	}
	for _, item := range resp {
		internalResp.Roles = append(internalResp.Roles, s.roleToProtoGetRoleResponse(item))
	}
	return internalResp
}

package user_serivce

import (
	"context"

	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
)

func (s *UserApiServiceServer) ListUserRoles(ctx context.Context, req *desc.ListUserRolesRequest) (*desc.ListUserRolesResponse, error) {
	resp, err := s.userRoleService.ListUserRoles(ctx, req.GetUserId())
	if err != nil {
		return nil, err
	}
	return s.userRolesToProtoListUserRolesResponse(resp), nil
}

func (s *UserApiServiceServer) userRolesToProtoListUserRolesResponse(resp []string) *desc.ListUserRolesResponse {
	internalResp := &desc.ListUserRolesResponse{
		UserRoles: make([]string, 0, len(resp)),
	}
	for _, item := range resp {
		internalResp.UserRoles = append(internalResp.UserRoles, item)
	}
	return internalResp
}

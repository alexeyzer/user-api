package service

import (
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/alexeyzer/user-api/internal/pkg/datastruct"
	"github.com/alexeyzer/user-api/internal/pkg/repository"
)

type RoleService interface {
	CreateRole(ctx context.Context, req datastruct.Role) (*datastruct.Role, error)
	UpdateRole(ctx context.Context, req datastruct.Role) (*datastruct.Role, error)
	GetRole(ctx context.Context, ID int64) (*datastruct.Role, error)
	DeleteRole(ctx context.Context, ID int64) error
	ListRoles(ctx context.Context) ([]*datastruct.Role, error)
}

type roleService struct {
	dao repository.DAO
}

func (s *roleService) UpdateRole(ctx context.Context, req datastruct.Role) (*datastruct.Role, error) {
	_, err := s.dao.RoleQuery().Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}

	resp, err := s.dao.RoleQuery().Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *roleService) ListRoles(ctx context.Context) ([]*datastruct.Role, error) {
	res, err := s.dao.RoleQuery().List(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *roleService) CreateRole(ctx context.Context, req datastruct.Role) (*datastruct.Role, error) {
	exists, err := s.dao.RoleQuery().Exists(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	if exists == true {
		return nil, status.Errorf(codes.InvalidArgument, "role with name = %s already exists", req.Name)
	}

	res, err := s.dao.RoleQuery().Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *roleService) GetRole(ctx context.Context, ID int64) (*datastruct.Role, error) {
	res, err := s.dao.RoleQuery().Get(ctx, ID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *roleService) DeleteRole(ctx context.Context, ID int64) error {
	_, err := s.dao.RoleQuery().Get(ctx, ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.InvalidArgument, "Role with id = %d doesn't exist", ID)
		}
		return err
	}
	err = s.dao.RoleQuery().Delete(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}

func NewRoleService(dao repository.DAO) RoleService {
	return &roleService{dao: dao}
}

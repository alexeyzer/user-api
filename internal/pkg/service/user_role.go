package service

import (
	"context"
	"database/sql"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/alexeyzer/user-api/internal/pkg/datastruct"
	"github.com/alexeyzer/user-api/internal/pkg/repository"
)

type UserRoleService interface {
	CreateUserRole(ctx context.Context, req datastruct.UserRole) (*datastruct.UserRole, error)
	GetUserRole(ctx context.Context, ID int64) (*datastruct.UserRole, error)
	DeleteUserRole(ctx context.Context, ID int64) error
	ListUserRoles(ctx context.Context, userID int64) ([]*datastruct.UserRoleWithName, error)
}

type userRoleService struct {
	dao repository.DAO
}

func (s *userRoleService) ListUserRoles(ctx context.Context, userID int64) ([]*datastruct.UserRoleWithName, error) {
	res, err := s.dao.UserRoleQuery().List(ctx, userID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *userRoleService) CreateUserRole(ctx context.Context, req datastruct.UserRole) (*datastruct.UserRole, error) {
	exists, err := s.dao.UserRoleQuery().Exists(ctx, req.UserID, req.RoleID)
	if err != nil {
		return nil, err
	}
	if exists == true {
		return nil, status.Errorf(codes.InvalidArgument, "userRole with userID = %d and roleID = %d  already exists", req.UserID, req.RoleID)
	}

	res, err := s.dao.UserRoleQuery().Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *userRoleService) GetUserRole(ctx context.Context, ID int64) (*datastruct.UserRole, error) {
	res, err := s.dao.UserRoleQuery().Get(ctx, ID)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *userRoleService) DeleteUserRole(ctx context.Context, ID int64) error {
	_, err := s.dao.UserRoleQuery().Get(ctx, ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return status.Errorf(codes.InvalidArgument, "userRole with id = %d doesn't exist", ID)
		}
		return err
	}
	err = s.dao.UserRoleQuery().Delete(ctx, ID)
	if err != nil {
		return err
	}
	return nil
}

func NewUserRoleService(dao repository.DAO) UserRoleService {
	return &userRoleService{dao: dao}
}

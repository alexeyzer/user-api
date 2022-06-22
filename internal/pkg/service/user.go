package service

import (
	"context"
	"github.com/alexeyzer/user-api/config"
	"github.com/alexeyzer/user-api/internal/client"
	"github.com/alexeyzer/user-api/internal/pkg/datastruct"
	"github.com/alexeyzer/user-api/internal/pkg/repository"
	desc "github.com/alexeyzer/user-api/pb/api/user/v1"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type UserService interface {
	CreateUser(ctx context.Context, req *desc.CreateUserRequest) (*datastruct.User, error)
	Login(ctx context.Context, req *desc.LoginRequest) ([]*datastruct.UserRoleWithName, string, *datastruct.User, error)
	SessionCheck(ctx context.Context, sessionID string) (*datastruct.UserWithRoles, error)
	DeleteSession(ctx context.Context, sessionID string) error
	GetUser(ctx context.Context, ID int64) (*datastruct.User, []*datastruct.UserRoleWithName, error)
	ListUsers(ctx context.Context) ([]*datastruct.User, error)
}

type userService struct {
	dao   repository.DAO
	redis client.RedisClient
}

func (s *userService) GetUser(ctx context.Context, ID int64) (*datastruct.User, []*datastruct.UserRoleWithName, error) {
	resp, err := s.dao.UserQuery().GetByID(ctx, ID)
	if err != nil {
		return nil, nil, err
	}
	roles, err := s.dao.UserRoleQuery().List(ctx, resp.ID)
	if err != nil {
		return nil, nil, err
	}
	return resp, roles, nil
}

func (s *userService) ListUsers(ctx context.Context) ([]*datastruct.User, error) {
	resp, err := s.dao.UserQuery().List(ctx)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *userService) createSession(ctx context.Context, user *datastruct.User) string {
	sessionID := uuid.New().String()
	err := s.redis.Set(ctx, sessionID, user.Email)
	if err != nil {
		log.Warnf("failed to create sessionID for user = %s, err: %s", user.Email, err)
	}

	return sessionID
}

func (s *userService) DeleteSession(ctx context.Context, sessionID string) error {
	err := s.redis.Delete(ctx, sessionID)
	if err != nil {
		return err
	}
	if err := grpc.SetHeader(ctx, metadata.Pairs(config.Config.Auth.SessionKey, sessionID)); err != nil {
		return err
	}
	if err := grpc.SetHeader(ctx, metadata.Pairs(config.Config.Auth.LogoutKey, config.Config.Auth.LogoutKey)); err != nil {
		return err
	}
	return nil
}

func (s *userService) SessionCheck(ctx context.Context, sessionID string) (*datastruct.UserWithRoles, error) {
	email, err := s.redis.Get(ctx, sessionID)
	if err != nil {
		return nil, err
	}
	user, err := s.dao.UserQuery().Get(ctx, email)
	if err != nil {
		return nil, err
	}
	roles, err := s.dao.UserRoleQuery().List(ctx, user.ID)
	if err != nil {
		return nil, err
	}
	rolesName := make([]string, 0, len(roles))
	for _, role := range roles {
		rolesName = append(rolesName, role.RoleName)
	}

	return &datastruct.UserWithRoles{
		ID:    user.ID,
		Email: user.Email,
		Roles: rolesName,
	}, nil
}

func (s *userService) Login(ctx context.Context, req *desc.LoginRequest) ([]*datastruct.UserRoleWithName, string, *datastruct.User, error) {

	user, err := s.dao.UserQuery().Get(ctx, req.Email)
	if err != nil {
		return nil, "", nil, err
	}
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(req.Password))
	if err != nil {
		return nil, "", nil, status.Errorf(codes.InvalidArgument, "Invalid password for user with email = %s", req.Email)
	}

	sessionID := s.createSession(ctx, user)
	extractedRoles, err := s.dao.UserRoleQuery().List(ctx, user.ID)
	if err != nil {
		return nil, "", nil, err
	}

	return extractedRoles, sessionID, user, nil
}

func (s *userService) CreateUser(ctx context.Context, req *desc.CreateUserRequest) (*datastruct.User, error) {
	exists, err := s.dao.UserQuery().Exists(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if exists == true {
		return nil, status.Errorf(codes.InvalidArgument, "User with email = %s already exists", req.Email)
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	res, err := s.dao.UserQuery().Create(ctx, s.serviceUserReqToDaoUser(req, encryptedPassword))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *userService) serviceUserReqToDaoUser(req *desc.CreateUserRequest, password []byte) datastruct.User {
	return datastruct.User{
		Name:       req.Name,
		Password:   password,
		Surname:    req.Surname,
		Patronymic: req.Patronymic,
		Phone:      req.Phone,
		Email:      req.Email,
	}
}

func NewUserService(dao repository.DAO, redis client.RedisClient) UserService {
	return &userService{dao: dao, redis: redis}
}

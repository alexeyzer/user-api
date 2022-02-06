package service

import (
	"context"
	"github.com/alexeyzer/user-api/config"
	"github.com/alexeyzer/user-api/internal/client"
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
	CreateUser(ctx context.Context, req *desc.CreateUserRequest) (*repository.User, error)
	GetBySession(ctx context.Context, sessionID string) (*repository.User, error)
	Login(ctx context.Context, req *desc.LoginRequest) (*string, error)
	SessionCheck(ctx context.Context, sessionID string) (*string, error)
	DeleteSession(ctx context.Context, sessionID string) error
}

type userService struct {
	dao   repository.DAO
	redis client.RedisClient
}

func (s *userService) GetBySession(ctx context.Context, sessionID string) (*repository.User, error) {
	email, err := s.SessionCheck(ctx, sessionID)
	if err != nil {
		return nil, err
	}
	user, err := s.dao.UserQuery().Get(ctx, *email)
	if err != nil {
		return nil, err
	}
	return user, nil
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

func (s *userService) SessionCheck(ctx context.Context, sessionID string) (*string, error) {

	email, err := s.redis.Get(ctx, sessionID)
	if err != nil {
		return nil, err
	}
	return &email, nil
}

func (s *userService) Login(ctx context.Context, req *desc.LoginRequest) (*string, error) {
	exists, err := s.dao.UserQuery().Exists(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if exists == false {
		return nil, status.Errorf(codes.InvalidArgument, "User with username = %s doesn't exist", req.Email)
	}
	user, err := s.dao.UserQuery().Get(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword(user.Password, []byte(req.Password))
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid password for username = %s", req.Email)
	}

	sessionID := uuid.New().String()
	err = s.redis.Set(ctx, sessionID, user.Email)
	if err != nil {
		log.Warnf("failed to create sessionID for user = %s", user.Email)
	}
	if err := grpc.SetHeader(ctx, metadata.Pairs(config.Config.Auth.SessionKey, sessionID)); err != nil {
		return nil, err
	}
	return &sessionID, nil
}

func (s *userService) CreateUser(ctx context.Context, req *desc.CreateUserRequest) (*repository.User, error) {
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
	sessionID := uuid.New().String()
	err = s.redis.Set(ctx, sessionID, res.Email)
	if err != nil {
		log.Warnf("failed to create sessionID for user = %s", res.Email)
	}

	return res, nil
}

func (s *userService) serviceUserReqToDaoUser(req *desc.CreateUserRequest, password []byte) repository.User {
	return repository.User{
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

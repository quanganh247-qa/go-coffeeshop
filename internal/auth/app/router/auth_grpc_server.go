package router

import (
	"context"

	"github.com/google/uuid"
	"golang.org/x/exp/slog"

	"github.com/google/wire"
	"github.com/thangchung/go-coffeeshop/cmd/auth/config"
	usecases "github.com/thangchung/go-coffeeshop/internal/auth/usecases/users"
	gen "github.com/thangchung/go-coffeeshop/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type authGRPCService struct {
	gen.UnimplementedAuthServiceServer
	cf *config.Config
	us usecases.UseCase
}

var _ gen.AuthServiceServer = (*authGRPCService)(nil)

var AuthGRPCServiceSet = wire.NewSet(NewAuthGRPCService)

func NewAuthGRPCService(
	grpcServer *grpc.Server,
	cf *config.Config,
	us usecases.UseCase,
) gen.AuthServiceServer {
	svc := authGRPCService{
		cf: cf,
		us: us,
	}
	gen.RegisterAuthServiceServer(grpcServer, &svc)
	reflection.Register(grpcServer)
	return &svc
}

func (s *authGRPCService) Register(
	ctx context.Context,
	request *gen.RegisterRequest,
) (*gen.RegisterResponse, error) {
	slog.Info("POST: Register User")
	res := gen.RegisterResponse{}
	err := s.us.Register(ctx, usecases.UserRegistrationRequest{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *authGRPCService) Login(
	ctx context.Context,
	request *gen.LoginRequest,
) (*gen.LoginResponse, error) {
	slog.Info("POST: Login User")
	resp, err := s.us.Login(ctx, usecases.UserLoginRequest{
		Email:    request.Email,
		Password: request.Password,
	})
	if err != nil {
		return nil, err
	}
	return &gen.LoginResponse{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}, nil
}

func (s *authGRPCService) VerifyToken(
	ctx context.Context,
	request *gen.VerifyTokenRequest,
) (*gen.VerifyTokenResponse, error) {
	slog.Info("POST: Verify Token")
	valid, err := s.us.VerifyToken(ctx, request.Token)
	if err != nil {
		return nil, err
	}
	return &gen.VerifyTokenResponse{
		Valid: valid,
	}, nil
}

func (s *authGRPCService) ListUsers(
	ctx context.Context,
	request *gen.ListUsersRequest,
) (*gen.ListUsersResponse, error) {
	slog.Info("GET: List Users")
	users, err := s.us.List(ctx)
	if err != nil {
		return nil, err
	}

	resUsers := make([]*gen.UserDto, 0, len(users))
	for _, u := range users {
		resUsers = append(resUsers, &gen.UserDto{
			Id:            u.ID.String(),
			Email:         u.Email,
			IsActive:      u.IsActive,
			EmailVerified: u.EmailVerified,
			CreatedAt:     u.CreatedAt.String(),
			UpdatedAt:     u.UpdatedAt.String(),
		})
	}

	return &gen.ListUsersResponse{
		Users: resUsers,
	}, nil
}

func (s *authGRPCService) GetUserDetail(
	ctx context.Context,
	request *gen.GetUserDetailRequest,
) (*gen.GetUserDetailResponse, error) {
	slog.Info("GET: Get User Detail")
	id, err := uuid.Parse(request.Id)
	if err != nil {
		return nil, err
	}

	user, err := s.us.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return &gen.GetUserDetailResponse{
		User: &gen.UserDto{
			Id:            user.ID.String(),
			Email:         user.Email,
			IsActive:      user.IsActive,
			EmailVerified: user.EmailVerified,
			CreatedAt:     user.CreatedAt.String(),
			UpdatedAt:     user.UpdatedAt.String(),
		},
	}, nil
}

func (s *authGRPCService) RefreshToken(
	ctx context.Context,
	request *gen.RefreshTokenRequest,
) (*gen.RefreshTokenResponse, error) {
	slog.Info("POST: Refresh Token")
	accessToken, err := s.us.RefreshToken(ctx, request.RefreshToken)
	if err != nil {
		return nil, err
	}
	return &gen.RefreshTokenResponse{
		AccessToken: accessToken,
	}, nil
}

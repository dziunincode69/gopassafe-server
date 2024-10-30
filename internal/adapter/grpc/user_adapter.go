package grpc

import (
	"context"

	userpb "github.com/dziunincode69/gopassafe-proto/protogen/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func validateUser(user *userpb.User) error {
	if user == nil {
		return status.Errorf(codes.InvalidArgument, "request missing required field: User")
	}
	if user.Email == "" {
		return status.Errorf(codes.InvalidArgument, "request missing required field: Email")
	}
	if user.Password == "" {
		return status.Errorf(codes.InvalidArgument, "request missing required field: Password")
	}
	return nil
}
func (g *GrpcAdapter) Login(ctx context.Context, req *userpb.LoginRequest) (*userpb.LoginResponse, error) {
	if err := validateUser(req.User); err != nil {
		return nil, err
	}
	token := g.userService.Authenticate(req.User.Email, req.User.Password)
	return &userpb.LoginResponse{
		User: &userpb.UserResponse{
			Id:    1,
			Email: req.User.Email,
			Token: token,
			Error: "",
		},
	}, nil
}

func (g *GrpcAdapter) Register(ctx context.Context, req *userpb.RegisterRequest) (*userpb.RegisterResponse, error) {
	if err := validateUser(req.User); err != nil {
		return nil, err
	}
	token := g.userService.Register(req.User.Email, req.User.Password)
	return &userpb.RegisterResponse{
		User: &userpb.UserResponse{
			Id:    1,
			Email: req.User.Email,
			Token: token,
			Error: "",
		},
	}, nil
}

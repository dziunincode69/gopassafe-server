package grpc

import (
	"context"

	"github.com/dziunincode69/gopassafe-proto/protogen/user"
)

func (g *GrpcAdapter) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	account := g.userService.GetUser(req.User.Email)

	return &user.LoginResponse{
		User: &user.UserResponse{
			Id:    1,
			Email: account,
			Token: "0000000000-000000-00000-00000",
			Error: "",
		},
	}, nil
}

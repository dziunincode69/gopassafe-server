package grpc

import (
	"context"

	vaultpb "github.com/dziunincode69/gopassafe-proto/protogen/vault"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (g *GrpcAdapter) VaultGet(ctx context.Context, req *vaultpb.VaultGetRequest) (*vaultpb.VaultGetResponse, error) {
	if req.Id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "request missing required field: Id")
	}
	GetVault := g.vaultService.GetVault(req.Id)
	Vaults := vaultpb.VaultGetResponse{
		Vaults: []*vaultpb.Vault{
			{
				Id:          1,
				Name:        GetVault,
				Description: GetVault,
				Url:         GetVault,
				Password:    GetVault,
			},
			{
				Id:          2,
				Name:        GetVault,
				Description: GetVault,
				Url:         GetVault,
				Password:    GetVault,
			},
		},
	}
	return &Vaults, nil

}

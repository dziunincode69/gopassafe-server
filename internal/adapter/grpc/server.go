package grpc

import (
	"fmt"
	"log"
	"net"

	userpb "github.com/dziunincode69/gopassafe-proto/protogen/user"
	vaultpb "github.com/dziunincode69/gopassafe-proto/protogen/vault"
	"github.com/dziunincode69/gopassafe-server/internal/port"
	"google.golang.org/grpc"
)

type GrpcAdapter struct {
	userService  port.UserServicePort
	vaultService port.VaultServicePort
	grpcPort     int
	server       *grpc.Server
	userpb.UnimplementedLoginServiceServer
	userpb.UnimplementedRegisterServiceServer
	vaultpb.UnimplementedVaultServiceServer
}

func NewGrpcAdapter(userService port.UserServicePort, vaultService port.VaultServicePort, grpcPort int) *GrpcAdapter {
	return &GrpcAdapter{
		userService:  userService,
		vaultService: vaultService,
		grpcPort:     grpcPort,
	}
}

func (g *GrpcAdapter) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", g.grpcPort))
	if err != nil {
		log.Fatalf("Failed to listen on port %d: %v", g.grpcPort, err)
	}
	log.Println("Starting gRPC server on port", g.grpcPort)

	grpcServer := grpc.NewServer()
	g.server = grpcServer

	userpb.RegisterLoginServiceServer(grpcServer, g)
	userpb.RegisterRegisterServiceServer(grpcServer, g)
	vaultpb.RegisterVaultServiceServer(grpcServer, g)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %d: %v", g.grpcPort, err)
	}
}

func (g *GrpcAdapter) Stop() {
	g.server.Stop()
}

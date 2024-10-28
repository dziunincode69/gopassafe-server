package grpc

import (
	"fmt"
	"log"
	"net"

	"github.com/dziunincode69/gopassafe-proto/protogen/user"
	"github.com/dziunincode69/gopassafe-server/internal/port"
	"google.golang.org/grpc"
)

type GrpcAdapter struct {
	userService port.UserServicePort
	grpcPort    int
	server      *grpc.Server
	user.LoginServiceServer
}

func NewGrpcAdapter(userService port.UserServicePort, grpcPort int) *GrpcAdapter {
	return &GrpcAdapter{
		userService: userService,
		grpcPort:    grpcPort,
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

	user.RegisterLoginServiceServer(grpcServer, g)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve gRPC server over port %d: %v", g.grpcPort, err)
	}
}

func (g *GrpcAdapter) Stop() {
	g.server.Stop()
}

package main

import (
	usergrpc "github.com/dziunincode69/gopassafe-server/internal/adapter/grpc"
	app "github.com/dziunincode69/gopassafe-server/internal/application"
)

func main() {
	us := &app.UserService{}
	grpcAdapter := usergrpc.NewGrpcAdapter(us, 9999)
	grpcAdapter.Run()
}

package main

import (
	usergrpc "github.com/dziunincode69/gopassafe-server/internal/adapter/grpc"
	app "github.com/dziunincode69/gopassafe-server/internal/application"
)

const PORT = 8899

func main() {
	userappservice := &app.UserService{}
	vaultappservice := &app.VaultService{}
	grpcAdapter := usergrpc.NewGrpcAdapter(userappservice, vaultappservice, PORT)
	grpcAdapter.Run()
}

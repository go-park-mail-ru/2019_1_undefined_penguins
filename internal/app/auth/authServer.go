package auth

import (
	"2019_1_undefined_penguins/internal/pkg/helpers"
	//"log"
	"net"

	"google.golang.org/grpc"
)

func Start() error {
	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		helpers.LogMsg("Can`t listen port ", err)
	}

	server := grpc.NewServer()

	RegisterAuthCheckerServer(server, NewAuthManager())

	helpers.LogMsg("AuthServer started at :8083")
	return server.Serve(lis)
}
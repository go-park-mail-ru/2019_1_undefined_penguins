package microChat

import (
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"google.golang.org/grpc"
	"net"
)

func Start() {
	lis, err := net.Listen("tcp", ":8083")
	if err != nil {
		helpers.LogMsg("cant listen port", err)
	}

	serverGrpc := grpc.NewServer()

	RegisterUserCheckerServer(serverGrpc, NewUserManager())

	helpers.LogMsg("starting ChatSserver at 8083")
	serverGrpc.Serve(lis)
}

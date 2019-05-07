package auth

import (
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"2019_1_undefined_penguins/internal/pkg/models"

	//"log"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func Start() error {

	viper.AddConfigPath("./configs")
	viper.SetConfigName("auth")
	var port string
	if err := viper.ReadInConfig(); err != nil {
		port = ":8083"
	} else {
		port = ":" + viper.GetString("port")
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		helpers.LogMsg("Can`t listen port ", err)
	}

	server := grpc.NewServer()

	models.RegisterAuthCheckerServer(server, NewAuthManager())

	helpers.LogMsg("AuthServer started at " + port)
	return server.Serve(lis)
}

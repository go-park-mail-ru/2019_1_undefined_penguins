package main

import (
	db "2019_1_undefined_penguins/internal/pkg/auth/database"
	"2019_1_undefined_penguins/internal/pkg/auth"
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"2019_1_undefined_penguins/internal/pkg/models"
	"net"

	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var SECRET = []byte("myawesomesecret")

func setConfig() string {
	viper.AddConfigPath("./configs")
	viper.SetConfigName("auth")
	var port string
	if err := viper.ReadInConfig(); err != nil {
		port = ":8083"
	} else {
		port = ":" + viper.GetString("port")
		SECRET = []byte(viper.GetString("secret"))
	}
	return port
}

func main() {
	port := setConfig()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		helpers.LogMsg("Can`t listen port ", err)
	}

	server := grpc.NewServer()

	models.RegisterAuthCheckerServer(server, auth.NewAuthManager())
	auth.UsersWantToPlay = make(map[string]*models.User)

	err = db.Connect()
	if err != nil {
		helpers.LogMsg("Connection error: ", err)
		return
	}
	defer db.Disconnect()

	helpers.LogMsg("AuthServer started at", port)
	server.Serve(lis)
}

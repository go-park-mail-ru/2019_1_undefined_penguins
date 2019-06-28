package main

import (
	"2019_1_undefined_penguins/internal/app/server"
	"2019_1_undefined_penguins/internal/pkg/controllers"
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"os"

	"github.com/spf13/viper"
)

func main() {
	params := server.Params{Port: os.Getenv("PORT")}
	viper.AddConfigPath("./configs/api")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err == nil {
		params.Port = viper.GetString("port")
		server.SetAuthAddress(viper.GetString("auth"))
		controllers.SetDefaultPictureAddress(viper.GetString("defaultPicture"))
		err = server.StartApp(params)
		if err != nil {
			helpers.LogMsg("Server error: ", err)
		}
	}

}

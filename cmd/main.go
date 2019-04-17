package main

import (
	"2019_1_undefined_penguins/internal/app/server"
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"os"
)

func main() {
	params := server.Params{Port: os.Getenv("PORT")}
	if params.Port == "" {
		params.Port = "8080"
	}

	err := server.StartApp(params)
	if err != nil {
		helpers.LogMsg("Server error: ", err)
	}
}

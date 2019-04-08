package main

import (
	"os"

	"2019_1_undefined_penguins/iternal/app/server"
)

func main() {

	params := server.Params{Port: os.Getenv("PORT")}
	if params.Port == "" {
		params.Port = "8080"
	}

	err := server.StartApp(params)
	if err != nil {
		panic(err)
	}
}

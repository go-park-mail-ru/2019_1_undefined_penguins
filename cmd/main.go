package main

import (
	"os"
)

func main() {
	params := server.Params{Port: os.Getenv("PORT")}
	if params.Port == "" {
		params.Port = "3000"
	}

	err := server.StartApp(params)
	if err != nil {
		panic(err)
	}
}

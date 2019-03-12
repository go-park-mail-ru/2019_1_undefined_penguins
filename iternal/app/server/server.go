package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Params struct {
	Port string
}

func StartApp(params Params) error {
	fmt.Println("Server starting at " + params.Port)

	router := mux.NewRouter()

	staticPath := "some/future/directory"
	staticHandler := http.StripPrefix(
		"/static",
		http.FileServer(http.Dir(staticPath)),
	)
	router.PathPrefix("/static").Handler(staticHandler)

	return http.ListenAndServe(":"+params.Port, router)
}

package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2019_1_undefined_penguins/iternal/pkg/controllers"
)

type Params struct {
	Port string
}

func StartApp(params Params) error {

	router := mux.NewRouter()

	router.HandleFunc("/", controllers.RootHandler)
	router.HandleFunc("/me", controllers.Me).Methods("GET")
	router.HandleFunc("/leaders", controllers.GetLeaders).Methods("GET")
	router.HandleFunc("/signup", controllers.SignUp).Methods("POST")
	router.HandleFunc("/signin", controllers.SignIn).Methods("POST")
	http.ListenAndServe(":8080", router)

	staticPath := "some/future/directory"
	staticHandler := http.StripPrefix(
		"/static",
		http.FileServer(http.Dir(staticPath)),
	)
	router.PathPrefix("/static").Handler(staticHandler)

	return http.ListenAndServe(":"+params.Port, router)
}

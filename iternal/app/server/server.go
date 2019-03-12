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
	router.HandleFunc("/me", controllers.Me).Methods("GET", "OPTIONS")
	router.HandleFunc("/leaders", controllers.GetLeaders).Methods("GET", "OPTIONS")
	router.HandleFunc("/signup", controllers.SignUp).Methods("POST", "OPTIONS")
	router.HandleFunc("/login", controllers.SignIn).Methods("POST", "OPTIONS")
	router.HandleFunc("/signin", controllers.SignOut).Methods("POST", "OPTIONS")

	return http.ListenAndServe(":"+params.Port, router)
}

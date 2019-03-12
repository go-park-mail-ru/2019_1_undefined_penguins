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

	router.HandleFunc("/", RootHandler)
	router.HandleFunc("/me", Me).Methods("GET")
	router.HandleFunc("/leaders", GetLeaders).Methods("GET")
	router.HandleFunc("/signup", SignUp).Methods("POST")
	//ВРЕМЕННО, ДАЛЕЕ НА ПУТ
	router.HandleFunc("/signin", SignIn).Methods("POST")
	http.ListenAndServe(":8080", router)

	staticPath := "some/future/directory"
	staticHandler := http.StripPrefix(
		"/static",
		http.FileServer(http.Dir(staticPath)),
	)
	router.PathPrefix("/static").Handler(staticHandler)

	return http.ListenAndServe(":"+params.Port, router)
}

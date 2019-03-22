package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2019_1_undefined_penguins/iternal/pkg/controllers"
)

type Params struct {
	Port string
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello penguins"))
}

func StartApp(params Params) error {

	router := mux.NewRouter()
	router.HandleFunc("/", RootHandler)
	router.HandleFunc("/me", controllers.Me).Methods("GET", "OPTIONS")
	router.HandleFunc("/leaders", controllers.GetLeaders).Methods("GET", "OPTIONS")
	router.HandleFunc("/signup", controllers.SignUp).Methods("POST", "OPTIONS")
	router.HandleFunc("/login", controllers.SignIn).Methods("POST", "OPTIONS")
	router.HandleFunc("/signout", controllers.SignOut).Methods("GET", "OPTIONS")
	router.HandleFunc("/change_profile", controllers.ChangeProfile).Methods("PUT", "OPTIONS")
	return http.ListenAndServe(":"+params.Port, router)
}

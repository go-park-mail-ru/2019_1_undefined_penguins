package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	c "github.com/go-park-mail-ru/2019_1_undefined_penguins/iternal/pkg/controllers"
	db "github.com/go-park-mail-ru/2019_1_undefined_penguins/iternal/pkg/database"
	mw "github.com/go-park-mail-ru/2019_1_undefined_penguins/iternal/pkg/middleware"
	"github.com/go-park-mail-ru/2019_1_undefined_penguins/iternal/pkg/helpers"
)

type Params struct {
	Port string
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello penguins"))
}

func StartApp(params Params) error {
	err := db.Connect()
	if err != nil {
		helpers.LogMsg("Connection error: ", err)
		return err
	}
	defer db.Disconnect()

	router := mux.NewRouter()

	router.Use(mw.LoggingMiddleware)
	router.Use(mw.CORSMiddleware)
	router.Use(mw.PanicMiddleware)
	router.Use(mw.AuthMiddleware)

	router.HandleFunc("/", RootHandler)
	router.HandleFunc("/me", c.Me).Methods("GET", "OPTIONS")
	router.HandleFunc("/leaders", c.GetLeaders).Methods("GET", "OPTIONS")
	router.HandleFunc("/leaderboard/{id:[0-9]+}", c.GetLeaderboardPage).Methods("GET", "OPTIONS")
	router.HandleFunc("/signup", c.SignUp).Methods("POST", "OPTIONS")
	router.HandleFunc("/login", c.SignIn).Methods("POST", "OPTIONS")
	router.HandleFunc("/signout", c.SignOut).Methods("GET", "OPTIONS")
	router.HandleFunc("/change_profile", c.ChangeProfile).Methods("PUT", "OPTIONS")
	fmt.Println("Server started")
	return http.ListenAndServe(":"+params.Port, router)
}

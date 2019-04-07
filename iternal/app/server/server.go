package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/go-park-mail-ru/2019_1_undefined_penguins/iternal/pkg/controllers"
	"github.com/go-park-mail-ru/2019_1_undefined_penguins/iternal/pkg/database"
	"github.com/go-park-mail-ru/2019_1_undefined_penguins/iternal/pkg/middleware"
)

type Params struct {
	Port string
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello penguins"))
}

func StartApp(params Params) error {
	//check this 2 funcs, is nil - return 'жёпка'
	//close connect
	database.InitConfig()
	database.Connect()
	router := mux.NewRouter()

	router.Use(middleware.LoggingMiddleware)
	router.Use(middleware.CORSMiddleware)
	router.Use(middleware.PanicMiddleware)
	router.Use(middleware.AuthMiddleware)


	router.HandleFunc("/", RootHandler)
	router.HandleFunc("/me", controllers.Me).Methods("GET", "OPTIONS")

	router.HandleFunc("/leaders", controllers.GetLeaders).Methods("GET", "OPTIONS")
	router.HandleFunc("/leaderboard/{id:[0-9]+}", controllers.GetLeaderboardPage).Methods("GET", "OPTIONS")
	router.HandleFunc("/signup", controllers.SignUp).Methods("POST", "OPTIONS")
	router.HandleFunc("/login", controllers.SignIn).Methods("POST", "OPTIONS")
	router.HandleFunc("/signout", controllers.SignOut).Methods("GET", "OPTIONS")
	router.HandleFunc("/change_profile", controllers.ChangeProfile).Methods("PUT", "OPTIONS")
	fmt.Println("Server started")
	return http.ListenAndServe(":"+params.Port, router)
}

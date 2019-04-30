package server

import (
	"2019_1_undefined_penguins/internal/pkg/fileserver"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	c "2019_1_undefined_penguins/internal/pkg/controllers"
	db "2019_1_undefined_penguins/internal/pkg/database"

	"2019_1_undefined_penguins/internal/pkg/helpers"
	mw "2019_1_undefined_penguins/internal/pkg/middleware"

	"github.com/gorilla/handlers"
)

type Params struct {
	Port string
}

func StartApp(params Params) error {
	err := db.Connect()
	if err != nil {
		helpers.LogMsg("Connection error: ", err)
		return err
	}
	defer db.Disconnect()
	router := mux.NewRouter()

	router.Use(mw.PanicMiddleware)
	router.Use(mw.CORSMiddleware)
	router.Use(mw.AuthMiddleware)

	router.HandleFunc("/", c.RootHandler)
	router.HandleFunc("/me", c.Me).Methods("GET", "OPTIONS")
	router.HandleFunc("/leaders/{id:[0-9]+}", c.GetLeaderboardPage).Methods("GET", "OPTIONS")
	router.HandleFunc("/leaders/info", c.GetLeaderboardInfo).Methods("GET", "OPTIONS")
	router.HandleFunc("/signup", c.SignUp).Methods("POST", "OPTIONS")
	router.HandleFunc("/login", c.SignIn).Methods("POST", "OPTIONS")
	router.HandleFunc("/signout", c.SignOut).Methods("GET", "OPTIONS")
	router.HandleFunc("/change_profile", c.ChangeProfile).Methods("PUT", "OPTIONS")
	router.HandleFunc("/upload", c.UploadImage).Methods("POST")
	router.HandleFunc("/ws", c.StartWS)

	helpers.LogMsg("Server started at " + params.Port)
	go func() {
		fileserver.Start()
	}()

	return http.ListenAndServe(":"+params.Port, handlers.LoggingHandler(os.Stdout, router))
}

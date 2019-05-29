package server

import (
	"2019_1_undefined_penguins/internal/app/metrics"
	"2019_1_undefined_penguins/internal/pkg/models"
	"google.golang.org/grpc"

	"2019_1_undefined_penguins/internal/pkg/fileserver"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"

	c "2019_1_undefined_penguins/internal/pkg/controllers"
	"2019_1_undefined_penguins/internal/pkg/helpers"
	mw "2019_1_undefined_penguins/internal/pkg/middleware"
	"github.com/gorilla/mux"

	"github.com/gorilla/handlers"
)

type Params struct {
	Port string
}


var authAddress string

func SetAuthAddress(address string) {
	authAddress = address
}

func StartApp(params Params) error {

	router := mux.NewRouter()
	authRouter := router.PathPrefix("/api").Subrouter()
	//gameRouter := router.PathPrefix("/data").Subrouter()

	grcpConn, err := grpc.Dial(
		//authAddress,
		"127.0.0.1:8083",
		grpc.WithInsecure(),
	)
	if err != nil {
		helpers.LogMsg("Can`t connect to grpc")
		return err
	}
	defer grcpConn.Close()

	models.AuthManager = models.NewAuthCheckerClient(grcpConn)

	router.Use(mw.PanicMiddleware)
	//router.Use(mw.MonitoringMiddleware)
	router.Use(mw.CORSMiddleware)
	router.Use(mw.AuthMiddleware)

	router.HandleFunc("/", c.RootHandler)
	authRouter.HandleFunc("/me", c.Me).Methods("GET", "OPTIONS")
	authRouter.HandleFunc("/leaders/{id:[0-9]+}", c.GetLeaderboardPage).Methods("GET", "OPTIONS")
	authRouter.HandleFunc("/leaders/info", c.GetLeaderboardInfo).Methods("GET", "OPTIONS")
	authRouter.HandleFunc("/signup", c.SignUp).Methods("POST", "OPTIONS")
	authRouter.HandleFunc("/login", c.SignIn).Methods("POST", "OPTIONS")
	authRouter.HandleFunc("/signout", c.SignOut).Methods("GET", "OPTIONS")
	authRouter.HandleFunc("/me", c.ChangeProfile).Methods("PUT")
	authRouter.HandleFunc("/upload", c.UploadImage).Methods("POST")
	authRouter.HandleFunc("/checkSingleWs", c.CheckWsSingle).Methods("GET", "OPTIONS")
	authRouter.HandleFunc("/checkMultiWs", c.CheckWsMulti).Methods("GET", "OPTIONS")


	//gameRouter.HandleFunc("/ws", c.StartWS)

	prometheus.MustRegister(metrics.Hits)
	router.Handle("/metrics", promhttp.Handler())



	helpers.LogMsg("Server started at " + params.Port)
	go func() {
		fileserver.Start()
	}()

	return http.ListenAndServe(":"+params.Port, handlers.LoggingHandler(os.Stdout, router))
}

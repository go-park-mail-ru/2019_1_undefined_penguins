package main

import (
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"2019_1_undefined_penguins/internal/app/metrics"
	mw "2019_1_undefined_penguins/internal/pkg/middleware"
	"2019_1_undefined_penguins/internal/pkg/models"
	"2019_1_undefined_penguins/internal/pkg/game"
	"google.golang.org/grpc"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func setConfig() (string, int, string) {
	viper.AddConfigPath("./configs/game")
	viper.SetConfigName("game")
	var port, authAddress string
	var maxRooms int
	if err := viper.ReadInConfig(); err != nil {
		port = ":8085"
		maxRooms = 20
		authAddress = "127.0.0.1:8083"
	} else {
		port = ":" + viper.GetString("port")
		maxRooms = viper.GetInt("maxRooms")
		authAddress = viper.GetString("auth")
	}
	return port, maxRooms, authAddress
}

func main() {
	port, maxRooms, authAddress := setConfig()
	game.PingGame = game.InitGame(uint(maxRooms))
	go game.PingGame.Run()

	grcpConn, err := grpc.Dial(
		authAddress,
		grpc.WithInsecure(),
	)
	if err != nil {
		helpers.LogMsg("Can`t connect to grpc")
		return
	}
	defer grcpConn.Close()

	models.AuthManager = models.NewAuthCheckerClient(grcpConn)

	router := mux.NewRouter()
	gameRouter := router.PathPrefix("/game").Subrouter()

	//TODO
	//router.Use(mw.PanicMiddleware)
	gameRouter.Use(mw.CORSMiddleware)

	prometheus.MustRegister(metrics.PlayersCountInGame, metrics.ActiveRooms)
	router.Handle("/metrics", promhttp.Handler())
	gameRouter.HandleFunc("/single", game.StartSingle)
	gameRouter.HandleFunc("/multi", game.StartMulti)

	helpers.LogMsg("GameServer started at", port)

	http.ListenAndServe(port, handlers.LoggingHandler(os.Stdout, router))
}
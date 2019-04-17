package controllers

import (
	"2019_1_undefined_penguins/internal/app/game"
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"net/http"

	"github.com/gorilla/websocket"
)

func StartWS(w http.ResponseWriter, r *http.Request) {
	//pingGame := game.NewGame(10)
	//go pingGame.Run()

	upgrader := &websocket.Upgrader{}

	cookie, err := r.Cookie("sessionid")
	if err != nil {
		helpers.LogMsg("Not authorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, http.Header{"Upgrade": []string{"websocket"}})
	if err != nil {
		helpers.LogMsg("Connection error: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	helpers.LogMsg("Connected to client")

	player := game.NewPlayer(conn, cookie.Value)
	go player.Listen()
	game.PingGame.AddPlayer(player)
}

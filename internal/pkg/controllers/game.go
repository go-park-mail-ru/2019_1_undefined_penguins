package controllers

import (
	"2019_1_undefined_penguins/internal/app/game"
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)


func StartWS(w http.ResponseWriter, r *http.Request) {
	pingGame := game.NewGame(10)
	go pingGame.Run()

	//http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
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
		pingGame.AddPlayer(player)
	//})

	fmt.Println("Started game")

	//http.ListenAndServe(":8082", nil)
	//if err != nil {
	//	helpers.LogMsg("Cannot start server: ", err)
	//}
}
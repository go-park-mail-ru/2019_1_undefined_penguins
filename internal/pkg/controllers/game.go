package controllers

import (
	"net/http"
)

func StartWS(w http.ResponseWriter, r *http.Request) {

	// if game.PingGame.RoomsCount() >= 10 {
	// 	//TODO check response on the client side
	// 	helpers.LogMsg("Too many clients")
	// 	w.WriteHeader(http.StatusTooManyRequests)
	// 	w.Write([]byte("Too many clients"))
	// 	return
	// }

	// upgrader := &websocket.Upgrader{}

	// //check for multi in micro!!!!!

	// //cookie, err := r.Cookie("sessionid")
	// //if err != nil {
	// //	cookie.Value = "Anonumys"
	// //}

	// upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	// conn, err := upgrader.Upgrade(w, r, nil)
	// if err != nil {
	// 	helpers.LogMsg("Connection error: ", err)
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	// helpers.LogMsg("Connected to client")

	// //TODO remove hardcore, get from front player value
	// player := game.NewPlayer(conn)
	// go player.Listen()
}

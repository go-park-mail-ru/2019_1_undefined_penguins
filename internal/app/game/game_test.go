package game

import (
	"testing"

	"github.com/gorilla/websocket"
)

func TestGameStart(t *testing.T) {

	go func() {
		err := Start()
		if err == nil {
			t.Error(err)

		}
	}()
}

func TestNewRoom(t *testing.T) {
	game := NewGame(10)
	if game == nil {
		t.Error("Игра не создана")
	}
	room := NewRoom(2)
	if room == nil {
		t.Error("Комната не создана")
	}
	game.AddRoom(room)
}

func TestPlayer(t *testing.T) {
	var conn *websocket.Conn
	player := NewPlayer(conn, "55")
	go player.Listen()
	var playerState PlayerState
	RotatePlayer(&playerState)

	room := NewRoom(2)
	if room == nil {
		t.Error("Комната не создана")
	}
	var bulletState BulletState
	// bulletState := CreateBullet(room)
	bulletState.X = 1
	bulletState.Y = 1
	ShotPlayer(&playerState, &bulletState)
}

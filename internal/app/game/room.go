package game

import (
	"time"
	"sync"
	"2019_1_undefined_penguins/internal/pkg/helpers"
)

type PlayerState struct {
	ID string
	X, Y int
}

type ObjectState struct {
	ID string
	Type string
	X, Y int
}

type RoomState struct {
	Players []PlayerState
	Objects []ObjectState
	CurrentTime time.Time
}

type Room struct {
	ID string
	MaxPlayers uint
	Players map[string]*Player
	mu *sync.Mutex
	register chan *Player
	unregister chan *Player
	ticker *time.Ticker
	state *RoomState
}

func NewRoom(MaxPlayers uint) *Room {
	return &Room{
		MaxPlayers: MaxPlayers,
		Players: make(map[string]*Player),
		register: make(chan *Player),
		unregister: make(chan *Player),
		ticker: time.NewTicker(1*time.Second),
	}
}

func (r *Room) Run() {
	helpers.LogMsg("Room loop started")
	for {
		select {
		case player := <-r.unregister:
			delete(r.Players, player.ID)
			helpers.LogMsg("Player "+player.ID+" was removed from room")
		case player := <-r.register:
			r.Players[player.ID] = player
			helpers.LogMsg("Player "+player.ID+" joined")
			player.SendMessage(&Message{"CONNECTED", nil})
		case <-r.ticker.C:
			helpers.LogMsg("Tick")

			//my code

			for _, player := range r.Players {
				player.SendState(r.state)
			}
		}
	}
}

func (r *Room) AddPlayer(player *Player) {
	player.room = r
	r.register <- player
}

func (r *Room) RemovePlayer(player *Player) {
	r.unregister <- player
}


package game

import (
	"fmt"
	"time"
	"sync"
	"2019_1_undefined_penguins/internal/pkg/helpers"
)

type PlayerState struct {
	ID string
	ClockwiseDirection bool
	Shoted bool
	X, Y int
	Alpha float64
	Score int
}

type BulletState struct {
	ID string
	X, Y int
	Alpha float64
}

type RoomState struct {
	Players map[string]*PlayerState
	Objects BulletState
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
		state: &RoomState{
			Players: make(map[string]*PlayerState),
		},
	}
}

func (r *Room) Run() {
	helpers.LogMsg("Room loop started")
	r.state.Objects = CreateBullet(r)
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
			for _, player := range r.Players {
				msg := <-player.in
				switch msg.Payload {
					case "SHOT":
						ShotPlayer(r.state.Players[msg.Type], &r.state.Objects)
					case "ROTATE":
						RotatePlayer(r.state.Players[msg.Type])
				}
				fmt.Println("sfw:", msg)
				player.SendState(r.state)
			}
		}
	}
}

func (r *Room) AddPlayer(player *Player) {
	ps := &PlayerState{
		ID: player.ID,
		X: 0,
		Y: 0,
		Alpha: 0,
		ClockwiseDirection: true,
		Shoted: false,
	}
	r.state.Players[player.ID] = ps
	player.room = r
	r.register <- player
}

func (r *Room) RemovePlayer(player *Player) {
	r.unregister <- player
}


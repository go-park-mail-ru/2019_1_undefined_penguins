package game

import (
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"sync"
	"time"
)

type PlayerState struct {
	ID                 string
	ClockwiseDirection bool
	Shoted             bool
	X, Y               int
	Alpha              float64
	Score              int
	//penguin("GOOD") or gun("BAD")
	Type string
}

type BulletState struct {
	ID    string
	X, Y  int
	Alpha float64
	Radious int
}

type FishState struct {
	ID int
	X, Y int
	Alpha float64
	Eaten bool
}

type RoomState struct {
	Players map[string]*PlayerState
	Bullet  *BulletState
	Fishes 	map[int]*FishState
	Radious int
	CurrentTime time.Time
}

type Room struct {
	ID         string
	MaxPlayers uint
	Players    map[string]*Player
	mu         *sync.Mutex
	register   chan *Player
	unregister chan *Player
	ticker     *time.Ticker
	state      *RoomState

	broadcast chan *Message
	finish chan *Message
}

func NewRoom(MaxPlayers uint) *Room {
	return &Room{
		MaxPlayers: MaxPlayers,
		Players:    make(map[string]*Player),
		register:   make(chan *Player),
		unregister: make(chan *Player),
		ticker:     time.NewTicker(1000 * time.Millisecond),
		state: &RoomState{
			Players: make(map[string]*PlayerState),
			Fishes: make(map[int]*FishState, 24),
			Radious: 250,
		},
		broadcast: make(chan *Message),
		finish: make(chan *Message),
	}
}

func (r *Room) Run() {
	helpers.LogMsg("Room loop started")
	r.state.Bullet = CreateBullet(r)
	GameInit(r)
	for {
		select {
		case player := <-r.unregister:
			delete(r.Players, player.ID)
			helpers.LogMsg("Player " + player.ID + " was removed from room")
		case player := <-r.register:
			r.Players[player.ID] = player
			helpers.LogMsg("Player " + player.ID + " joined")
			player.SendMessage(&Message{"CONNECTED", nil})
		case message := <- r.broadcast:
			for _, player := range r.Players {
				select {
				case player.out <- message:
				default:
					close(player.out)
				}
			}
			HandleCommand(r, message)
		case <-r.ticker.C:
			ProcessGameSingle(r)
		case <- r.finish:
			FinishGame(r)
			return
		}
	}
}

func (r *Room) AddPlayer(player *Player) {
	ps := &PlayerState{
		ID:                 player.ID,
		X:                  0,
		Y:                  0,
		Alpha:              0,
		ClockwiseDirection: true,
		Shoted:             false,
		Score:				0,
		//TODO it is for single (ПНС)
		Type: 				"GOOD",
	}
	r.state.Players[player.ID] = ps
	player.room = r
	r.register <- player
}

func (r *Room) RemovePlayer(player *Player) {
	r.unregister <- player
}

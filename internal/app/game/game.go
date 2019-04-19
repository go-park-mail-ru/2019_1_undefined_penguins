package game

import (
	"sync"
	"2019_1_undefined_penguins/internal/pkg/helpers"
)

var PingGame *Game

func InitGame() *Game {
	g := &Game{
		MaxRooms: 10,
		register: make(chan *Player),
	}
	return g
}

type Game struct {
	MaxRooms uint
	rooms []*Room
	mu *sync.Mutex
	register chan *Player
}

func NewGame(maxRooms uint) *Game {
	return &Game{
		MaxRooms: maxRooms,
		register: make(chan *Player),
	}
}

func (g *Game) Run()  {
	//helpers.LogMsg("Main loop started")

//TODO remove goto metka by reversing condition
LOOP:
	for {
		player := <-g.register

		for _, room := range g.rooms {
			if len(room.Players) < int(room.MaxPlayers) {
				room.AddPlayer(player)
				continue LOOP
			}
		}

		//если все комнаты заняты - делой новую
		room := NewRoom(2)
		g.AddRoom(room)
		go room.Run()

		room.AddPlayer(player)
	}
}

func (g *Game) AddPlayer(player *Player)  {
	helpers.LogMsg("Player " + player.ID + " queued to add")
	g.register <- player
}

func (g *Game) AddRoom(room *Room)  {
	g.rooms = append(g.rooms, room)
}

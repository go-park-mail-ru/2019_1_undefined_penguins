package game

import (
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"fmt"
	"sync"
)

var PingGame *Game

const (
	SINGLE = "SINGLE"
	MULTI = "MULTI"

	WAIT = "SIGNAL_TO_WAIT_OPPONENT"
	START = "SIGNAL_START_THE_GAME"
	FINISH = "SIGNAL_FINISH_GAME"
	STATE = "SIGNAL_NEW_GAME_STATE"

	NEWPLAYER = "newPlayer"
	NEWCOMMAND = "newCommand"
)

func InitGame(maxRooms uint) *Game {
	return NewGame(maxRooms)
}

type Game struct {
	MaxRooms uint
	roomsSingle []*RoomSingle
	roomsMulti []*RoomMulti
	//mu *sync.Mutex
	mu sync.RWMutex
	register chan *Player
}

func NewGame(maxRooms uint) *Game {
	return &Game{
		MaxRooms: maxRooms,
		register: make(chan *Player),
	}
}

func (g *Game) Run()  {
LOOP:
	for {
		player, _ := <-g.register
		//fmt.Println("register ch is ", ok)
		//fmt.Println("State is "+ player.GameMode)

		switch player.GameMode {
		case SINGLE:
			//start roomSingle
			for _, room := range g.roomsSingle {
				if room.Player == nil {
					g.mu.Lock()
					room.AddPlayer(player)
					g.mu.Unlock()
					continue LOOP
				}
			}

			//если все комнаты заняты - делой новую
			room := NewRoomSingle(1)
			g.mu.Lock()
			g.AddToRoomSingle(room)
			g.mu.Unlock()

			go room.Run()

			g.mu.Lock()
			room.AddPlayer(player)
			g.mu.Unlock()

		case MULTI:
			//start roomMulty
			for _, room := range g.roomsMulti {
				if len(room.Players) < int(room.MaxPlayers) {
					g.mu.Lock()
					room.AddPlayer(player)
					g.mu.Unlock()
					continue LOOP
				}
			}

			//если все комнаты заняты - делой новую
			room := NewRoomMulti(2)
			g.mu.Lock()
			g.AddToRoomMulti(room)
			g.mu.Unlock()

			go room.Run()

			g.mu.Lock()
			room.AddPlayer(player)
			g.mu.Unlock()
		default:
			fmt.Println("Empty")
		}
	}
}

func (g *Game) AddToRoomSingle(room *RoomSingle) {
	g.roomsSingle = append(g.roomsSingle, room)
}

func (g *Game) AddToRoomMulti(room *RoomMulti) {
	g.mu.Lock()
	g.roomsMulti = append(g.roomsMulti, room)
	g.mu.Unlock()
}

func (g *Game) AddPlayer(player *Player)  {
	helpers.LogMsg("Player " + player.ID + " queued to add")
	g.register <- player
}


func (g *Game) RoomsCount() int {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return len(g.roomsSingle) + len(g.roomsMulti)
}

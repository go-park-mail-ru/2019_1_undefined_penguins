package main

import (
	//"game/helpers"
	"fmt"
	"game/helpers"
	"game/metrics"
	"sync"
)

var PingGame *Game
var maxRooms uint
var SingleRoomsCount int
var MultiRoomsCount int

func InitGame(rooms uint) *Game {
	maxRooms = rooms
	game :=  NewGame(maxRooms)
	return game
}

type Game struct {
	MaxRooms    uint
	roomsSingle map[int]*RoomSingle
	roomsMulti  map[int]*RoomMulti
	mu       sync.RWMutex
	register chan *Player
	unregister chan *Player
	Players    map[string]*Player
}

func NewGame(maxRooms uint) *Game {
	return &Game{
		MaxRooms: maxRooms,
		register: make(chan *Player),
		unregister: make(chan *Player),
		roomsSingle: make(map[int]*RoomSingle),
		roomsMulti:  make(map[int]*RoomMulti),
		Players:    make(map[string]*Player),
	}
}

func (g *Game) Run() {
	//defer helpers.RecoverPanic()
//LOOP:
	for {
		select {

		case player, _ := <-g.register:
			switch player.GameMode {
				case SINGLE:
					g.ProcessSingle(player)
				case MULTI:
					g.ProcessMulti(player)
				default:
					fmt.Println("Empty")
				}
		case player, _ := <-g.unregister:
			delete(g.Players, player.ID)
			for _, room := range g.roomsMulti {
				fmt.Println(len(room.Players))
				if room != nil && len(room.Players) == 0 {
					delete(g.roomsMulti, room.ID)
				}
			}
			for _, room := range g.roomsSingle {
				if room != nil && room.Player == nil {
					delete(g.roomsSingle, room.ID)
				}
			}
			helpers.LogMsg("Player " + player.ID + " was removed from PingGame")
		}

	}
}

func (g *Game) AddToRoomSingle(room *RoomSingle) {
	metrics.ActiveRooms.Inc()
	g.roomsSingle[SingleRoomsCount] = room
	SingleRoomsCount++
	//g.roomsSingle = append(g.roomsSingle, room)
}

func (g *Game) AddToRoomMulti(room *RoomMulti) {
	metrics.ActiveRooms.Inc()
	g.roomsMulti[MultiRoomsCount] = room
	MultiRoomsCount++
	//g.roomsMulti = append(g.roomsMulti, room)
}

func (g *Game) AddPlayer(player *Player) {
	helpers.LogMsg("Player " + player.ID + " queued to add")
	g.mu.Lock()
	g.Players[player.ID] = player
	g.mu.Unlock()
	metrics.PlayersCountInGame.Inc()
	g.register <- player
}

func (g *Game) RoomsCount() int {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return len(g.roomsSingle) + len(g.roomsMulti)
}

func (g *Game) ProcessSingle(player *Player) {
	//TODO remove loop
	for _, room := range g.roomsSingle {
		if room.Player == nil {
			g.mu.Lock()
			room.AddPlayer(player)
			//TODO add game states
			g.mu.Unlock()
			room.state = CreateInitialStateSingle(room)
			room.state.Penguin.ID = player.ID
			room.gameState = RUNNING
			player.out <- &OutcomeMessage{
				Type: START,
				Payload:OutPayloadMessage{
					Gun:GunMessage{
						Bullet:BulletMessage{
							Alpha: 0,
							DistanceFromCenter: 0,
						},
						Alpha: 0,
						Name: string(GUN),
						Result: "",
					},
					Penguin:PenguinMessage{
						Clockwise:false,
						Alpha: 0,
						Name: player.ID,
						Result: "",
					},
				},
			}
			//continue LOOP
			return
		}
	}

	//если все комнаты заняты - делой новую
	room := NewRoomSingle(1, SingleRoomsCount)
	g.mu.Lock()
	g.AddToRoomSingle(room)
	g.mu.Unlock()

	go room.Run()

	g.mu.Lock()
	room.AddPlayer(player)
	g.mu.Unlock()

	room.state = CreateInitialStateSingle(room)
	room.state.Penguin.ID = player.ID
	room.gameState = WAITING

	room.SendRoomState(&OutcomeMessage{
		Type: FINISHROUND,
		Payload:OutPayloadMessage{
			Gun:GunMessage{
				Bullet:BulletMessage{
					Alpha: 0,
					DistanceFromCenter: 0,
				},
				Alpha: 0,
				Name: string(GUN),
				Result: "",
			},
			Penguin:PenguinMessage{
				Clockwise:false,
				Alpha: 0,
				Name: player.ID,
				Result: "",
			},
			Round: 1,
			PiscesCount: 24,
		},
	})
}

func (g *Game) ProcessMulti(player *Player) {
	for _, room := range g.roomsMulti {
		if room.gameState == PICKINGUP { //len(room.Players) < int(room.MaxPlayers) {
			g.mu.Lock()
			room.AddPlayer(player)
			g.mu.Unlock()

			room.state = CreateInitialState(room)

			penguin, gun := room.SelectPlayersRoles()
			room.state.Penguin.ID = penguin
			room.state.Gun.ID = gun
			room.gameState = WAITING
			room.state.Round = 1
			room.round = 1
			room.SendRoomState(&OutcomeMessage{
				Type:INITIALIZEDGAME,
				Payload:OutPayloadMessage{
					Penguin:PenguinMessage{
						Name: room.state.Penguin.ID,
						Score: uint(room.state.Penguin.Score),
					},
					Gun:GunMessage{
						Name: room.state.Gun.ID,
						Score: uint(room.state.Gun.Score),
					},
					Round: uint(room.state.Round),
					PiscesCount: 24,
				}})
			//room.StartNewRound()
			//continue LOOP
			return
		}
	}
	//если все комнаты заняты - делой новую
	room := NewRoomMulti(2, SingleRoomsCount)
	g.mu.Lock()
	g.AddToRoomMulti(room)
	g.mu.Unlock()

	go room.Run()

	g.mu.Lock()
	room.AddPlayer(player)
	player.out <- &OutcomeMessage{Type: WAIT}
	g.mu.Unlock()
	room.gameState = PICKINGUP
	room.state.Round = 1
	room.round = 1
}

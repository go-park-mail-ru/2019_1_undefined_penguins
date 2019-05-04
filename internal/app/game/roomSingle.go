package game

import (
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"sync"
	"time"
)

type RoomSingle struct {
	ID         string
	MaxPlayers uint
	Player     *Player
	mu         sync.Mutex
	register   chan *Player
	unregister chan *Player
	ticker     *time.Ticker
	state      *RoomState

	broadcast chan *OutcomeMessage
	finish chan *Player
}

func NewRoomSingle(MaxPlayers uint) *RoomSingle {
	return &RoomSingle{
		MaxPlayers: MaxPlayers,
		Player:    new(Player),
		register:   make(chan *Player),
		unregister: make(chan *Player),
		ticker:     time.NewTicker(100 * time.Millisecond),
		state: &RoomState{
			Penguin: new(PenguinState),
			Fishes: make(map[int]*FishState, 24),
		},
		broadcast: make(chan *OutcomeMessage),
		finish: make(chan *Player),
	}
}

func (r *RoomSingle) Run() {
	helpers.LogMsg("Room Single loop started")
	//r.state.Gun.Bullet = CreateBullet(r)
	//GameInit(r)
	for {
		select {
		case player := <-r.unregister:
			r.Player = nil
			helpers.LogMsg("Player " + player.ID + " was removed from room")
		case player := <-r.register:
			r.mu.Lock()
			r.Player = player
			r.mu.Unlock()
			helpers.LogMsg("Player " + player.ID + " joined")
			player.out <- &OutcomeMessage{Type:START}
		case message := <- r.broadcast:
				select {
				case r.Player.out <- message:
				default:
					close(r.Player.out)
				}
			//HandleCommand(r, message)
		case <-r.ticker.C:
			//ProcessGameSingle(r)
		case player := <- r.finish:
			helpers.LogMsg("Player " + player.ID + " finished game")
			player.out <- &OutcomeMessage{Type:FINISH}
			r.state.Penguin = nil
			//FinishGame(r)
		}
	}
}

func (r *RoomSingle) AddPlayer(player *Player) {
	ps := &PenguinState{
		ID:                 player.ID,
		Alpha:              0,
		ClockwiseDirection: true,
		Score:				0,
	}
	r.mu.Lock()
	r.state.Penguin = ps
	r.mu.Unlock()
	player.roomSingle = r
	r.register <- player
}

//func (r *RoomSingle) RemovePlayer(player *Player) {
//	r.unregister <- player
//}

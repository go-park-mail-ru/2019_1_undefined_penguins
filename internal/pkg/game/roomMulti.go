package main

import (
	"fmt"
	"game/helpers"
	"game/models"
	"golang.org/x/net/context"
	"math/rand"
	"sync"
	"time"
)

type RoomMulti struct {
	ID         int
	MaxPlayers uint
	Players    map[string]*Player
	mu         sync.Mutex
	register   chan *Player
	unregister chan *Player
	ticker     *time.Ticker
	state      *RoomState
	gameState GameCurrentState
	round int

	broadcast chan *OutcomeMessage
	finish chan *Player
}

func NewRoomMulti(MaxPlayers uint, id int) *RoomMulti {
	return &RoomMulti{
		ID: id,
		MaxPlayers: MaxPlayers,
		Players:    make(map[string]*Player),
		register:   make(chan *Player),
		unregister: make(chan *Player),
		ticker:     time.NewTicker(50 * time.Millisecond),
 		state: &RoomState{
			Penguin: new(PenguinState),
			Gun: new(GunState),
			Fishes: make(map[int]*FishState, 24),
			Round: 1,
		},
		round: 1,
		broadcast: make(chan *OutcomeMessage),
		finish: make(chan *Player),
	}
}

func (r *RoomMulti) Run() {
	//defer helpers.RecoverPanic()
	helpers.LogMsg("Room Multi loop started")
	for {
		select {
		case player := <-r.unregister:
			delete(r.Players, player.ID)
			helpers.LogMsg("Player " + player.ID + " was removed from room")
		case player := <-r.register:
			r.mu.Lock()
			r.Players[player.ID] = player
			r.mu.Unlock()
			helpers.LogMsg("Player " + player.ID + " joined")
		//case message := <- r.broadcast:
			//r.SendRoomState(message)
		case <-r.ticker.C:
			if r.gameState == RUNNING {
				  message := RunMulti(r)
				  if message.Type != STATE {
					  switch message.Type {
					  case FINISHROUND:
					  		fmt.Println(FINISHROUND)
					  		fmt.Println(r.gameState)
					  case FINISHGAME:
							//message = r.FinishGame()
					  }
				  }
				if r.round > LastRound && r.gameState == FINISHED {
					r.SaveResult()
					message := r.FinishGame()
					r.SendRoomState(message)
					r.state.Round = 1
					r.round = 1
					return
				} else {
					r.SendRoomState(message)
				}
			}
		}
	}
}

func (r *RoomMulti) AddPlayer(player *Player) {
	ps := &PenguinState{
		ID:                 player.ID,
		Alpha:              0,
		ClockwiseDirection: true,
		Score:				0,
	}
	r.mu.Lock()
	r.state.Penguin = ps
	r.mu.Unlock()
	player.roomMulti = r
	r.register <- player
}

func (r *RoomMulti) RemovePlayer(player *Player) {
	//r.unregister <- player
	r.mu.Lock()
	delete(r.Players, player.ID)
	r.mu.Unlock()
	helpers.LogMsg("Player " + player.ID + " was removed from room")
}

func (r *RoomMulti) SelectPlayersRoles() (string, string) {
	var penguin, gun string
	digit := rand.Intn(2)
	time.Sleep(10* time.Millisecond)
	for _, player := range r.Players {
		if digit == 0 {
			player.Type = PENGUIN
			penguin = player.ID
			digit = 1
		} else {
			player.Type = GUN
			gun = player.ID
			digit = 0
		}
	}
	return penguin, gun
}

func (r *RoomMulti) ProcessCommand(message *IncomeMessage) {
	login := message.Payload.Name
	for _, player := range r.Players {
		if player.ID != login {
			continue
		}
		fmt.Println(r.state)
		switch player.Type {
		case PENGUIN:
			r.state.RotatePenguin()
		case GUN:
			r.state.RotateGun()
		default:
			fmt.Println("Incorrect player type!")
		}
		break
	}
}

func (r *RoomMulti) FinishGame() *OutcomeMessage {
	message := new(OutcomeMessage)
	for _, player := range r.Players {
		helpers.LogMsg("Player " + player.ID + " finished game")
	}
	//r.gameState = FINISHED
	if r.state.Penguin.Score > r.state.Gun.Score {
		message = &OutcomeMessage{
			Type: FINISHGAME,
			Payload: OutPayloadMessage{
				Penguin: PenguinMessage{
					Name:   r.state.Penguin.ID,
					Score:  uint(r.state.Penguin.Score),
					Result: WIN,
				},
				Gun: GunMessage{
					Name:   r.state.Gun.ID,
					Score:  uint(r.state.Gun.Score),
					Result: LOST,
				},
			},
		}

		r.gameState = FINISHED
		r.round = 1
		r.state.Round = 1
	} else {
		message = &OutcomeMessage{
			Type: FINISHGAME,
			Payload: OutPayloadMessage{
				Penguin: PenguinMessage{
					Name:   r.state.Penguin.ID,
					Score:  uint(r.state.Penguin.Score),
					Result: LOST,
				},
				Gun: GunMessage{
					Name:   r.state.Gun.ID,
					Score:  uint(r.state.Gun.Score),
					Result: WIN,
				},
			},
		}
		r.gameState = FINISHED
		r.round = 1
		r.state.Round = 1
	}
	return  message
}

func (r *RoomMulti) FinishRound() {
	for _, player := range r.Players {
		helpers.LogMsg("Player " + player.ID + " finished round")
	}
	r.gameState = WAITING
	if r.round > LastRound {
		r.gameState = FINISHED
	}
}

func (r *RoomMulti) SendRoomState(message *OutcomeMessage) {
	for _, player := range r.Players {
		if message == nil {
			continue
		}
		player.out <- message
		//select {
		//case player.out <- message:
		//default:
		//	close(player.out)
		//}
	}
}

func (r *RoomMulti) StartNewRound() {
	//time.Sleep(1000 * time.Millisecond)
	if r.state != nil && r.round <= LastRound {
		r.round += 1
		r.state.Round = r.round
		//penguin, gun := r.SelectPlayersRoles()
		message := &OutcomeMessage{
			Type: START,
			Payload: OutPayloadMessage{
				Gun: GunMessage{
					//Name: gun,
					Name: r.state.Gun.ID,
					Score: uint(r.state.Gun.Score),
				},
				Penguin: PenguinMessage{
					//Name: penguin,
					Name: r.state.Penguin.ID,
					Score: uint(r.state.Penguin.Score),
				},
				PiscesCount: 24,
				Round:       uint(r.round),
			},
		}
		r.SendRoomState(message)
		r.state = CreateInitialState(r)
		r.gameState = RUNNING
	} else {
		if r.round > LastRound {
			message := r.FinishGame()
			r.round = 1
			r.SendRoomState(message)
			//r.gameState = FINISHED
		}
	}
}

func (r *RoomMulti) SaveResult() {
	players := r.Players
	for _, player := range players {
		if player.Type == PENGUIN {
			player.instance.Score = uint64(player.roomMulti.state.Penguin.Score)
		}
		if player.Type == GUN {
			player.instance.Score = uint64(player.roomMulti.state.Gun.Score)
		}
		ctx := context.Background()
		_, err := models.AuthManager.SaveUserGame(ctx, player.instance)
		_, err = models.AuthManager.DeleteUserFromGame(ctx, player.instance)
		fmt.Println(err)
	}
}

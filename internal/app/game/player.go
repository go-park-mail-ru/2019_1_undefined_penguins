package game

import (
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"fmt"
	"github.com/gorilla/websocket"
)

type Player struct {
	conn *websocket.Conn
	ID   string
	in   chan *IncomeMessage
	out  chan *Message
	room *Room
}

func NewPlayer(conn *websocket.Conn, id string) *Player {
	return &Player{
		conn: conn,
		ID:   id,
		in:   make(chan *IncomeMessage),
		out:  make(chan *Message),
	}
}

func (p *Player) Listen() {
	go func() {
		for {
			message := &IncomeMessage{}
			err := p.conn.ReadJSON(message)
			if websocket.IsUnexpectedCloseError(err) {
				p.room.RemovePlayer(p)
				helpers.LogMsg("Player " + p.ID + " disconnected")
				return
			}
			if err != nil {
				helpers.LogMsg("Cannot read json: ", err)
				continue
			}

			p.in <- message
		}
	}()

	for {
		select {
		case message := <-p.out:
			p.conn.WriteJSON(message)
		case message := <-p.in:
			fmt.Printf("income: %#v", message)
		}
	}
}

func (p *Player) SendState(state *RoomState) {
	//TODO: send to front
	if state != nil {
		p.out <- &Message{"STATE", state}
	}
}

func (p *Player) SendMessage(message *Message) {
	p.out <- message
}

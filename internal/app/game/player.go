package game

import (
	"2019_1_undefined_penguins/internal/pkg/helpers"
	"fmt"
	"github.com/gorilla/websocket"
)

type Player struct {
	conn *websocket.Conn
	ID   string
	//in   chan *IncomeMessage
	out  chan *Message
	room *Room
}

func NewPlayer(conn *websocket.Conn, id string) *Player {
	return &Player{
		conn: conn,
		ID:   id,
		//in:   make(chan *IncomeMessage),
		out:  make(chan *Message),
	}
}

func (p *Player) Listen() {
	go func() {
		for {
			message := &Message{}
			err := p.conn.ReadJSON(message)
			if websocket.IsUnexpectedCloseError(err) {
				p.room.RemovePlayer(p)
				helpers.LogMsg("Player " + p.ID + " disconnected")
				return
			}
			//message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

			if err != nil {
				helpers.LogMsg("Cannot read json: ", err)
				continue
			}
			p.room.broadcast <- message
			//p.in <- message
		}
	}()

	for {
		select {
		case message := <-p.out:
			p.conn.WriteJSON(message)
		//case message := <-p.in:
			fmt.Printf("income: %#v", message)
			fmt.Println("")
		}
	}
}

func (p *Player) SendState(state *RoomState) {
	//TODO: send to front
	if state != nil {
		//TODO create norm state
		p.out <- &Message{"SINGLE", PayloadMessage{"STATE", "SOME-STATE"}}
	}
}

func (p *Player) SendMessage(message *Message) {
	p.out <- message
}

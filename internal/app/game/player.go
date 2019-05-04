package game

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type Player struct {
	conn *websocket.Conn
	ID   string
	in   chan *IncomeMessage
	out  chan *OutcomeMessage
	roomSingle *RoomSingle
	roomMulti *RoomMulti
	GameMode string
}

func NewPlayer(conn *websocket.Conn, id string) *Player {
	return &Player{
		conn: conn,
		ID:   id,
		in:   make(chan *IncomeMessage),
		out:  make(chan *OutcomeMessage),
		roomMulti: nil,
		roomSingle: nil,
	}
}

func (p *Player) Listen() {
	go func() {
		for {
			//слушаем фронт
			message := &IncomeMessage{}
			err := p.conn.ReadJSON(message)
			fmt.Println(err)
			p.in <- message
		}
	}()

	for {
		select {
		//получаем состояние игры от фронтов
		case message := <-p.in:
			//оработать, посчитать
			fmt.Printf("Front says: %#v", message)
			fmt.Println("")
			//if websocket.IsUnexpectedCloseError(err) {
			//	helpers.LogMsg("Player " + p.ID + " disconnected")
			//	return
			//}
			////message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
			//
			//if err != nil {
			//	helpers.LogMsg("Cannot read json: ", err)
			//	continue
			//}
			//стартовая инициализация, производится строго вначале один раз
			if message.Payload.Mode != "" {
				p.GameMode = message.Payload.Mode
				p.ID = message.Payload.Name
			}

			//обработка пришедшей команды
			if message.Payload.Command != "" {
				//process command
			}

	//		p.out <- &OutcomeMessage{"<3", OutPayloadMessage{}}

		case message := <-p.out:
			fmt.Printf("Back says: %#v", message)
			_ = p.conn.WriteJSON(message)
			//шлем всем фронтам текущее состояние
			if p.GameMode != "" {
				if p.roomSingle != nil {
					p.roomSingle.broadcast <- message
				} else {
					p.roomMulti.broadcast <- message
				}

			}
		}
	}
}

//func (p *Player) SendState(state *RoomState) {
//	//TODO: send to front
//	if state != nil {
//		//TODO create norm state
//		p.out <- &Message{"SINGLE", PayloadMessage{"STATE", "SOME-STATE"}}
//	}
//}
//
//func (p *Player) SendMessage(message *Message) {
//	p.out <- message
//}

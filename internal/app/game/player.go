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
		out:  make(chan *OutcomeMessage, 1),
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
			fmt.Println("ReadJSON error: ", err)
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

			//стартовая инициализация, производится строго вначале один раз
			if message.Payload.Mode != "" {
				p.GameMode = message.Payload.Mode
				p.ID = message.Payload.Name
				PingGame.AddPlayer(p)
			}

			//обработка пришедшей команды
			if message.Payload.Command != "" {
				//process command
			}

		case message := <-p.out:
			fmt.Printf("Back says: %#v", message)
			//шлем всем фронтам текущее состояние
			//switch message.Type {
			//	case START:
			//		_ = p.conn.WriteJSON(message)
			//	case WAIT:
			//		_ = p.conn.WriteJSON(message)
			//	case FINISH:
			//		_ = p.conn.WriteJSON(message)
			//	case STATE:
			//		_ = p.conn.WriteJSON(message)
			//	default:
			//		fmt.Println("Default in Player.Listen()")
			//}
			_ = p.conn.WriteJSON(message)
			//if p.GameMode != "" {
			//	if p.roomSingle != nil {
			//		p.roomSingle.broadcast <- message
			//	} else {
			//		p.roomMulti.broadcast <- message
			//	}
			//}
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
//func (p *Player) SendMessageSingle(message *OutcomeMessage) {
//	p.roomSingle.broadcast <- message
//}

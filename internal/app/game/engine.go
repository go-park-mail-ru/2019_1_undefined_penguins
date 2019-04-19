package game

import (
	"math"
)

const ownRandom = 0.25

func RotatePlayer(ps *PlayerState) {
	if ps.ClockwiseDirection {
		ps.ClockwiseDirection = false
	} else {
		ps.ClockwiseDirection = true
	}
}

func ShotPlayer(ps *PlayerState, b *BulletState) {
	if ps.Shoted {
		return
	}

	RecountBullet(ps, b)

	if ps.X == b.X && ps.Y == b.Y {
		//you loose
		ps.Shoted = true
		return
	}
}

func RecountBullet(ps *PlayerState, b *BulletState) {
	if ps.ClockwiseDirection {
		//TODO: create own random
		b.Alpha = ps.Alpha + ownRandom*100
	} else {
		b.Alpha = ps.Alpha - ownRandom*100
	}

	b.Radious = 20

	b.X = int(math.Floor(250 + math.Sin(b.Alpha*(math.Pi/180))*float64(b.Radious)))
	b.Y = int(math.Floor(250 - math.Cos(b.Alpha*(math.Pi/180))*float64(b.Radious)))
	b.Radious += 5
}

func CreateBullet(r *Room) *BulletState {
	return &BulletState{
		ID:    r.ID,
		X:     0,
		Y:     0,
		Alpha: 0,
	}
}

//recount coordinates (and maybe you win)
func ProcessGameSingle(r *Room) {
	var penguin *PlayerState
	//var gun *PlayerState
	for _, player := range r.state.Players {
		if player.Type == "GOOD" {
			penguin = player
		} else {
			//gun = player
		}
	}

	if penguin.Alpha == 360 {
		penguin.Alpha = 0
	}

	if penguin.Alpha == -1 {
		penguin.Alpha = 359
	}

	fishCount := 24
	for i := 0; i < fishCount; i++ {
		if penguin.Alpha == r.state.Fishes[i].Alpha {
			penguin.Score ++

			r.state.Fishes[i].Eaten = true
			break
		}
	}

	for i := 0; i < fishCount; i++ {
		if r.state.Fishes[i].Eaten == false {
			break
		}
			//you win???????
			// условие победы должно быть len(r.state.Fishes) == 0
			// но при этом, если рыбу съели, ее надо удалять. А еще здесь будет гг, потому что мы выделяем память сразу на 24 рыбы. Не надо так(
			//вопрос: как из функции game() понять, что мы победили и куда передать сообщение ??

		for t, _ := range r.state.Players {
				r.Players[t].out <- &Message{penguin.ID, "WIN"}
		}

		r.finish <- &Message{penguin.ID, "WIN"}
		return
	}

	if penguin.ClockwiseDirection {
		penguin.Alpha ++
	} else {
		penguin.Alpha --
	}

	alphaRad := degreesToRadians(penguin.Alpha)

	penguin.X = int(math.Floor(float64(r.state.Radious) + math.Sin(alphaRad)*float64(r.state.Radious)))
	penguin.Y = int(math.Floor(float64(r.state.Radious) - math.Cos(alphaRad)*float64(r.state.Radious)))
}

func degreesToRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}

//prepare room for game
func GameInit(r *Room) {
	fishCount := 24
	for i := 0; i < fishCount; i++ {
		Alpha := float64((360/24)*i)
		alphaRad := degreesToRadians(Alpha)
		X := int(math.Floor(float64(r.state.Radious) + math.Sin(alphaRad)*float64(r.state.Radious)))
		Y := int(math.Floor(float64(r.state.Radious) - math.Cos(alphaRad)*float64(r.state.Radious)))
		r.state.Fishes[i] = &FishState{i, X,Y,Alpha, false}
	}

	//for i, fish := range r.state.Fishes {
	//	fish.Alpha = float64((360/len(r.state.Fishes))*i)
	//	alphaRad := degreesToRadians(fish.Alpha)
	//	fish.X = int(math.Floor(float64(r.state.Radious) + math.Sin(alphaRad)*float64(r.state.Radious)))
	//	fish.Y = int(math.Floor(float64(r.state.Radious) - math.Cos(alphaRad)*float64(r.state.Radious)))
	//}

	for _, player := range r.state.Players {
		if player.Type == "GOOD" {
			player.Alpha = math.Floor(ownRandom*360)
			alphaRad := degreesToRadians(player.Alpha)
			player.X = int(math.Floor(float64(r.state.Radious) + math.Sin(alphaRad)*float64(r.state.Radious)))
			player.Y = int(math.Floor(float64(r.state.Radious) - math.Cos(alphaRad)*float64(r.state.Radious)))
		}
	}
}

func HandleCommand(r *Room, msg *Message) {
	//if message, ok := <-r.broadcast; ok {
		switch msg.Payload {
			case "SHOT":
				ShotPlayer(r.state.Players[msg.Type], r.state.Bullet)
				for t, player := range r.state.Players {
					if player.Shoted {
						r.Players[t].out <- &Message{msg.Type, "KILLED"}
					} else {
						r.Players[t].out <- &Message{msg.Type, "WIN"}
					}
				}
				//r.broadcast <- &Message{msg.Type, "KILLED"}
				//}
			case "ROTATE":
				RotatePlayer(r.state.Players[msg.Type])
				r.Players[msg.Type].out <- &Message{msg.Type, "KILLED"}
		}
		//for _, player := range r.Players {
		//	select {
		//	case player.out <- message:
		//		//if r.state.Players[message.Type].Shoted {
		//		//	r.Players[message.Type].out	<- &Message{message.Type, "KILLED"}
		//		//}
		//	default:
		//		close(player.out)
		//	}
		//}
	//}
}

func FinishGame(r *Room) {
	for _, player := range r.Players {
		message := &Message{player.ID, "GAME FINISHED"}
		player.conn.WriteJSON(message)
	}
}

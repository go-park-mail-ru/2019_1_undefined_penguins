package game

import (
	"math"
)

func RotatePlayer(ps *PlayerState) {
	if ps.ClockwiseDirection {
		ps.ClockwiseDirection = false
	} else {
		ps.ClockwiseDirection = true
	}
}

func ShotPlayer(ps *PlayerState, b *BulletState) {
	//helpers.LogMsg(b.X, b.Y)
	if ps.Shoted {
		return
	}

	RecountBullet(ps, b)
	//
	//helpers.LogMsg(b.X, b.Y)

	if ps.X == b.X && ps.Y == b.Y {
		ps.Shoted = true

		return
	}
}

func RecountBullet(ps *PlayerState, b *BulletState) {
	const ownRandom = 0.25
	if ps.ClockwiseDirection {
		//TODO: create own random
		b.Alpha = ps.Alpha + ownRandom*100
	} else {
		b.Alpha = ps.Alpha - ownRandom*100
	}

	b.X = b.X + int(math.Floor(math.Sin(b.Alpha*(math.Pi/180))))
	b.Y = b.Y - int(math.Floor(math.Cos(b.Alpha*(math.Pi/180))))
}

func CreateBullet(r *Room) BulletState {
	return BulletState{
		ID:    r.ID,
		X:     0,
		Y:     0,
		Alpha: 0,
	}
}

//recount coordinates
func ProcessGame() {

}

//prepare room for game
func GameInit(r *Room) {

}

func HandleCommand(r *Room) {
	if message, ok := <-r.broadcast; ok {
		switch message.Payload {
			case "SHOT":
				ShotPlayer(r.state.Players[message.Type], &r.state.Objects)
				//fmt.Println("Shooted ", r.state.Players[message.Type])
				//if r.state.Players[message.Type].Shoted {
				//	r.Players[message.Type].out	<- &Message{message.Type, "KILLED"}
				//}
			case "ROTATE":
				RotatePlayer(r.state.Players[message.Type])
				//fmt.Println("Rotated ", r.state.Players[message.Type])
		}
		for _, player := range r.Players {
			select {
			case player.out <- message:
				//if r.state.Players[message.Type].Shoted {
				//	r.Players[message.Type].out	<- &Message{message.Type, "KILLED"}
				//}
			default:
				close(player.out)
			}
		}
	}
}

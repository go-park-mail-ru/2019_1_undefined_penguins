package game

import (
	"2019_1_undefined_penguins/internal/pkg/helpers"
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
	helpers.LogMsg(b.X, b.Y)
	if ps.Shoted {
		return
	}

	RecountBullet(ps, b)

	helpers.LogMsg(b.X, b.Y)

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

	b.X = 20 + int(math.Floor(math.Sin(b.Alpha*(math.Pi/180))))
	b.Y = 20 - int(math.Floor(math.Cos(b.Alpha*(math.Pi/180))))
}

func CreateBullet(r *Room) BulletState {
	return BulletState{
		ID:    r.ID,
		X:     0,
		Y:     0,
		Alpha: 0,
	}
}

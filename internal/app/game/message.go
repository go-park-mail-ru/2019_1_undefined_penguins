package game

//from back to front
type OutcomeMessage struct {
	Type string `json:"type"`
	Payload OutPayloadMessage `json:"payload"`
}

type OutPayloadMessage struct {
	Penguin PenguinMessage `json:"penguin"`
	Bullet BulletMessage `json:"bullet"`
}

type PenguinMessage struct {
	Name string `json:"name"`
	Clockwise bool `json:"clockwise"`
	Alpha int `json:"alpha"`
}

type GunMessage struct {
	Name string `json:"name"`
	Bullet BulletMessage `json:"bullet"`
	Alpha int `json:"alpha"`
}

type BulletMessage struct {
	DistanceFromCenter int `json:"distance_from_center"`
	Alpha int `json:"alpha"`
}

//from front to back
type IncomeMessage struct {
	Type string `json:"type"`
	Payload InPayloadMessage `json:"payload,omitempty"`
}

type InPayloadMessage struct {
	Name string `json:"name"`
	Command string `json:"command"`
}

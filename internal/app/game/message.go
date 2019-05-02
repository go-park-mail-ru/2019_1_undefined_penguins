package game

type Message struct {
	Type string `json:"type"`
	Payload PayloadMessage `json:"payload,omitempty"`
}

type PayloadMessage struct {
	Player string `json:"player"`
	Command string `json:"command"`
}


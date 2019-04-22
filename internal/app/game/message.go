package game

type Message struct {
	Type string `json:"type"`
	Payload PayloadMessage `json:"payload,omitempty"`
}

type PayloadMessage struct {
	Player string `json:"player"`
	Command string `json:"command"`
}

type IncomeMessage struct {
	Type string `json:"type"`
	//Payload json.RawMessage `json:"payload,omitempty"`
	Payload string `json:"payload,omitempty"`
} 
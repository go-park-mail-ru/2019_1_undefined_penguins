package game

type Message struct {
	Type string `json:"type"`
	Payload interface{} `json:"payload,omitempty"`
}

type IncomeMessage struct {
	Type string `json:"type"`
	//Payload json.RawMessage `json:"payload,omitempty"`
	Payload string `json:"payload,omitempty"`
} 
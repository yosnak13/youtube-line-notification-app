package model

type Message struct {
	Type   string  `json:"type"`
	Hero   *Hero   `json:"hero"`
	Body   *Body   `json:"body"`
	Footer *Footer `json:"footer"`
}

func NewMessage(contentType string, hero *Hero, body *Body, footer *Footer) *Message {
	return &Message{
		Type:   contentType,
		Hero:   hero,
		Body:   body,
		Footer: footer,
	}
}

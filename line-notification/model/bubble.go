package model

type Bubble struct {
	Type   string  `json:"type"`
	Hero   *Hero   `json:"hero"`
	Body   *Body   `json:"body"`
	Footer *Footer `json:"footer"`
}

func NewBubble(contentType string, hero *Hero, body *Body, footer *Footer) *Bubble {
	return &Bubble{
		Type:   contentType,
		Hero:   hero,
		Body:   body,
		Footer: footer,
	}
}

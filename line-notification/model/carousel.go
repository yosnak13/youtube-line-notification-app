package model

type Carousel struct {
	Type    string    `json:"type"`
	Bubbles []*Bubble `json:"contents"`
}

func NewCarousel(contentType string, bubbles []*Bubble) *Carousel {
	return &Carousel{
		Type:    contentType,
		Bubbles: bubbles,
	}
}

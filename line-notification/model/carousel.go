package model

type Carousel struct {
	Type             string     `json:"type"`
	CarouselContents []*Message `json:"contents"`
}

func NewCarousel(contentType string, messages []*Message) *Carousel {
	return &Carousel{
		Type:             contentType,
		CarouselContents: messages,
	}
}

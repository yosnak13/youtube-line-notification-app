package model

type FlexMessage struct {
	Type     string    `json:"type"`
	AltText  string    `json:"altText"`
	Carousel *Carousel `json:"contents"`
}

func NewFlexMessage(contentType string, altText string, carousel *Carousel) *FlexMessage {
	return &FlexMessage{
		Type:     contentType,
		AltText:  altText,
		Carousel: carousel,
	}
}

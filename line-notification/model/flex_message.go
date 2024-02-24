package model

type FlexMessage struct {
	Type     string    `json:"type"`
	AltText  string    `json:"altText"`
	Carousel []Content `json:"contents"`
}

func NewFlexMessage(contentType string, altText string, carousel []Content) *FlexMessage {
	return &FlexMessage{
		Type:     contentType,
		AltText:  altText,
		Carousel: carousel,
	}
}

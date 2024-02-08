package model

type FooterContent struct {
	Type   string  `json:"type"`
	Style  string  `json:"style"`
	Height string  `json:"height"`
	Action *Action `json:"action"`
}

func NewFooterContent(footerType string, style string, height string, action *Action) *FooterContent {
	return &FooterContent{
		Type:   footerType,
		Style:  style,
		Height: height,
		Action: action,
	}
}

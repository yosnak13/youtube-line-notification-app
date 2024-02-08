package model

type FooterContent struct {
	Type   string `json:"type"`
	Style  string `json:"style"`
	Height string `json:"height"`
	Action Action `json:"action"`
}

func NewFooterContent() *FooterContent {
	return &FooterContent{
		Type:   "button",
		Style:  "link",
		Height: "sm",
		Action: *NewActionForFooter(),
	}
}

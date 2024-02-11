package model

type Content struct {
	Type     string    `json:"type"`
	Text     string    `json:"text,omitempty"`
	Weight   string    `json:"weight,omitempty"`
	Size     string    `json:"size,omitempty"`
	Wrap     bool      `json:"wrap,omitempty"`
	Layout   string    `json:"layout,omitempty"`
	Margin   string    `json:"margin,omitempty"`
	Spacing  string    `json:"spacing,omitempty"`
	Color    string    `json:"color,omitempty"`
	Flex     int       `json:"flex,omitempty"`
	Action   Action    `json:"action,omitempty"`
	Contents []Content `json:"contents,omitempty"`
}

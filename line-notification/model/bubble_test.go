package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBubble(t *testing.T) {
	action := NewAction("test", "test", "test")
	hero := NewHero("image", "https://example.com", "full", "20:30", "cover", action)

	var contents []*Content
	for i := 0; i < 2; i++ {
		content := &Content{
			Type:  "text",
			Text:  "ch",
			Flex:  i,
			Wrap:  true,
			Size:  "sm",
			Color: "aaaaaa",
		}
		contents = append(contents, content)
	}
	body := NewBody("type", "xl", contents)

	footerContent := NewFooterContent("button", "link", "sm", action)
	footer := NewFooter("box", "vertical", "sm", []*FooterContent{footerContent}, 1)

	contentType := "bubble"

	expect := &Bubble{
		Type:   contentType,
		Hero:   hero,
		Body:   body,
		Footer: footer,
	}

	actual := NewBubble(contentType, hero, body, footer)

	assert.Equal(t, expect, actual)
}

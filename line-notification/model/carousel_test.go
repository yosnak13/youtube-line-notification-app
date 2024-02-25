package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCarousel(t *testing.T) {
	contentType := "carousel"

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
			Color: "#aaaaaa",
		}
		contents = append(contents, content)
	}
	body := NewBody("type", "xl", contents)
	footerContent := NewFooterContent("button", "link", "sm", action)
	footer := NewFooter("box", "vertical", "sm", []*FooterContent{footerContent}, 1)

	bubble := NewBubble("bubble", hero, body, footer)
	bubbles := []*Bubble{bubble}

	expect := &Carousel{
		Type:    contentType,
		Bubbles: bubbles,
	}

	actual := NewCarousel(contentType, bubbles)

	assert.Equal(t, expect, actual)
}

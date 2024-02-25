package model

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewFlexMessage(t *testing.T) {
	contentType := "flex"
	altText := "本日の動画です。"

	var carousels []*Carousel
	var contents []*Content
	for i := 0; i < 2; i++ {
		action := NewAction("test", "test", "test")
		url := fmt.Sprintf("https://example.com/%v", i)
		hero := NewHero("image", url, "full", "20:30", "cover", action)

		for k := 0; k < 2; k++ {
			content := &Content{
				Type:  "text",
				Text:  "ch",
				Flex:  k,
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

		carousel := NewCarousel("carousel", bubbles)
		carousels = append(carousels, carousel)
	}

	expect := &FlexMessage{
		Type:     contentType,
		AltText:  altText,
		Carousel: carousels,
	}

	actual := NewFlexMessage(contentType, altText, carousels)

	assert.Equal(t, expect, actual)
}

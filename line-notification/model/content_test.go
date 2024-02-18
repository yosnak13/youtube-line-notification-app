package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewContentMovieTitle(t *testing.T) {

	contentType := "text"
	text := "タイトル"
	weight := "bold"
	size := "xl"
	wrap := true

	expect := &Content{
		Type:   contentType,
		Text:   text,
		Weight: weight,
		Size:   size,
		Wrap:   wrap,
	}

	actual := NewContentMovieTitle(contentType, text, weight, size, wrap)

	assert.Equal(t, expect, actual)
}

func TestNewContentBodyContainer(t *testing.T) {

	contentType := "text"
	layout := "vertical"
	margin := "lg"
	spacing := "sm"
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

	expect := &Content{
		Type:     contentType,
		Layout:   layout,
		Margin:   margin,
		Spacing:  spacing,
		Contents: contents,
	}
	actual := NewContentBodyContainer(contentType, layout, margin, spacing, contents)

	assert.Equal(t, expect, actual)
}

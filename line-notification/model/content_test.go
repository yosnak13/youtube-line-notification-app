package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewContentMovieProperty(t *testing.T) {
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

	actual := NewContentMovieProperty(contentType, text, weight, size, wrap)

	assert.Equal(t, expect, actual)
}

func TestNewContentMovieValue(t *testing.T) {
	contentType := "text"
	layout := "vertical"
	margin := "lg"
	spacing := "sm"
	var contents []*Content
	for i := 0; i < 2; i++ {
		content := &Content{
			Type:  "text",
			Text:  "ch",
			Flex:  1,
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
	actual := NewContentMovieValue(contentType, layout, margin, spacing, contents)

	assert.Equal(t, expect, actual)
}

func TestNewContentBodyBlockChannelRoot(t *testing.T) {
	contentType := "box"
	layout := "baseline"
	var contents []*Content
	for i := 0; i < 2; i++ {
		content := &Content{
			Type:  "text",
			Text:  "ch",
			Flex:  1,
			Wrap:  true,
			Size:  "sm",
			Color: "aaaaaa",
		}
		contents = append(contents, content)
	}

	expect := &Content{
		Type:     contentType,
		Layout:   layout,
		Contents: contents,
	}

	actual := NewContentBodyBlockChannelRoot(contentType, layout, contents)

	assert.Equal(t, expect, actual)
}

func TestNewContentBodyBlockChannelPropertyValue(t *testing.T) {
	contentType := "text"
	text := "ch"
	flex := 1
	wrap := true
	size := "sm"
	color := "aaaaaa"

	expect := &Content{
		Type:  contentType,
		Text:  text,
		Flex:  flex,
		Wrap:  wrap,
		Size:  size,
		Color: color,
	}

	actual := NewContentBodyBlockChannelPropertyValue(contentType, text, flex, wrap, size, color)

	assert.Equal(t, expect, actual)
}

func TestNewContentBodyBlockUrlRoot(t *testing.T) {
	contentType := "box"
	layout := "baseline"
	spacing := "sm"
	var contents []*Content
	for i := 0; i < 2; i++ {
		content := &Content{
			Type:  "text",
			Text:  "ch",
			Flex:  1,
			Wrap:  true,
			Size:  "sm",
			Color: "aaaaaa",
		}
		contents = append(contents, content)
	}

	expect := &Content{
		Type:     contentType,
		Layout:   layout,
		Spacing:  spacing,
		Contents: contents,
	}

	actual := NewContentBodyBlockUrlRoot(contentType, layout, spacing, contents)

	assert.Equal(t, expect, actual)
}

func TestNewContentBodyBlockUrlProperty(t *testing.T) {
	contentType := "text"
	movieUrl := "https://example.com"
	color := "#aaaaaa"
	size := "sm"
	flex := 1

	expect := &Content{
		Type:  contentType,
		Text:  movieUrl,
		Color: color,
		Size:  size,
		Flex:  flex,
	}

	actual := NewContentBodyBlockUrlProperty(contentType, movieUrl, color, size, flex)

	assert.Equal(t, expect, actual)
}

func TestNewContentBodyBlockUrlValue(t *testing.T) {
	contentType := "text"
	text := "test"
	wrap := true
	color := "#aaaaaa"
	size := "sm"
	flex := 1
	action := NewAction("uri", "", "https://example.com")

	expect := &Content{
		Type:   contentType,
		Text:   text,
		Wrap:   wrap,
		Color:  color,
		Size:   size,
		Flex:   flex,
		Action: action,
	}

	actual := NewContentBodyBlockUrlValue(contentType, text, wrap, color, size, flex, action)

	assert.Equal(t, expect, actual)
}

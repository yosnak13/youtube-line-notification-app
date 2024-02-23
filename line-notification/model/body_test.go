package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewBody(t *testing.T) {
	contentType := "type"
	layout := "xl"
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

	expect := &Body{
		Type:     contentType,
		Layout:   layout,
		Contents: contents,
	}

	actual := NewBody(contentType, layout, contents)

	assert.Equal(t, expect, actual)
}

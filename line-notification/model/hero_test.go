package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHero(t *testing.T) {

	typeOfHero := "image"
	thumbnail := "sample@example.com"
	size := "full"
	aspectRatio := "20:30"
	aspectMode := "cover"
	action := NewAction("test", "test", "test")

	expect := &Hero{
		Type:         typeOfHero,
		ThumbnailURL: thumbnail,
		Size:         size,
		AspectRatio:  aspectRatio,
		AspectMode:   aspectMode,
		Action:       action,
	}
	actual := NewHero(typeOfHero, thumbnail, size, aspectRatio, aspectMode, action)

	assert.Equal(t, expect, actual)
}

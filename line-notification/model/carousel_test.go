package model

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/stretchr/testify/assert"
	mocks "line-notification/model/line_bot/mock"
	"testing"
)

func TestNewCarousel(t *testing.T) {

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

	contentType := "carousel"
	expect := &Carousel{
		Type:    contentType,
		Bubbles: bubbles,
	}

	actual := NewCarousel(contentType, bubbles)

	assert.Equal(t, expect, actual)
}

func TestCarousel_RequestLineMessagingAPI(t *testing.T) {
	t.Setenv("LineBotChannelSecret", "test_secret")
	t.Setenv("LineBotChannelToken", "test_token")

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	action := *NewAction("uri", "", "https://example.com")
	hero := *NewHero("image", "https://example.com", "full", "20:30", "cover", &action)
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
	body := *NewBody("type", "xl", contents)
	footerContent := *NewFooterContent("button", "link", "sm", &action)
	footer := *NewFooter("box", "vertical", "sm", []*FooterContent{&footerContent}, 1)

	bubble := NewBubble("bubble", &hero, &body, &footer)
	bubbles := []*Bubble{bubble}

	contentType := "carousel"
	validCarousel := NewCarousel(contentType, bubbles)

	t.Run("正常系: ブロードキャスト成功", func(t *testing.T) {
		mockCall := mocks.NewMockLineBroadcastMessageCall(ctrl)
		mockCall.EXPECT().Do().Return(&linebot.BasicResponse{}, nil)

		mockBot := mocks.NewMockLineClient(ctrl)
		mockBot.EXPECT().
			BroadcastMessage(gomock.Any()).
			Return(mockCall)

		err := validCarousel.RequestLineMessagingAPI(mockBot)
		assert.NoError(t, err)
	})

	t.Run("異常系: ブロードキャスト失敗", func(t *testing.T) {
		mockCall := mocks.NewMockLineBroadcastMessageCall(ctrl)
		mockCall.EXPECT().Do().Return(nil, errors.New("api error"))

		mockBot := mocks.NewMockLineClient(ctrl)
		mockBot.EXPECT().
			BroadcastMessage(gomock.Any()).
			Return(mockCall)

		err := validCarousel.RequestLineMessagingAPI(mockBot)
		assert.ErrorContains(t, err, "api error")
	})
}

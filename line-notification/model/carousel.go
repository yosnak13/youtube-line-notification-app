package model

import (
	"encoding/json"
	"fmt"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"line-notification/model/line_bot"
)

const AnnounceTodayMovie string = "本日の動画です"

type Carousel struct {
	Type    string    `json:"type"`
	Bubbles []*Bubble `json:"contents"`
}

func NewCarousel(contentType string, bubbles []*Bubble) *Carousel {
	return &Carousel{
		Type:    contentType,
		Bubbles: bubbles,
	}
}

func (c Carousel) RequestLineMessagingAPI(botProvider line_bot.LineClient) error {
	messageJSON, err := json.Marshal(c)

	flexContainer, err := linebot.UnmarshalFlexMessageJSON(messageJSON)
	if err != nil {
		fmt.Printf("failed to UnmarshalFlexMessageJSON %s", err)
		return err
	}

	flexMessage := linebot.NewFlexMessage(AnnounceTodayMovie, flexContainer)
	if _, err := botProvider.BroadcastMessage(flexMessage).Do(); err != nil {
		fmt.Printf("failed to BroadcastMessage %s", err)
		return err
	}

	fmt.Println("Messaging request is succeeded!")
	return nil
}

package line_bot

import "github.com/line/line-bot-sdk-go/v7/linebot"

type LineBroadcastMessageCall interface {
	Do() (*linebot.BasicResponse, error)
}

type LineClient interface {
	BroadcastMessage(message linebot.SendingMessage) LineBroadcastMessageCall
}

type LineBotClient struct {
	client *linebot.Client
}

func NewLineBotClient(channelSecret, channelToken string) (LineClient, error) {
	client, err := linebot.New(channelSecret, channelToken)
	return &LineBotClient{client: client}, err
}

func (c *LineBotClient) BroadcastMessage(msg linebot.SendingMessage) LineBroadcastMessageCall {
	return c.client.BroadcastMessage(msg)
}

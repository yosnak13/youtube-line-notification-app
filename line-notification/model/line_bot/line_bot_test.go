package line_bot_test

import (
	"github.com/golang/mock/gomock"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/stretchr/testify/assert"
	"line-notification/model/line_bot"
	mock "line-notification/model/line_bot/mock"
	"os"
	"testing"
)

func TestNewLineBotClient_Success(t *testing.T) {
	t.Setenv("LINE_CHANNEL_SECRET", "test_secret")
	t.Setenv("LINE_CHANNEL_TOKEN", "test_token")

	client, err := line_bot.NewLineBotClient(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_TOKEN"),
	)

	assert.NoError(t, err)
	assert.NotNil(t, client)

	_, ok := client.(line_bot.LineClient)
	assert.True(t, ok)
}

func TestNewLineBotClient_InvalidCredentials(t *testing.T) {
	t.Setenv("LineBotChannelSecret", "")
	t.Setenv("LineBotChannelToken", "")

	t.Run("空の認証情報", func(t *testing.T) {
		_, err := line_bot.NewLineBotClient("", "")
		assert.Error(t, err)
	})
}

func TestLineBotClient_BroadcastMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCall := mock.NewMockLineBroadcastMessageCall(ctrl)
	mockClient := mock.NewMockLineClient(ctrl)

	t.Run("BroadcastMessage Success", func(t *testing.T) {
		expectedResp := &linebot.BasicResponse{}
		testMessage := linebot.NewTextMessage("test")

		mockClient.EXPECT().
			BroadcastMessage(testMessage).
			Return(mockCall).
			Times(1)

		mockCall.EXPECT().
			Do().
			Return(expectedResp, nil).
			Times(1)

		call := mockClient.BroadcastMessage(testMessage)
		resp, err := call.Do()

		// 検証
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if resp != expectedResp {
			t.Error("response mismatch")
		}
	})
}

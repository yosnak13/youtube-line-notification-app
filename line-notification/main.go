package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"google.golang.org/api/youtube/v3"
	"line-notification/model"
	"line-notification/model/line_bot"
	"log"
	"net/http"
	"os"

	"google.golang.org/api/googleapi/transport"
)

const (
	Blank                  string = ""
	ContentUri             string = "uri"
	ContentUrlLarge        string = "URL"
	ContentImg             string = "image"
	ContentText            string = "text"
	ContentBox             string = "box"
	SpaceSmall             string = "sm"
	SrgbGray               string = "#aaaaaa"
	LayoutBaseLine         string = "baseline"
	LayoutVertical         string = "vertical"
	FlexOne                int    = 1
	FlexFive               int    = 5
	NoFlex                 int    = 0
	IsWrap                 bool   = true
	AnnounceTodayMovie     string = "本日の動画です"
	AnnounceTapHere        string = "動画はこちらをタップ"
	AnnounceGoTopOfYouTube string = "Youtubeトップへ"
)

func main() {
	lambda.Start(handler)
}

func handler() {
	// 暫定的に下手打ち。DynamoDBに移行する。
	channelIDs := []string{
		"UCPVr7clenPjpD7WNsSI3UBQ", // レトルト
		"UCZMRuagdTBKmmrFtSMN48Xw", // 牛沢
		"UCWcEgYIOqq1BVr4Qm1sPuVg", // ガッチマン
		"UCcAGc7BqTIyEjeXZM6QhGiQ", // towaco
		// 追加のチャンネルIDをここに追加
	}

	apiKey := os.Getenv("API_KEY")
	client := &http.Client{
		Transport: &transport.APIKey{Key: apiKey},
	}

	service, err := youtube.New(client)
	if err != nil {
		log.Fatalf("Error creating new YouTube client: %v", err)
	}

	var bubbles []*model.Bubble

	for _, channelID := range channelIDs {
		call := service.Search.List([]string{"snippet"}).ChannelId(channelID).Type("video").Order("date").MaxResults(1)
		response, err := call.Do()
		if err != nil {
			log.Fatalf("Error making search API call for channel %s: %v", channelID, err)
		}

		item := response.Items[0] // MaxResultsが1なので、配列の最初の要素でよい
		movieTitle := item.Snippet.Title
		thumbnail := item.Snippet.Thumbnails
		channelName := item.Snippet.ChannelTitle
		videoID := item.Id.VideoId

		thumbnailURL := thumbnail.Default.Url
		movieURL := fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoID)

		bubble := buildBubble(movieTitle, thumbnailURL, channelName, movieURL)
		bubbles = append(bubbles, bubble)
	}

	typeCarousel := "carousel"
	carousel := *model.NewCarousel(typeCarousel, bubbles)

	bot, _ := line_bot.NewLineBotClient(os.Getenv("LineBotChannelSecret"), os.Getenv("LineBotChannelToken"))
	err = carousel.RequestLineMessagingAPI(bot)
}

func buildBubble(movieTitle string, thumbnailURL string, channelTitle string, movieURL string) *model.Bubble {
	hero := buildHero(thumbnailURL, movieURL)
	body := buildBody(movieTitle, channelTitle, movieURL)
	footer := buildFooter()

	typeBubble := "bubble"
	bubble := *model.NewBubble(typeBubble, hero, body, footer)
	return &bubble
}

func buildHero(thumbnailURL string, movieURL string) *model.Hero {
	fullSize := "full"
	aspect := "1:1"
	aspectMode := "cover"

	action := *model.NewAction(ContentUri, Blank, movieURL)
	hero := *model.NewHero(ContentImg, thumbnailURL, fullSize, aspect, aspectMode, &action)
	return &hero
}

func buildBody(movieTitle string, channelTitle string, movieURL string) *model.Body {
	urlProperty := *model.NewContentBodyBlockUrlProperty(ContentText, ContentUrlLarge, SrgbGray, SpaceSmall, FlexOne)

	urlValueAction := *model.NewAction(ContentUri, Blank, movieURL)

	darkGray := "#666666"
	urlValue := *model.NewContentBodyBlockUrlValue(ContentText, AnnounceTapHere, IsWrap, darkGray, SpaceSmall, FlexFive, &urlValueAction)
	urlComponents := []*model.Content{&urlProperty, &urlValue}
	urlRootComponent := *model.NewContentBodyBlockUrlRoot(ContentBox, LayoutBaseLine, SpaceSmall, urlComponents)

	textChannel := "ch"
	channelProperty := *model.NewContentBodyBlockChannelPropertyValue(ContentText, textChannel, FlexOne, IsWrap, SpaceSmall, SrgbGray)
	channelValue := *model.NewContentBodyBlockChannelPropertyValue(ContentText, channelTitle, FlexFive, IsWrap, SpaceSmall, SrgbGray)
	channelNameComponents := []*model.Content{&channelProperty, &channelValue}
	channelRootComponent := *model.NewContentBodyBlockChannelRoot(ContentBox, LayoutBaseLine, channelNameComponents)

	movieInfo := []*model.Content{&channelRootComponent, &urlRootComponent}

	largeMargin := "lg"
	movieComponent := *model.NewContentMovieValue(ContentBox, LayoutVertical, largeMargin, SpaceSmall, movieInfo)
	fontWeightBold := "bold"
	sizeXl := "xl"
	titleComponent := *model.NewContentMovieProperty(ContentText, movieTitle, fontWeightBold, sizeXl, IsWrap)

	bodyComponents := []*model.Content{&titleComponent, &movieComponent}

	body := *model.NewBody(ContentBox, LayoutVertical, bodyComponents)
	return &body
}

func buildFooter() *model.Footer {
	youTubeTopURL := "https://youtube.com"
	action := *model.NewAction(ContentUri, AnnounceGoTopOfYouTube, youTubeTopURL)

	typeButton := "button"
	styleLink := "link"
	footerContent := *model.NewFooterContent(typeButton, styleLink, SpaceSmall, &action)
	content := []*model.FooterContent{&footerContent}
	footer := *model.NewFooter(ContentBox, LayoutVertical, SpaceSmall, content, NoFlex)
	return &footer
}

# アプリケーション概要

## Youtube通知アプリ

気に入ったYoutubeチャンネルに動画がアップロードされた時、LINEに通知するサービス。  
毎時17:30に、LINEへ通知する。

![RPReplay_Final1711515036](https://github.com/yosnak13/youtube-line-notification-app/assets/64535376/f6427cc3-e56b-40d7-9c2e-c911174d5a38)


## 開発背景

YouTubeを見る機会が多いが、毎度登録チャンネルの最新動画を確認するのが面倒なため、毎日1回定期的にLINEに通知するようにして、確認先を1箇所にとどめられるようにした。  
LINEチャンネルを作成し、チャンネルを登録することで利用できる。

## 採用技術

### 言語

- Golang
  - Lambda関数に使用

### インフラ

主にAWSサービスを活用
- AWS SAM（Serverless Application Model）
  - 開発・ビルド・デプロイに使用
- Lambda
- EventBridge
- S3
  - SAMでビルドしたアプリケーションのデプロイ先
- CloudFormation
  - インフラリソースのプロビジョニング
  - SAM templateに記載したものがCloudFormationテンプレートに変換されデプロイされる

### ミドルウェア

- Docker
  - SAMを利用した開発に必須

### APIサービス

- [LINE Developers](https://developers.line.biz/ja/)
  - LINEチャンネル作成に使用
  - [LINE Messaging API](https://developers.line.biz/en/services/messaging-api/)
- [YouTube Data API](https://developers.google.com/youtube/v3?hl=ja)


## アウトプット

- [インフラ構成図](https://github.com/yosnak13/youtube-line-notification-app/issues/4#issuecomment-1909375809)
- [Lambda関数](https://github.com/yosnak13/youtube-line-notification-app/blob/main/line-notification/main.go)
- [DynamoDBテーブル定義](https://github.com/yosnak13/youtube-line-notification-app/issues/8)
  - 未実装。改修で実装予定
- [シーケンス図](https://github.com/yosnak13/youtube-line-notification-app/issues/7#issuecomment-1909590119)
  - DynamoDB部分が未実装。改修で実装予定
- [LINE MessagingAPI インターフェース](https://github.com/yosnak13/youtube-line-notification-app/issues/9#issuecomment-1908007166)
- [LINEアイコン](https://github.com/yosnak13/youtube-line-notification-app/issues/27#issuecomment-1974153867)

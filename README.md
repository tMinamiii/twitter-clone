# Twitter Clone

## Tools

以下のツールを用いて開発を行います。

- air (https://github.com/air-verse/air)
- gomock (https://github.com/uber-go/mock)
- golang-migrate (https://github.com/golang-migrate/migrate)
- staticcheck (https://staticcheck.io/docs/)
- pre-commit (https://github.com/pre-commit/pre-commit)

```sh
go install github.com/air-verse/air@latest
go install go.uber.org/mock/mockgen@latest
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
go install honnef.co/go/tools/cmd/staticcheck@latest
brew install pre-commit # or pip install pre-commit
pre-commit install
```

## ローカル環境立ち上げ

client

```sh
cd clinet
cp dot.env.local .env.local
npm install
npm run dev
```

server

```sh
cd server
cp dot.env.local .env.local
docker compose up -d
go mod download

# db 起動後
make migrate-up
make seed
make dev
```

## テストの実行

client

```sh
cd server
npm run test
```

server

```sh
cd server
docker compose up -d # db起動済みなら不要
make migrate-up # migration済みなら不要
make test
```

## その他

### server

#### rdb, usecase パッケージを追加・変更する場合

rdb, usecase パッケージは interface に沿って mock ファイルを生成しています。
パッケージに新規ファイルを作成する場合は、ファイル先頭行に下記を記述してください。

```go
//go:generate mockgen -source=$GOFILE -destination=./mock/$GOFILE -package=mock_$GOPACKAGE
````

新規ファイル作成後や既存ファイルの interface を変更した際は generate コマンドを実行してください。

```sh
cd sever
make generate
```

#### migration について

新規マイグレーションファイルを作成する場合は、golan-migrate を用いて生成して下し。生成コマンドは下記のとおりです。

```sh
migrate create -ext sql -dir server/migrations <ファイル名>
```

#### seeder が作成するユーザー一覧

`make seed` することで以下のユーザーが作成されます。

| account id          | username          |
| ------------------- | ----------------- |
| 01_taro_yamada      | 01\_山田 太郎     |
| 02_hanako_sato      | 02\_佐藤 花子     |
| 03_ichiro_suzuki    | 03\_鈴木 一郎     |
| 04_yuki_tanaka      | 04\_田中 由紀     |
| 05_kei_kobayashi    | 05\_小林 慶       |
| 06_miki_yoshida     | 06\_吉田 美樹     |
| 07_satoshi_watanabe | 07\_渡辺 智       |
| 08_kana_ito         | 08\_伊藤 香奈     |
| 09_naoki_yamamoto   | 09\_山本 直樹     |
| 10_haruto_nakamura  | 10\_中村 春人     |
| 11_mio_matsumoto    | 11\_松本 美緒     |
| 12_yuto_inoue       | 12\_井上 優斗     |
| 13_ayaka_kimura     | 13\_木村 綾香     |
| 14_ryota_shimizu    | 14\_清水 亮太     |
| 15_erika_sasaki     | 15\_佐々木 絵里香 |
| 16_sho_kondo        | 16\_近藤 翔       |
| 17_kana_fujimoto    | 17\_藤本 香奈     |
| 18_akira_takahashi  | 18\_高橋 晃       |
| 19_nana_morita      | 19\_森田 奈々     |
| 20_kenta_murakami   | 20\_村上 健太     |

package main

import (
	"context"
	"fmt"
	"log"
	"tMinamiii/Tweet/domain"
	"tMinamiii/Tweet/infra/rdb"
	"tMinamiii/Tweet/project"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println(project.Root())
	err := godotenv.Load(project.Root() + "/.env.local")
	if err != nil {
		log.Fatalln(err)
	}

	sess := rdb.GetTweetSession()
	sess.Exec("TRUNCATE TABLE users")
	sess.Exec("TRUNCATE TABLE posts")
	sess.Exec("TRUNCATE TABLE follows")

	tx, err := sess.Begin()
	if err != nil {
		log.Fatalln(err)
	}
	defer tx.RollbackUnlessCommitted()

	ctx := context.Background()
	usersLength := len(usersData)
	// contentsLength := len(contentsData)
	// insert users
	for _, v := range usersData {
		tx.InsertInto("users").
			Pair("account_id", v.AccountID).
			Pair("username", v.Username).
			ExecContext(ctx)
	}

	users := &[]domain.User{}
	_, err = tx.Select("*").From("users").LoadContext(ctx, users)
	if err != nil {
		log.Fatalln(err)
	}

	p := rdb.NewPostsTable()
	// insert post
	for i, v := range contentsData {
		userID := i%usersLength + 1
		_, err = p.CreateTx(ctx, tx, int64(userID), v)
		if err != nil {
			log.Fatalln(err)
		}
		time.Sleep(time.Second)
	}

	tx.Commit()
}

var usersData = []domain.User{
	{AccountID: "01_taro_yamada", Username: "01_山田 太郎"},
	{AccountID: "02_hanako_sato", Username: "02_佐藤 花子"},
	{AccountID: "03_ichiro_suzuki", Username: "03_鈴木 一郎"},
	{AccountID: "04_yuki_tanaka", Username: "04_田中 由紀"},
	{AccountID: "05_kei_kobayashi", Username: "05_小林 慶"},
	{AccountID: "06_miki_yoshida", Username: "06_吉田 美樹"},
	{AccountID: "07_satoshi_watanabe", Username: "07_渡辺 智"},
	{AccountID: "08_kana_ito", Username: "08_伊藤 香奈"},
	{AccountID: "09_naoki_yamamoto", Username: "09_山本 直樹"},
	{AccountID: "10_haruto_nakamura", Username: "10_中村 春人"},
	{AccountID: "11_mio_matsumoto", Username: "11_松本 美緒"},
	{AccountID: "12_yuto_inoue", Username: "12_井上 優斗"},
	{AccountID: "13_ayaka_kimura", Username: "13_木村 綾香"},
	{AccountID: "14_ryota_shimizu", Username: "14_清水 亮太"},
	{AccountID: "15_erika_sasaki", Username: "15_佐々木 絵里香"},
	{AccountID: "16_sho_kondo", Username: "16_近藤 翔"},
	{AccountID: "17_kana_fujimoto", Username: "17_藤本 香奈"},
	{AccountID: "18_akira_takahashi", Username: "18_高橋 晃"},
	{AccountID: "19_nana_morita", Username: "19_森田 奈々"},
	{AccountID: "20_kenta_murakami", Username: "20_村上 健太"},
}

var contentsData = []string{
	"今日は晴れていて気持ちいいですね！",
	"新しいカフェを見つけました。コーヒーが美味しい！",
	"映画を見に行ったけど、すごく感動しました。",
	"最近読んだ本がとても面白かった。おすすめです。",
	"ランニングを始めました。気分がスッキリします。",
	"友達と久しぶりに会って楽しかった！",
	"新しいレシピに挑戦してみました。意外と簡単でした。",
	"旅行に行きたいなぁ。次はどこに行こうかな？",
	"今日は仕事がたくさんあって忙しかった。",
	"最近ハマっているドラマが面白すぎる！",
	"ペットの写真をシェアします。かわいいでしょ？",
	"今日は特別な日。大切な人と過ごせて幸せです。",
	"美味しいご飯を食べに行きました。最高！",
	"新しい趣味を見つけました。楽しい！",
	"今日はゆっくり休む日。リラックスしています。",
	"友達の誕生日パーティーに行ってきました。",
	"新しい音楽を発見しました。リピート中。",
	"今日は家でまったりと映画鑑賞。",
	"最近健康に気を使っています。運動大事。",
	"今日のランチは最高でした。また行きたい。",
	"新しい服を買いました。お気に入りです。",
	"天気がいいので散歩に出かけました。気持ちいい！",
	"今日は仕事がうまくいった。達成感があります。",
	"最近始めた趣味が楽しくて仕方ないです。",
	"新しい映画を見て感動しました。おすすめです。",
	"友達とカフェでおしゃべり。楽しい時間でした。",
	"今日は家でのんびり。リフレッシュできました。",
	"新しいレストランに行きました。美味しかった！",
	"最近の出来事をシェアします。今日はいい日でした。",
	"次の旅行先を計画中。ワクワクします！",
	"今日は久しぶりに早起きしました。朝の空気は気持ちいいですね。",
	"今日は料理をしてみたけど、意外と上手にできました。",
	"新しいゲームを始めました。かなりハマってます。",
	"今日はジムに行って汗を流しました。すごく気持ちよかった！",
	"新しいスマホを買いました。使い心地が最高です。",
	"今日は家でまったり読書。リラックスできました。",
	"最近の仕事が忙しすぎて疲れています。でも頑張ります！",
	"新しいプロジェクトが始まりました。ワクワクしています。",
	"今日は家族と一緒に過ごしました。楽しかったです。",
	"新しい靴を買いました。おしゃれで気に入ってます。",
	"今日は外でランチ。美味しかったです。",
	"最近ヨガを始めました。リラックスできていい感じです。",
	"今日は映画館で映画を見ました。大きなスクリーンはやっぱりいいですね。",
	"新しい友達ができました。これからが楽しみです。",
	"今日は少し早めに帰宅。家でゆっくり過ごします。",
	"最近の天気が良くて気分も晴れやかです。",
	"新しいバッグを買いました。どこに持って行こうかな？",
	"今日は公園でピクニック。楽しい一日でした。",
	"最近の出来事を振り返ると、いいことばかりです。",
	"新しいチャレンジに挑戦中。結果が楽しみです。",
	"今日は久しぶりに友達と会って、楽しい時間を過ごしました。",
	"新しいヘアスタイルに挑戦しました。気に入ってます！",
	"今日は家で料理をして、一日中キッチンにいました。",
	"最近、写真を撮るのが趣味になりました。面白い！",
	"新しいスポーツを始めました。体を動かすのは気持ちいいですね。",
	"今日は図書館で静かに過ごしました。落ち着きます。",
	"新しいカメラを買いました。写真撮影が楽しみです。",
	"今日は久しぶりに外でランチ。リフレッシュできました。",
	"最近、お絵かきを始めました。自分でも意外と上手！",
	"新しい靴を買ったので、今日はたくさん歩きました。",
	"今日は家で映画三昧。至福の時間です。",
	"新しいガジェットを試してみました。便利！",
	"最近、筋トレを始めました。体が引き締まってきました。",
	"今日は仕事が早く終わったので、ゆっくりしています。",
	"新しい音楽を発見。ヘビロテ中です！",
	"最近、お菓子作りにハマってます。楽しい！",
	"今日は美術館に行ってきました。素晴らしい作品がたくさんありました。",
	"新しいスマホケースを買いました。お気に入りです。",
	"今日はのんびり読書。時間がゆっくり流れています。",
	"最近、夜ランニングを始めました。気持ちいい！",
	"今日は公園でジョギング。リフレッシュできました。",
	"新しいバッグを買いました。どこに持って行こうかな？",
	"今日は友達とドライブ。楽しい一日でした。",
	"最近、料理の腕が上がってきました。自分でも驚きです。",
	"新しいイヤホンを買いました。音質が最高です。",
	"今日は家で映画を見ました。リラックスできました。",
	"最近、DIYにハマってます。新しい家具を作りました。",
	"今日は友達と一緒に過ごしました。楽しかったです。",
	"新しいシャツを買いました。お気に入りです。",
	"今日は家でのんびり。リフレッシュできました。",
	"最近、アートに興味を持ち始めました。面白い！",
	"新しい腕時計を買いました。スタイリッシュで気に入ってます。",
	"今日は外でピクニック。気持ちのいい一日でした。",
	"最近、料理に凝ってます。今日は新しいレシピに挑戦しました。",
	"新しいサングラスを買いました。おしゃれ！",
	"今日は友達と映画鑑賞。楽しかったです。",
	"最近、読書の時間を増やしています。心が落ち着きます。",
	"新しい靴を買いました。履き心地がいいです。",
	"今日は家でリラックス。疲れが取れました。",
	"最近、音楽を作ることにハマっています。楽しい！",
	"新しいギターを買いました。練習が楽しみです。",
	"今日は友達とランチ。美味しかったです。",
	"最近、早朝ランニングを始めました。爽快です。",
	"新しいカメラを買いました。写真撮影が楽しみです。",
	"今日はペットと遊びました。癒されました。",
	"新しいゲームをクリアしました。達成感があります。",
	"今日は友達とオンラインゲーム。楽しかったです。",
}

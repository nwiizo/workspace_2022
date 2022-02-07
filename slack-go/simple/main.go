package main

import (
	"github.com/slack-go/slack"
)

func main() {
	/*
		「Slack API: Applications」 にアクセス。
		ワークスペースにログインしていない場合はする。
		Create New App をクリック。
		From scratch をクリック。
		任意のアプリ名を入力、ワークスペースを選択し、 Create App をクリック。
		これでアプリが作成される。
		スコープを設定する
		OAuth & Permissions をクリック。
		ページの中ほどに Scopes セクションがあり、ここでアプリにスコープを設定することができる。
		ボットとして API を実行する際のスコープを設定する Bot Token Scopes と、ユーザーとして API を実行する際のスコープを設定する User Token Scopes がある。
		まず Bot Token Scopes を設定する。
		Add an OAuth Scope をクリック。
		メッセージを投稿するには chat:write スコープが必要なので、追加する。
		User Token Scope も同様の手順で chat:write スコープを追加する。
		アプリをワークスペースにインストールする
		スコープを設定すると同じページの一番上のところで Install to Workspace ボタンが有効になっているので、クリックする。
		権限をリクエストされるので、 許可する をクリック。
		アプリのインストールに成功すると 2 つのトークンが生成される。
		xoxp- から始まる User OAuth Token ( ユーザーとして API を実行するためのトークン )
		xoxb- から始まる Bot User OAuth Token ( ボットとして API を実行するためのトークン )
		アプリをチャンネルに追加する
		ボットとして API を実行するために必要。ユーザーとして API を実行する際には不要。
		Slack でショートカットボタン -> このチャンネルにアプリを追加する の順にクリック。
		今回作成したアプリを探して 追加 をクリック。
	*/
	tkn := "xoxb-**********"
	c := slack.New(tkn)

	_, _, err := c.PostMessage("#*********", slack.MsgOptionText("", true))
	if err != nil {
		panic(err)
	}
}

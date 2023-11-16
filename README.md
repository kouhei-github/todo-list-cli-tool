# TODO List Cli Tool
Todo listが確認できるCLIツールを作成した。
![スクリーンショット 2023-11-17 3 45 55](https://github.com/kouhei-github/todo-list-cli-tool/assets/49782052/33da00a2-c095-44ad-87a6-5261434d7d9e)

---

## 1. 動かすために必要なこと

### 1.1 .envファイルの作成
```shell
cp .env.sample .env
```

・Lineでアカウントを作成して、チャンネルシークレットとアクセストークをセットする<br>
・OpenAIのAPIキーをセットする
```text
LINE_CHANNEL_SECRET=
LINE_CHANNEL_ACCESS_TOKEN=

OPENAI_API_KEY=
```

---

## 2. 起動・停止方法
### 2.1 起動方法
imageの作成
```shell
docker compose build
```

imageからコンテナの起動
```shell
docker compose up -d
```

---

### 2.2 停止方法
```shell
docker compose down
```

---




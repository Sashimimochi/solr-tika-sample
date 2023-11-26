# Tika Server Mode

Apache Tika をサーバモードで起動して使います。

## Preparation

プロファイルで tika-server を指定して起動します。

```bash
$ docker-compose --profile tika-server up -d
```

## Usage

コンテナを起動したらテキスト抽出を行いたいファイルを用意して、エンドポイントに対して投入します。

```bash
$ curl -T your_pdf_file.pdf http://localhost:9998/tika --header "Accept: application/json" > result.json
```

抽出結果は標準出力に JSON 形式で返ってきます。
上記のサンプルコマンドでは、`result.json` で受け取ります。

> [!TIP]
> Docker image だと日本語に対応していないので、自前で server アプリケーションをダウンロードして使用しています。
> cf. https://hub.docker.com/r/apache/tika

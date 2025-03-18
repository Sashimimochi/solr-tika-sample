# Go + Tika

[go-tika](https://pkg.go.dev/github.com/google/go-tika/tika) を使用してGo言語からTikaを呼び出して、抽出結果をMySQLに保存するまでの流れを実装したサンプルコードです。

必要に応じてgo-tikaリポジトリのGitHubリポジトリも参照ください。
https://github.com/google/go-tika/tree/main

## Preparation

事前に以下の2つを準備します。

1. サーバ版のtikaファイル
1. 抽出対象のバイナリファイル

Tikaのダウンロードページからサーバ版のアプリケーションをダウンロードして `app` ディレクトリに配置しておきます。
https://archive.apache.org/dist/tika/

ファイル名は、`tika-server.jar` という名前にしておいてください。

```
$ wget https://archive.apache.org/dist/tika/tika-server-1.21.jar -O go/app/tika-server.jar
```

> [!TIP]
> go-tika で使用できる Tika のバージョンは 1.19~1.21 のみです。
>
> ```go
> fmt.Println("Available Tika Version:", tika.Versions)
> // Available Tika Version: [1.19 1.20 1.21]
> ```
>
> 本サンプルコードでは、バージョン 1.21 を想定しています。
> 別のバージョンを使用する場合は、
>
> ```go
> err := tika.DownloadServer(context.Background(), tika.Version121, "tika-server.jar")
> ```
>
> の箇所を書き換えてください。

抽出対象のファイルはお好みのものをご用意ください。
抽出対象のファイルも `app` ディレクトリ配下に配置しておきます。
基本的には、PDFファイルを想定していますが、Wordファイルなども対応しています。
本サンプルコードでは、`book.pdf` というファイル名を想定しています。
別のファイル名を使用する場合は、

```go
filepath := "book.pdf"
```

を適当なものに書き換えてください。

ファイルが準備できたら、プロファイルでgoを指定してコンテナを起動します。

```bash
$ docker-compose --profile go up -d
```

## Usage

`main.go` ファイルを実行すると、抽出結果がMySQL上のテーブルに保存され、保存結果が標準出力に表示されます。

```bash
$ docker-compose exec go-app go run main.go
テーブルが作成されました
データが挿入されました
挿入されたデータ:
ID: 1, Title: タイトル, Body: 本文, CreatedDate: 2022-09-05 03:53:01 +0000 UTC, ModifiedDate: 2022-09-05 03:53:01 +0000 UTC, Pages: 24, CharLength: 9322
```

# Python + Tika

[tika-python](https://github.com/chrismattmann/tika-python) を使用して、Python から Tika を呼び出して、抽出結果を MySQL に保存するまでの流れを実装したサンプルコードです。

## Preparation

事前に以下を準備します。

1. 抽出対象のバイナリファイル

抽出対象のファイルはお好みのものをご用意ください。
基本的には、PDF ファイルを想定していますが、Word ファイルなども対応しています。
本サンプルコードでは、`app` ディレクトリ配下に `book.pdf` というファイル名が配置されている状態を想定しています。
別のファイル名を使用する場合は、

```python
FILEPATH = 'book.pdf'
```

を適当なものに書き換えてください。

ファイルが準備できたら、プロファイルで python を指定してコンテナを起動します。

```bash
$ docker-compose --profile python up -d
```

## Usage

`app.py` ファイルを実行すると、抽出結果が MySQL 上のテーブルに保存され、保存結果が標準出力に表示されます。

```bash
$ docker-compose exec python-app python app.py
テーブルが作成されました
データが挿入されました
挿入されたデータ:
(1, 'タイトル', '本文', datetime.datetime(2022, 9, 5, 3, 53, 1), datetime.datetime(2022, 9, 5, 3, 53, 1), 24, 9322)
```

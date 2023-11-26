# Solr Tika Sample

本リポジトリは、[「今日から始める Solr ベクトル検索～ Supplemental Book ～」](https://techbookfest.org/product/n3mPcAsU2Bymu1s0x8uj8t?productVariantID=2zhs8fRjBaAAFH0d1WtBJu)に掲載されている Apache Solr と Apache Tika の連携および、Tika の使用方法に関するサンプルコードです。

## Usage

立ち上げたいサービスを指定して立ち上げます。

```bash
$ docker-compose --profile [profile-name] up -d
```

例えば、Solr+Zookeeper のコンテナセットを立ち上げたければ profile-name で `solr` を指定します。

```bash
$ docker-compose --profile solr up -d
```

各サービスがどのプロファイルに属しているかは、`docker-compose.yml` を参照してください。

起動後の使用方法は各ディレクトリの README を参照ください。

停止時はプロファイル名の指定は不要です。

```bash
$ docker-compose down
```

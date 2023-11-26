# Solr + Zookeeper

Solr および Zookeeper を起動して、直接 Tika 経由で PDF ファイルのインデックスを行います。

## Preparation

プロファイルで solr を指定して起動します。

```bash
$ docker-compose --profile solr up -d
```

## Usage

### サンプルの configset を使用する

まずは、サンプルの configset を使用してコレクションを作成します。

```bash
$ curl "http://localhost:8983/solr/admin/collections?action=CREATE&name=gettingstarted&collection.configName=_default&numShards=1&replicationFactor=1&maxShardsPerNode=1"
```

続いて、コレクションに対して ExtractingRequestHandler の設定を行います。

```bash
$ curl -X POST -H 'Content-type:application/json' -d '{
    "add-requesthandler": {
    "name": "/update/extract",
    "class": "solr.extraction.ExtractingRequestHandler",
      "defaults":{ "lowernames": "true", "captureAttr":"true"}
    }
  }' "http://localhost:8983/solr/gettingstarted/config"
```

最後にインデックスさせたいバイナリファイルを指定して投入します。
ここでは、`data` ディレクトリ配下にある `book.pdf` を投入する想定です。
ファイルはご自身でご用意ください。

```bash
$ curl "http://localhost:8983/solr/gettingstarted/update/extract?literal.id=doc1&commit=true" -F "myfile=@data/book.pdf"
```

投入出来たら通常通り検索ができます。

```bash
$ curl http://localhost:8983/solr/gettingstarted/select?indent=true&q=*:*&wt=json
```

### 自前の configset を使う

あらかじめ用意した configset を使用してコレクションを作成します。
ここでは、`conf/book` 配下に用意した configset を使用します。

```bash
# 設定ファイルをアップロード
$ docker-compose exec solr_node1 server/scripts/cloud-scripts/zkcli.sh -zkhost zookeeper1:2181 -cmd upconfig -confdir /opt/solr/server/solr/configsets/book/conf -confname book
# 設定ファイルからコレクションを作成
$ curl "http://localhost:8983/solr/admin/collections?action=CREATE&name=book&collection.configName=book&numShards=1&replicationFactor=1&maxShardsPerNode=1"
# ファイルからメタデータを抽出してインデックス
$ curl "http://localhost:8983/solr/book/update/extract?literal.id=doc1&commit=true" -F "myfile=@solr/data/book.pdf"
```

正常にインデックスができたら検索ができるようになります。

```bash
$ curl http://localhost:8983/solr/book/select?indent=true&q=*:*&wt=json
```

> [!CAUTION]
> Solr Cell の使用は想定外の例外や負荷によって Solr のプロセスに多大なダメージを与える可能性があります。
> そのため、本番環境での使用は極力控えることが推奨されています。
> cf. https://solr.apache.org/guide/solr/latest/indexing-guide/indexing-with-tika.html#solr-cell-performance-implications

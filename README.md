# go-wiki-revel

[go-wiki](https://github.com/snowman-mh/go-wiki) on Revel Framework.

## Description

- [go-wiki](https://github.com/snowman-mh/go-wiki)をRevelフレームワークに乗せて作ってみた。
- [go-wiki](https://github.com/snowman-mh/go-wiki)ではWikiデータをファイル形式で保存したいたが、このリポジトリではデータベースに保存するようにした。
  - データベースにはMySQLを採用
  - OR Mapperには[gorm](https://github.com/jinzhu/gorm)を採用

## How to run

Install libraries.

```
$ cd $GOPATH
$ go get -u github.com/revel/cmd/revel
$ go get -u github.com/jinzhu/gorm
```

Create database.

```
$ mysql -u root
mysql> CREATE DATABASE go_wiki_revel;
mysqr> exit;
```

Clone this repository.

```
$ cd $GOPATH
$ cd src/
$ git clone https://github.com/snowman-mh/go-wiki-revel.git
$ revel run go-wiki-revel
```

Then, visit http://localhost:9000/

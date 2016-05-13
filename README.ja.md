# fanlin

[![Circle CI](https://circleci.com/gh/jobtalk/fanlin/tree/master.svg?style=shield)](https://circleci.com/gh/jobtalk/fanlin/tree/master)

[English](README.md) | 日本語

## 概要

## 環境
* 確認済み
    * go1.4
    * go1.5
    * go1.6

## OS X の環境構築
## master pushの悲劇を防ぐために
[ここを参考に設定すること](http://ganmacs.hatenablog.com/entry/2014/06/18/224132)
## 事前に入れるもの

## Linux用にクロスコンパイルする
go1.5の場合  
1.4はやり方が変わります  
```
$ GOOS=linux GOARCH=amd64 go build github.com/jobtalk/fanlin/cmd/fanlin
```

## サーバーに配布するもの
ビルドして作った実行ファイル  
設定ファイル

## API
getパラメータに値を渡して操作する  
`w`画像の横幅  
`h`画像の縦幅  
色を指定しない時`w`と`h`を指定した時は小さい方に合わせて縮尺を変更する  
この時アスペクト比は維持する  
また一方が`0`の時は`0`ではない値を基準に縮尺を変更する  
`w`及び`h`が`0`であるときは元のサイズで表示する  
あまりにも大きいサイズが指定された時は設定ファイルにかかれている上限の大きさで拡大する  
`rgb`で色を指定した場合`w`と`h`で指定された大きさの画像を生成する  
この時画像のアスペクト比が違うときは隙間を指定した色で塗りつぶす
例:  
```
http://localhost:8080/path/to/image/?h=400&w=400&rgb=100,100,100
```

## testing
```
$ go test -cover ./...
```

## 設定項目に関して
だいたいこんな感じでかけます
```
{
    "port": 8080,
    "max_width": 1000,
    "max_height": 1000,
    "404_img_path": "/path/to/404/image",
    "access_log_path": "/path/to/access/log",
    "error_log_path": "/path/to/error/log",
    "providers": [
        {
            "alias/0" : {
                "type" : "web",
                "src" : "http://aaa.com/bbb"
            }
        },
        {
            "alias/1" : {
                "type" : "web",
                "src" : "https://ccc.com"
            }
        },
        {
            "alias/3" : {
                "type" : "s3",
                "src" : "s3://bucket/path",
                "region" : "ap-northeast-1"
            }
        }
    ]
}
```
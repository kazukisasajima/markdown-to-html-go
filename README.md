# markdown-to-html-go

## 概要

マークダウンを HTML に変換する goスクリプトです。このプロジェクトはコンピュータサイエンス学習サービス[Recursion](https://recursion.example.com)の課題で取り組みました。

## インストール

- スクリプトを実行するにはgomarkdown/markdownライブラリをインストールをインストールする必要があります。

```sh
go get github.com/gomarkdown/markdown
```

## 実行方法

- 以下のコマンドを使用して、マークダウンファイルを HTML ファイルに変換できます。

```sh
go run file_converter.go markdown sample.md output.html
```

## イメージ

例えば、以下のようなマークダウンファイルを、

<img width="600" alt="スクリーンショット 2024-11-24 144321" src="https://github.com/user-attachments/assets/d42c49fb-0992-41bb-aed1-54fcb7509ed8">


以下のようなHTMLファイルに変換します。
<img width="549" alt="{24A45D24-5BD1-4AD5-ADB2-17C09EDC2BD9}" src="https://github.com/user-attachments/assets/2e072f94-75b2-45e2-aeed-df72562aa5f9">


変換したHTMLをブラウザで表示すると、
<img width="510" alt="{CD418CD0-C57D-4584-BB2C-1DEAB5C2D283}" src="https://github.com/user-attachments/assets/22f11132-b994-4a65-b4f6-1006db0ef4da">

のようになります。

# ImageScraper
googleのHPからイメージ画像を自動で取得する

# DEMO
Googleのホームページから自動で画像を取得
さらに自分のダウンロードフォルダにそれを保存
ファイル名を`20XX/XX/XX.jpg`という様に取得した日付に変更

# Features
Googleのホームページから自動で画像を取得
さらに自分のダウンロードフォルダにそれを保存
ファイル名を`20XX/XX/XX.jpg`という様に取得した日付に変更

# Requirement
最新のもので動きます 
* go
* goquery
 
# Installation
goをインストール
```bash
go get 	github.com/PuerkitoBio/goquery
git clone 
```
 
# Usage
```bash
git clone https://github.com/shibuya365/imgscanner
cd examples
go mod init
```
`writeDir := ""`をダウンロードフォルダに変更
```go
writeDir := "/Users/your_username/Downloads/"
```

## 好きなURLから画像を取得
```bash
go run . https://favoriteurl
```
## 省略するとGoogleHPから取得します
```bash
go run .
```
# Note
`writeDir := ""`を設定しなかった場合、
画像は起動したプログラムがあるフォルダに保存されます
 
# Author
* shibuya365
* shibuya365days@gmail.com
 
# License
 
"imgscanner" is under [MIT license](https://en.wikipedia.org/wiki/MIT_License).
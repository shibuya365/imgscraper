package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// 書き込みファイルの初期値
	// writeDir := "/Users/your_username/Downloads/"
	writeDir := ""

	// 今日の日付
	day := time.Now()
	today := day.Format("2006-01-02")

	// 取得URLの初期値を`https://www.google.com`に
	url := "https://www.google.com"
	if len(os.Args) != 1 {
		// 読み込むサイトをURLを取得
		url = os.Args[1]
	}

	// HTMLの読み込み
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()

	// 上手くGETできなかったら
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s\n", res.StatusCode, res.Status)
	}

	// タイトルの部分の抜き出し
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Println(err)
	}
	section := doc.Find("img")

	// 複数の画像を取得
	section.Each(func(i int, line *goquery.Selection) {
		// src属性を取得
		src, _ := line.Attr("src")
		// 拡張子でスライスする
		slice := strings.Split(src, ".")
		// 保存するファイルの名前を`日付+画像番号+拡張子`とする
		filename := today + "_" + strconv.Itoa(i) + "." + slice[len(slice)-1]
		// 最初の画像の場合は画像番号をつけない
		if i == 0 {
			filename = today + "." + slice[len(slice)-1]
		} else {
		}

		// 画像を保存
		if err := DownloadFile(writeDir+filename, url+src); err != nil {
			panic(err)
		}

	})
}

func DownloadFile(filepath string, url string) error {

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

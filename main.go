package main

import (
	"github.com/sclevine/agouti"
	"strings"
	"time"
	"log"
)

func main() {
	/*
	options := agouti.ChromeOptions(
		"args", []string{
			"--headless",
			"-disable-gpu",                // 暫定的に必要とのこと
			"--ignore-certificate-errors", // 認証関係のエラー無視
			"--no-sandbox",
			"--disable-xss-auditor",
		})

	driver := agouti.ChromeDriver(options)
	*/

	log.Println("initialize start")

	options := []agouti.Option{
		agouti.Browser("firefox"),
	}

	capabilities := agouti.Capabilities{
		"moz:firefoxOptions": map[string]interface{}{
			"args": []string{
				"--headless",
			},
			"prefs": map[string]interface{}{
				"network.proxy.type":	0,
				"network.proxy.socks":	"127.0.0.1",
				"network.proxy.socks_port": 8080,
			},
		},
	}

	driver := agouti.GeckoDriver(append(options, agouti.Desired(capabilities))...)

	if err := driver.Start(); err != nil {
		log.Fatalf("Start error: %s\n", err)
	}

	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("NewPage error: %s\n", err)
	}

	log.Println("page.Navigate start")
	url := "https://xxx.atlassian.net/projects")
	page.Navigate(url)
	if err != nil {
		log.Fatalf("Navigate error: %s\n", err)
	}

	// ページ表示までちょっと待つ
	sec := 1
	time.Sleep(time.Duration(sec) * time.Second)

	t, err := page.Title()
	if err != nil {
		log.Printf("Get Title error: %s\n", err)
	}
	log.Println("Title: " + t)

	html, err := page.HTML()
	if err != nil {
		log.Fatalf("Get HTML error: %s\n", err)
	}
	// log.Println("HTML: " + html)

	checkStr := "プロジェクトまたは課題にアクセスできません"
	n := strings.Index(html, checkStr)
	log.Printf("%v matched n: %d\n", checkStr, n)
	if n < 0 {
		errStr := "名前"
		en := strings.Index(html, errStr)
		log.Printf("%v matched n: %d\n", errStr, en)
		if en >= 0 {
			// エラーチェックに引っかかった
			log.Fatalf("設定が誤っている可能性があります。ただしに見直してください。")
		} else {
			// エラーチェックに引っかからないが、想定外の何かが発生した
			log.Fatalf("想定外の状況が発生していますが、設定が誤っているとは限りません。状況を確認してください。")
		}
	} else {
		log.Println("問題ありません。")
	}
}

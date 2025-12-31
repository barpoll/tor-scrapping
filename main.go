package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	content, err := ioutil.ReadFile("targets.yaml")
	if err != nil {
		log.Fatal("[ERR] targets.yaml bulunamadi")
	}
	urls := strings.Split(string(content), "\n")

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.ProxyServer("socks5://127.0.0.1:9150"),
	)

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	fmt.Println("[INFO] Tarama baslatiliyor...")

	for _, url := range urls {
		url = strings.TrimSpace(url)
		if url == "" || strings.HasPrefix(url, "#") {
			continue
		}

		ctx, cancel := chromedp.NewContext(allocCtx)

		var buf []byte
		var htmlContent string
		fmt.Printf("[INFO] Scanning: %s -> ", url)

		err := chromedp.Run(ctx,
			chromedp.Navigate(url),
			chromedp.Sleep(3*time.Second),
			chromedp.FullScreenshot(&buf, 90),
			// HTML'i yakalamak için JavaScript Evaluate kullanalım
			chromedp.Evaluate(`document.documentElement.innerHTML`, &htmlContent),
		)

		if err != nil {
			fmt.Printf("FAILED (%v)\n", err)
			cancel()
			continue
		}

		timestamp := time.Now().UnixNano()

		// PNG Kaydı
		ssName := fmt.Sprintf("ss_%d.png", timestamp)
		_ = ioutil.WriteFile(ssName, buf, 0644)

		// HTML Kaydı (Eğer string doluysa)
		if htmlContent != "" {
			htmlName := fmt.Sprintf("data_%d.html", timestamp)
			_ = ioutil.WriteFile(htmlName, []byte(htmlContent), 0644)
			fmt.Println("SUCCESS (Saved SS and HTML)")
		} else {
			fmt.Println("SUCCESS (SS Saved, HTML Empty)")
		}

		cancel()
	}
}

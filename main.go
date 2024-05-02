package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type Video struct {
	URL     string
	Title   string
	Quality string
	Format  string
	Sound   bool
}

func getDownloadURL(videoID string) ([]Video, error) {
	var videos []Video

	c := colly.NewCollector()

	c.OnHTML("a.downloadBtn", func(e *colly.HTMLElement) {
		url := e.Attr("href")
		title := e.Attr("download")
		quality := e.Text
		format := "mp4"
		sound := true

		switch quality {
		case "720p":
			format = "mp4"
		case "360p":
			format = "mp4"
		case "144p":
			format = "3gp"
		case "1080p50":
			sound = false
		}

		videos = append(videos, Video{
			URL:     url,
			Title:   title,
			Quality: quality,
			Format:  format,
			Sound:   sound,
		})
	})

	err := c.Visit(fmt.Sprintf("https://10downloader.com/download?v=%s&lang=en&type=video", videoID))
	if err != nil {
		return nil, err
	}

	return videos, nil
}

func main() {
	var linkYete string
	fmt.Print("Please insert your YouTube link: ")
	_, err := fmt.Scanln(&linkYete)
	if err != nil {
		log.Fatal(err)
	}

	urls, err := getDownloadURL(linkYete)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(urls[0].URL)
}

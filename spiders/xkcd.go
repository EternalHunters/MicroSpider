package spiders

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"spider/client"
	"sync"
)

var Mux sync.WaitGroup

func NewXkcdSpider() {
	defer Mux.Done()
	resp, err := client.CommonGet("https://xkcd.com", "firefox", false)
	if err != nil {
		return
	}
	doc := soup.HTMLParse(resp)
	links := doc.Find("div", "id", "comicLinks").FindAll("a")
	for _, link := range links {
		fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
	}
}

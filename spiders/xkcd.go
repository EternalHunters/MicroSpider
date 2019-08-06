package spiders

import (
	"fmt"
	"github.com/anaskhan96/soup"
	"spider/client"
)

// xkcd main function
func SpiderXkcd() {
	for i := 0; i < 3; i++ {
		mux.Add(1)
		go newXkcdSpider("https://xkcd.com")
	}
}

func newXkcdSpider(path string) {
	defer mux.Done()
	resp, err := client.CommonGet(path, false, true, false)
	if err != nil {
		return
	}
	doc := soup.HTMLParse(resp)
	links := doc.Find("div", "id", "comicLinks").FindAll("a")
	for _, link := range links {
		fmt.Println(link.Text(), "| Link :", link.Attrs()["href"])
	}
}

package main

import "spider/spiders"

func main() {
	spiders.Mux.Add(1)
	spiders.NewXkcdSpider()
	spiders.Mux.Wait()
}

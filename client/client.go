package client

import (
	"io/ioutil"
	"log"
	"net/http"
)

func CommonGet(path string, brower string, hasCookie bool) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		log.Printf("New request error: %s", err.Error())
		return "", err
	}
	if brower == "firefox" {
		req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:69.0) Gecko/20100101 Firefox/69.0")
	} else {
		req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36")
	}
	if hasCookie {
		cookie := &http.Cookie{}
		req.AddCookie(cookie)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("http get error: %s", err.Error())
		return "", err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	return string(b), err
}

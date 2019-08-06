package client

import (
	"io/ioutil"
	"log"
	"net/http"
)

func CommonGet(path string, hasCookie bool, randomAgent, isMobile bool) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", path, nil)
	if err != nil {
		log.Printf("New request error: %s", err.Error())
		return "", err
	}
	if randomAgent == true {
		if isMobile {
			req.Header.Add("User-Agent", getRandomMobileAgent())
		} else {
			req.Header.Add("User-Agent", getRandomWebAgent())
		}
	} else {
		if isMobile {
			req.Header.Add("User-Agent", "Mozilla/5.0 (iPad; CPU OS 6_0_1 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A523 Safari/8536.25")
		} else {
			req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/75.0.3770.142 Safari/537.36")
		}

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

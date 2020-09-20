package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello from Docker using WSL2, nin ouan")

	client := &http.Client{}

	data := url.Values{}
	data.Set("username", os.Getenv("VA_USER"))
	data.Set("password", os.Getenv("VA_PASS"))
	data.Set("grant_type", "password")

	req, err := http.NewRequest(http.MethodPost, "https://hal.virginactive.com.sg/token", strings.NewReader(data.Encode()))
	if err != nil {
		log.Panicln(err)

	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req.Header.Add("Host", "hal.virginactive.com.sg")
	req.Header.Add("Connection", "keep-alive")
	req.Header.Add("Content-Length", "50")
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("x-mylocker-language", "en-SG")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.102 Safari/537.36 Edg/85.0.564.51")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Origin", "https://mylocker.virginactive.com.sg")
	req.Header.Add("Sec-Fetch-Site", "same-site")
	req.Header.Add("Sec-Fetch-Mode", "cors")
	req.Header.Add("Sec-Fetch-Dest", "empty")
	req.Header.Add("Referer", "https://mylocker.virginactive.com.sg/")
	req.Header.Add("Accept-Encoding", "gzip, deflate, br")
	req.Header.Add("Accept-Language", "en-US,en;q=0.9")

	res, err := client.Do(req)

	if err != nil {
		log.Panicln(err)
	} else {
		log.Println("Response:", res)
		resBody, _ := ioutil.ReadAll(res.Body)
		log.Println("Response body:", string(resBody))
	}

}

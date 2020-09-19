package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	fmt.Println("Hello from Docker using WSL2, nin ouan")

	res, err := http.PostForm("https://hal.virginactive.com.sg/token",
		url.Values{"username": {os.Getenv("VA_USER")}, "password": {os.Getenv("VA_PASS")}, "grant_type": {"password"}})

	log.Println(res, err)

}

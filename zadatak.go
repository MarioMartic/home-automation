package main

import (
	"time"
	"fmt"
	"net/http"
	"strings"
	"strconv"
	"net/url"
)

func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}

func action(t time.Time) {
	apiUrl := "http://192.168.1.117:3000"
	fmt.Println("URL:>", apiUrl)

	data := url.Values{}
	data.Set("led_4", "1")


	client := &http.Client{}
	r, _ := http.NewRequest("POST", apiUrl, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Authorization", "059b9576-89ea-468e-81fb-564d1331055c")
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, _ := client.Do(r)
	fmt.Println(resp.Status)

}

func keepAlive(t time.Time) {
	apiUrl := "http://192.168.1.117:3000"
	fmt.Println("URL:>", apiUrl)

	data := url.Values{}

	client := &http.Client{}
	r, _ := http.NewRequest("POST", apiUrl, strings.NewReader(data.Encode())) // URL-encoded payload
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	resp, _ := client.Do(r)
	fmt.Println(resp)
	
}

func main() {
	doEvery(60*time.Second, keepAlive)
}

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	sites := [3]string{
		"https://facebook.com",
		"https://google.com",
		"https://instagram.com",
	}

	c := make(chan string)

	for _, site := range sites {
		go checkLink(site, c)
	}

	// for {
	// 	go checkLink(<-c, c)
	// }

	// Equivalent

	for l := range c {
		go func(l string, c chan string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(l, c)
	}
}

func checkLink(link string, c chan string) {
	r, _ := http.Get(link)

	if r.StatusCode == 200 {
		fmt.Println(link, "is up")
		c <- link
	} else {
		fmt.Println(link, "is down")
		c <- link
	}
}

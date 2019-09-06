package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string) // channel: go routine同士を連携させる(stringの交換)

	for _, link := range links {
		go checkLink(link, c)
	}
	
	for l := range c{
		go checkLink(l, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		c <- "Might be down I think"
		fmt.Println(link, "might be down!")
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}

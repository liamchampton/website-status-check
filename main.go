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
		"https://golang.org",
		"http://amazon.com",
		"http://ibm.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	fmt.Println(<-c)

}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- "might be down I think!"
		return
	}

	fmt.Println(link, "is up!")
	c <- "Yep, it's up!"
}

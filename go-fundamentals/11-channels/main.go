package main

import (
	"net/http"
	"time"

	"github.com/fatih/color"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	strCh := make(chan string)

	// Initialization. We call checkLink function with every link we have to check
	for _, link := range links {
		go checkLink(link, strCh)
	}

	// Infinite loop
	// for {
	// 	go checkLink(<-strCh, strCh)
	// }

	// Loop all the values of the channel
	// for l := range strCh {
	// 	go checkLink(l, strCh)
	// }

	// Loop all the values of the channel with a delay
	for l := range strCh {
		go func(link string) { // unnamed function
			time.Sleep(5 * time.Second)
			checkLink(link, strCh)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		color := color.New(color.FgRed, color.Bold)
		color.Println(link, "might be down!")
	} else {
		color := color.New(color.FgGreen)
		color.Println(link, "is up!")
	}
	c <- link
}

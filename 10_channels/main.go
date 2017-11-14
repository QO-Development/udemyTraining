package main

import "net/http"
import "fmt"
import "time"

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		//Keyword go creates a new thread go routine
		go checkLink(link, c)
	}

	//Infinite loop
	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {

	//http.Get seems to be a synchronous network call, run on the main thread
	_, err := http.Get(link)

	if err != nil {
		fmt.Println("There was a problem:", err)
		c <- link
	} else {
		fmt.Println("Link", link, "is good to go.")
		c <- link
	}
}

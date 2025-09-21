package main

import (
	"fmt"
	"io"
	"net/http"
)

type HomePageSize struct {
	URL  string
	Size int
}

func main() {
	urls := []string{
		"http://www.google.com/",
		"http://www.amazon.com/",
		"http://www.apple.com/",
		"http://www.microsoft.com/",
	}

	results := make(chan HomePageSize)

	for _, url := range urls {
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("Error fetching ", url, ":", err)
				panic(err)
			}
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Error reading ", url, ":", err)
				panic(err)
			}
			results <- HomePageSize{url, len(body)}
		}(url)
	}

	var biggest HomePageSize

	for range urls {
		result := <-results
		fmt.Println(result.URL, result.Size)
		if result.Size > biggest.Size {
			biggest = result
		}
	}

	fmt.Println("\nBiggest page: ", biggest.URL, "with size ", biggest.Size)
}

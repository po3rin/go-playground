package ch4

import (
	"fmt"
	"net/http"
)

func P100() {
	checkStatus := func(
		done <-chan interface{},
		urls ...string,
	) <-chan *http.Response {
		responses := make(chan *http.Response)
		go func() {
			defer close(responses)
			for _, url := range urls {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err) // <1>
					continue
				}
				select {
				case <-done:
					return
				case responses <- resp:
				}
			}
		}()
		fmt.Println("ee")
		return responses
	}

	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://www.google.com", "https://badhost"}
	for response := range checkStatus(done, urls...) {
		fmt.Println("pp")
		fmt.Printf("Response: %v\n", response.Status)
	}
}

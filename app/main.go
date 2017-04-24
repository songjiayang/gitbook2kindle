package main

import (
	"sync"

	"github.com/gitbook2kindle/gitbook"
	"github.com/gitbook2kindle/kindle"
)

func main() {
	gClient := gitbook.NewClient()
	books := gClient.ListStarBooks()

	var wg sync.WaitGroup
	wg.Add(len(books))

	queue := make(chan string, 4)

	for _, book := range books {
		queue <- book.ID

		go func(q chan string, wg *sync.WaitGroup) {
			defer wg.Done()
			id := <-q
			kindle.SendBook(id, gClient.DownloadBook(id))
		}(queue, &wg)
	}

	wg.Wait()
}

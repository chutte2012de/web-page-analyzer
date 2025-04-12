package usecase

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/useinsider/go-pkg/insrequester"
)

// https://medium.com/insiderengineering/concurrent-http-requests-in-golang-best-practices-and-techniques-f667e5a19dea
// https://stackoverflow.com/questions/45337881/go-routinemaking-concurrent-api-requests

type LinkManager struct {
}

type Job struct {
	URL string
}

func (l *LinkManager) CategorizeLinks(url string, links []string) ([]string, []string, error) {

	internalLinks := make([]string, 0, 100)
	externalLinks := make([]string, 0, 100)

	return internalLinks, externalLinks, nil
}

func worker(requester *insrequester.Request, jobs <-chan Job, results chan<- *http.Response, wg *sync.WaitGroup) {
	for job := range jobs {
		res, _ := requester.Get(insrequester.RequestEntity{Endpoint: job.URL})
		results <- res
		wg.Done()
	}
}

func (l *LinkManager) ValidateLinks(preffix string, links []string) ([]string, error) {

	internalLinks := make([]string, 0, 100)

	requester := insrequester.NewRequester().Load()

	urls := []string{"http://example.com", "https://www.freecodecamp.org",
		"https://www.behance.net", "https://codepen.io", "http://example.org",
		"http://example.net", "http://exampleapple.net/"}
	numWorkers := 2 // Define the number of workers in the pool

	jobs := make(chan Job, len(urls))
	results := make(chan *http.Response, len(urls))
	var wg sync.WaitGroup

	// Start workers
	for w := 0; w < numWorkers; w++ {
		go worker(requester, jobs, results, &wg)
	}

	// Sending jobs to the worker pool
	wg.Add(len(urls))
	for _, url := range urls {
		jobs <- Job{URL: url}
	}
	close(jobs)
	wg.Wait()

	// Collecting results
	for i := 0; i < len(urls); i++ {
		fmt.Println("Result")
		fmt.Println(<-results)
		fmt.Println()
	}

	return internalLinks, nil
}

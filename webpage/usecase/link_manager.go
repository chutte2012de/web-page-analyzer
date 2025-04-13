package usecase

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"sync"

	"github.com/chutte2012de/web-page-analyzer/webpage/model"
	"github.com/useinsider/go-pkg/insrequester"
)

// https://medium.com/insiderengineering/concurrent-http-requests-in-golang-best-practices-and-techniques-f667e5a19dea
// https://stackoverflow.com/questions/45337881/go-routinemaking-concurrent-api-requests

type LinkManager struct {
}

type Job struct {
	URL string
}

func (l *LinkManager) CategorizeLinks(mainUrlStr string, links []string) ([]string, []string, error) {

	internalLinks := make([]string, 0, 100)
	externalLinks := make([]string, 0, 100)

	mainUrl, _ := url.Parse(mainUrlStr)
	log.Println("Main URL Host: ", mainUrl.Host)
	log.Println("Main URL Host Name: ", mainUrl.Hostname())
	currentHostName := strings.TrimPrefix(mainUrl.Hostname(), "www.")
	log.Println("currentHostName: ", currentHostName)

	for _, linkUrlStr := range links {
		log.Println("linkUrl: ", linkUrlStr)

	}

	return internalLinks, externalLinks, nil
}

func worker(requester *insrequester.Request, jobs <-chan Job, results chan<- *model.Link, wg *sync.WaitGroup) {
	for job := range jobs {
		result := model.Link{
			Url: job.URL,
		}
		res, err := requester.Get(insrequester.RequestEntity{Endpoint: job.URL})
		//defer res.Body.Close()

		if err != nil {
			fmt.Println("ERROR: Failed to reach:", job.URL, " with error", err)
			result.Status = string(err.Error())
		} else {
			result.Status = res.Status
			//result.Body = string(res.Body)
		}
		// results <- res
		// results <- &model.Link{
		// 	Url:    job.URL,
		// 	Status: string(res.),
		// }
		results <- &result

		wg.Done()
	}
}

func (l *LinkManager) ValidateLinks(preffix string, links []string) ([]model.Link, error) {
	requester := insrequester.NewRequester().Load()

	numWorkers := 2 // Define the number of workers in the pool

	jobs := make(chan Job, len(links))
	results := make(chan *model.Link, len(links))
	var wg sync.WaitGroup

	// Start workers
	for w := 0; w < numWorkers; w++ {
		go worker(requester, jobs, results, &wg)
	}

	// Sending jobs to the worker pool
	wg.Add(len(links))
	for _, url := range links {
		jobs <- Job{URL: url}
	}
	close(jobs)
	wg.Wait()

	output := make([]model.Link, 0, len(links))

	// Collecting results
	for i := 0; i < len(links); i++ {
		fmt.Println("Result")
		r := <-results
		fmt.Println(r)
		// fmt.Println(<-results)
		output = append(output, *r)
		fmt.Println()
	}
	close(results)
	wg.Wait()
	//wg.Done()

	return output, nil
}

package usecase

import (
	"log/slog"
	"net/url"
	"strings"
	"sync"

	"github.com/chutte2012de/web-page-analyzer/webpage/model"
	"github.com/useinsider/go-pkg/insrequester"
)

// https://medium.com/insiderengineering/concurrent-http-requests-in-golang-best-practices-and-techniques-f667e5a19dea
// https://stackoverflow.com/questions/45337881/go-routinemaking-concurrent-api-requests

type (
	ILinkManager interface {
		GetLinksInfo(inputUrlStr string, linksInPage []string) (model.LinksInfo, error)
		GetHostAndWebsiteBaseNames(inputUrlStr string) (string, string, error)
		CategorizeLinks(websiteBaseName string, links []string) ([]string, []string, []string, error)
		GetLinksSummaryStat(links []model.Link) (model.SummaryStat, error)
		ValidateLinks(preffix string, links []string) ([]model.Link, error)
	}

	linkManager struct {
	}
)

func NewLinkManager() ILinkManager {
	return &linkManager{}
}

type Job struct {
	URL string
}

func (l *linkManager) GetLinksInfo(inputUrlStr string, linksInPage []string) (model.LinksInfo, error) {
	slog.Info("GetLinksInfo", "inputUrlStr", inputUrlStr)
	inputUrlHostName, inputUrlWebsiteBaseName, _ := l.GetHostAndWebsiteBaseNames(inputUrlStr)
	internalLinksWithoutPreffix, internalLinksWithPreffix, externalLinks, _ := l.CategorizeLinks(inputUrlWebsiteBaseName, linksInPage)

	internalWithoutPreffixLinksInfo, _ := l.ValidateLinks(inputUrlHostName, internalLinksWithoutPreffix)
	internalWithoutPreffixLinksSummaryStat, _ := l.GetLinksSummaryStat(internalWithoutPreffixLinksInfo)

	internalWithPreffixLinksInfo, _ := l.ValidateLinks("", internalLinksWithPreffix)
	internalWithPreffixLinksSummaryStat, _ := l.GetLinksSummaryStat(internalWithPreffixLinksInfo)

	internalWithoutPreffixLinksInfo = append(internalWithoutPreffixLinksInfo, internalWithPreffixLinksInfo...)
	internalWithoutPreffixLinksSummaryStat.TotalCount += internalWithPreffixLinksSummaryStat.TotalCount
	internalWithoutPreffixLinksSummaryStat.ReachableCount += internalWithPreffixLinksSummaryStat.ReachableCount
	internalWithoutPreffixLinksSummaryStat.UnreachableCount += internalWithPreffixLinksSummaryStat.UnreachableCount

	externalLinksInfo, _ := l.ValidateLinks("", externalLinks)
	externalLinksSummaryStat, _ := l.GetLinksSummaryStat(externalLinksInfo)

	linksInfo := model.LinksInfo{
		Summary: model.LinksSummary{
			Internal: internalWithoutPreffixLinksSummaryStat,
			External: externalLinksSummaryStat,
		},
		Detail: model.LinksDetail{
			Internal: internalWithoutPreffixLinksInfo,
			External: externalLinksInfo,
		},
	}

	return linksInfo, nil
}

func (l *linkManager) GetHostAndWebsiteBaseNames(inputUrlStr string) (string, string, error) {
	slog.Info("GetHostAndWebsiteBaseNames", "inputUrlStr", inputUrlStr)
	inputUrl, err := url.Parse(inputUrlStr)
	if err != nil {
		slog.Error("Input Url String Parsing failed", "inputUrlStr", inputUrlStr)
		return "", "", err
	}
	websiteBaseName := strings.TrimPrefix(inputUrl.Hostname(), "www.")
	slog.Info("GetHostAndWebsiteBaseNames", "websiteBaseName", websiteBaseName)

	inputUrl.Path = ""
	inputUrl.RawQuery = ""
	inputUrl.Fragment = ""

	return inputUrl.String(), websiteBaseName, nil
}

func (l *linkManager) CategorizeLinks(websiteBaseName string, links []string) ([]string, []string, []string, error) {
	slog.Info("CategorizeLinks", "websiteBaseName", websiteBaseName)

	internalLinksWithoutPreffix := make([]string, 0, 100)
	internalLinksWithPreffix := make([]string, 0, 100)
	externalLinks := make([]string, 0, 100)

	for _, linkUrlStr := range links {
		slog.Info("Parsing Link URL", "linkUrlStr", linkUrlStr)
		linkUrl, err := url.Parse(linkUrlStr)
		if err != nil {
			slog.Error("Link URL Parsed failed", "linkUrlStr", linkUrlStr)
			continue
		}

		slog.Info("Link URL Parsed", "Hostname", linkUrl.Hostname())
		if linkUrl.Hostname() == "" {
			internalLinksWithoutPreffix = append(internalLinksWithoutPreffix, linkUrlStr)
		} else if strings.Contains(linkUrl.Hostname(), websiteBaseName) {
			internalLinksWithPreffix = append(internalLinksWithPreffix, linkUrlStr)
		} else {
			externalLinks = append(externalLinks, linkUrlStr)
		}
	}

	return internalLinksWithoutPreffix, internalLinksWithPreffix, externalLinks, nil
}

func (l *linkManager) GetLinksSummaryStat(links []model.Link) (model.SummaryStat, error) {
	total := 0
	reachable := 0
	unreachable := 0

	for _, link := range links {
		total++
		if link.Reachable == "YES" {
			reachable++
		} else if link.Reachable == "NO" {
			unreachable++
		}
	}

	summary := model.SummaryStat{
		TotalCount:       uint16(total),
		ReachableCount:   uint16(reachable),
		UnreachableCount: uint16(unreachable),
	}
	return summary, nil
}

func worker(requester *insrequester.Request, jobs <-chan Job, results chan<- *model.Link, wg *sync.WaitGroup) {
	for job := range jobs {
		result := model.Link{
			Url: job.URL,
		}
		res, err := requester.Get(insrequester.RequestEntity{Endpoint: job.URL})

		if err != nil {
			slog.Info("Failed to reach by Link Worker", "url", job.URL, "error", err)
			result.Status = string(err.Error())
			result.Reachable = "NO"
		} else {
			result.Status = res.Status
			result.Reachable = "YES"
		}
		results <- &result

		wg.Done()
	}
}

func (l *linkManager) ValidateLinks(preffix string, links []string) ([]model.Link, error) {
	requester := insrequester.NewRequester().Load()

	numWorkers := 5 // Define the number of workers in the pool

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
		jobs <- Job{URL: preffix + url}
	}
	close(jobs)
	wg.Wait()

	output := make([]model.Link, 0, len(links))

	// Collecting results
	for i := 0; i < len(links); i++ {
		//fmt.Println("Result")
		r := <-results
		// fmt.Println(r)
		// fmt.Println(<-results)
		output = append(output, *r)
		//fmt.Println()
	}
	close(results)
	wg.Wait()
	//wg.Done()

	return output, nil
}

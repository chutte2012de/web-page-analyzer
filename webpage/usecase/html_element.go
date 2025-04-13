package usecase

import (
	"fmt"
	"io"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/chutte2012de/web-page-analyzer/webpage/model"
	"golang.org/x/net/html"
)

type HtmlElement struct {
	LinkMan *LinkManager
}

func (h *HtmlElement) ExtractFromUrl(url string) (model.HtmlStat, error) {

	slog.Info("Start ExtractFromUrl", "url", url)
	resp, err := http.Get(url)

	now := time.Now().UTC()

	if err != nil {
		slog.Error("Failed http Get", "url", url, "error", err)
		return model.HtmlStat{
			Id:        rand.Uint64(),
			Url:       url,
			CreatedAt: &now,
		}, err
	}

	b := resp.Body
	defer b.Close() // close Body when the function completes

	htmlStat, linksInPage, err := h.ExtractFromIoReader(b)

	if err != nil {
		slog.Error("Failed to Extract Html Elements", "url", url, "error", err)
		return model.HtmlStat{
			Id:        rand.Uint64(),
			Url:       url,
			CreatedAt: &now,
		}, err
	}

	// TODO:
	inputUrlHostName, inputUrlWebsiteBaseName, _ := h.LinkMan.GetHostAndWebsiteBaseNames(url)
	internalLinksWithoutPreffix, _, externalLinks, _ := h.LinkMan.CategorizeLinks(inputUrlWebsiteBaseName, linksInPage)
	// internalLinksWithoutPreffix, internalLinksWithPreffix, externalLinks, _ := h.LinkMan.CategorizeLinks(inputUrlWebsiteBaseName, linksInPage)

	internalWithoutPreffixLinksInfo, _ := h.LinkMan.ValidateLinks(inputUrlHostName, internalLinksWithoutPreffix)
	internalWithoutPreffixLinksSummaryStat, _ := h.LinkMan.GetLinksSummaryStat(internalWithoutPreffixLinksInfo)

	// linksInfo, _ := h.LinkMan.ValidateLinks("", linksInPage)
	// linksSummaryStat, _ := h.LinkMan.GetLinksSummaryStat(linksInfo)

	externalLinksInfo, _ := h.LinkMan.ValidateLinks("", externalLinks)
	externalLinksSummaryStat, _ := h.LinkMan.GetLinksSummaryStat(externalLinksInfo)

	htmlStat.LinksInfo = model.LinksInfo{
		Summary: model.LinksSummary{
			Internal: internalWithoutPreffixLinksSummaryStat,
			External: externalLinksSummaryStat,
		},
		Detail: model.LinksDetail{
			Internal: internalWithoutPreffixLinksInfo,
			External: externalLinksInfo,
		},
	}

	htmlStat.Url = url
	slog.Info("Complete ExtractFromUrl", "url", url)
	return htmlStat, nil
}

func (h *HtmlElement) ExtractFromIoReader(r io.Reader) (model.HtmlStat, []string, error) {
	now := time.Now().UTC()
	htmlStat := model.HtmlStat{
		Id: rand.Uint64(),
		Headers: model.Headers{
			H1: 0,
			H2: 0,
			H3: 0,
			H4: 0,
			H5: 0,
			H6: 0,
		},
		CreatedAt: &now,
	}

	tokenizer := html.NewTokenizer(r)

	links := make([]string, 0, 100)

	currentTokenTag := ""

	for {
		tokenType := tokenizer.Next()
		token := tokenizer.Token()
		if tokenType == html.ErrorToken {
			fmt.Printf("links: %v\n", links)
			fmt.Println("links len: ", len(links))
			fmt.Println("links cap: ", cap(links))
			if tokenizer.Err() == io.EOF {
				fmt.Printf("End of File")
				return htmlStat, links, nil
			}
			fmt.Printf("Error: %v", tokenizer.Err())
			return htmlStat, links, tokenizer.Err()
		}

		// fmt.Printf("Token: %v\n", html.UnescapeString(token.String()))

		if tokenType == html.DoctypeToken {
			htmlStat.Version = token.Data
			currentTokenTag = ""
			continue
		}

		switch tokenType {
		case html.StartTagToken:
			currentTokenTag = token.Data
		case html.EndTagToken:
			currentTokenTag = ""
		case html.SelfClosingTagToken:
			currentTokenTag = token.Data
		case html.TextToken:
			// do nothing, because it probably between StartTagToken and EndTagToken
		default: //This will also include contents of <script>, <style> tags content
			currentTokenTag = ""
		}

		// fmt.Printf("currentTokenTag: %v\n", currentTokenTag)

		if html.TextToken == tokenType && currentTokenTag == "title" {
			htmlStat.Title = token.Data
		}

		if html.StartTagToken == tokenType {
			switch token.Data {
			case "h1":
				htmlStat.Headers.H1++
			case "h2":
				htmlStat.Headers.H2++
			case "h3":
				htmlStat.Headers.H3++
			case "h4":
				htmlStat.Headers.H4++
			case "h5":
				htmlStat.Headers.H5++
			case "h6":
				htmlStat.Headers.H6++
			}
		}

		if currentTokenTag == "a" || currentTokenTag == "link" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					//links = append(links, attr.Val)
					// fmt.Println("a Link: [", attr.Val, "]")
					links = append(links, attr.Val)
				}

			}
		}

		if html.SelfClosingTagToken == tokenType {
			// fmt.Printf("Token: %v\n", html.UnescapeString(token.String()))
			// fmt.Printf("currentTokenTag: %v\n", currentTokenTag)
			if currentTokenTag == "img" {
				for _, attr := range token.Attr {
					if attr.Key == "src" {
						//links = append(links, attr.Val)
						// fmt.Println("a Link: [", attr.Val, "]")
						links = append(links, attr.Val)
					}

				}
			}
			currentTokenTag = ""
		}

		// 	// Check if the token is an <a> tag
		// 	isAnchor := t.Data == "a"
		// 	if !isAnchor {
		// 		continue
		// 	}

		// 	// Extract the href value, if there is one
		// 	// ok, url := getHref(t)
		// 	// if !ok {
		// 	// 	continue
		// 	// }

		// }
	}

	return htmlStat, links, nil
}

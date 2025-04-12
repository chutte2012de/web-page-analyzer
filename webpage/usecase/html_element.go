package usecase

import (
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/chutte2012de/web-page-analyzer/webpage/model"
	"golang.org/x/net/html"
)

type HtmlElement struct {
}

func (h *HtmlElement) ExtractFromUrl(url string) (model.HtmlStat, error) {

	fmt.Println("Begin Extract url: [", url, "]")
	resp, err := http.Get(url)

	now := time.Now().UTC()

	if err != nil {
		fmt.Println("ERROR: Failed to crawl:", url)
		return model.HtmlStat{
			Id:        rand.Uint64(),
			Url:       url,
			CreatedAt: &now,
		}, err
	}

	b := resp.Body
	defer b.Close() // close Body when the function completes

	htmlStat, err := h.ExtractFromIoReader(b)

	if err != nil {
		fmt.Println("ERROR: Failed to ExtractFromIoReader of Url:", url)
		return model.HtmlStat{
			Id:        rand.Uint64(),
			Url:       url,
			CreatedAt: &now,
		}, err
	}

	fmt.Println("End Extract url: [", url, "]")

	htmlStat.Url = url

	return htmlStat, nil
}

func (h *HtmlElement) ExtractFromIoReader(r io.Reader) (model.HtmlStat, error) {
	// fmt.Println("Begin url: [", url, "]")
	// resp, err := http.Get(url)

	// if err != nil {
	// 	fmt.Println("ERROR: Failed to crawl:", url)
	// 	return
	// }

	// b := resp.Body
	// defer b.Close() // close Body when the function completes

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

	// elementsMap := make(map[string]uint16)

	// links := []string{}
	// links := make([]string, 100)

	// titleStr := ""

	// docScanningCompleted := false

	currentTokenTag := ""

	for {
		tokenType := tokenizer.Next()
		token := tokenizer.Token()
		if tokenType == html.ErrorToken {
			if tokenizer.Err() == io.EOF {
				fmt.Printf("End of File")
				return htmlStat, nil
			}
			fmt.Printf("Error: %v", tokenizer.Err())
			return htmlStat, tokenizer.Err()
		}

		fmt.Printf("Token: %v\n", html.UnescapeString(token.String()))

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

		fmt.Printf("currentTokenTag: %v\n", currentTokenTag)

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
				// case "script":
				// 	fmt.Printf("Script Token: %v\n", html.UnescapeString(token.String()))
				// case "style":
				// 	fmt.Printf("Style Token: %v\n", html.UnescapeString(token.String()))
				// default: //This will also include contents of <script>, <style> tags content
				// 	fmt.Printf("Others: %v\n", html.UnescapeString(token.String()))
			}
		}

		if html.SelfClosingTagToken == tokenType {
			currentTokenTag = ""
		}

		// continue

		// if true == docScanningCompleted {
		// 	break
		// }
		// tt := z.Next()

		// fmt.Println("tt: [", tt.String(), "]")

		// switch {

		// case tt == html.ErrorToken:
		// 	fmt.Println("ErrorToken TOKEN REACHED.")
		// 	// fmt.Println("elementsMap: ", elementsMap)
		// 	// fmt.Println("links: ", links)
		// 	fmt.Println("LEN No. of links: ", len(links))
		// 	fmt.Println("CAP No. of links: ", cap(links))
		// 	fmt.Println("titleStr: ", titleStr)
		// 	docScanningCompleted = true
		// 	// End of the document, we're done
		// 	// return
		// 	//break
		// case tt == html.DoctypeToken:
		// 	t := z.Token()
		// 	fmt.Println("DoctypeToken Type: [", t.Type, "], n Data: [", t.Data, "]")

		// case tt == html.SelfClosingTagToken:
		// 	t := z.Token()
		// 	fmt.Println("SelfClosingTagToken BEGIN Type: [", t.Type, "], n Data: [", t.Data, "]")

		// 	for _, attr := range t.Attr {
		// 		fmt.Println("Attr Key: [", attr.Key, "], n Val: [", attr.Val, "]")
		// 	}
		// 	fmt.Println("SelfClosingTagToken END Type: [", t.Type, "], n Data: [", t.Data, "]")

		// case tt == html.TextToken:
		// 	// t := z.Token()
		// 	// fmt.Println("TextToken Type: [", t.Type, "], n Data: [", t.Data, "]")
		// case tt == html.StartTagToken:
		// 	t := z.Token()

		// 	//fmt.Println("t Type: [", t.Type, "], n Data: [", t.Data, "]")

		// 	v, ok := elementsMap[t.Data]
		// 	if ok {
		// 		elementsMap[t.Data] = v + 1
		// 	} else {
		// 		elementsMap[t.Data] = 1
		// 	}

		// 	if "title" == t.Data {
		// 		fmt.Println("title contents: [", t.String(), "]")

		// 		// for _, attr := range t.Attr {
		// 		// 	fmt.Println("title contents: Key[", attr.Key, "],  Val[", attr.Val, "]")
		// 		// }
		// 		if tt = z.Next(); tt == html.TextToken {
		// 			titleStr = z.Token().Data
		// 			fmt.Println(z.Token().Data)
		// 		}
		// 	}

		// 	if "a" == t.Data {
		// 		for _, attr := range t.Attr {
		// 			if attr.Key == "href" {
		// 				//links = append(links, attr.Val)
		// 				// fmt.Println("a Link: [", attr.Val, "]")
		// 				links = append(links, attr.Val)
		// 			}

		// 		}
		// 	}

		// 	if "h1" == t.Data {
		// 		html_stat.Headers.H1++
		// 	}

		// 	if "h2" == t.Data {
		// 		html_stat.Headers.H2++
		// 	}

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

	// fmt.Println("elementsMap:", elementsMap)

	return htmlStat, nil
}

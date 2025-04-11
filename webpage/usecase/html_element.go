package usecase

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/net/html"
)

type HtmlElement struct {
}

func (h *HtmlElement) ExtractFromUrl(url string) {

	fmt.Println("Begin Extract url: [", url, "]")
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("ERROR: Failed to crawl:", url)
		return
	}

	b := resp.Body
	defer b.Close() // close Body when the function completes

	h.ExtractFromIoReader(b)

	fmt.Println("End Extract url: [", url, "]")
}

func (h *HtmlElement) ExtractFromIoReader(r io.Reader) {
	// fmt.Println("Begin url: [", url, "]")
	// resp, err := http.Get(url)

	// if err != nil {
	// 	fmt.Println("ERROR: Failed to crawl:", url)
	// 	return
	// }

	// b := resp.Body
	// defer b.Close() // close Body when the function completes

	z := html.NewTokenizer(r)

	elementsMap := make(map[string]uint16)

	// links := []string{}
	links := make([]string, 100)

	titleStr := ""

	for {
		tt := z.Next()

		//fmt.Println("tt: [", tt.String(), "]")

		switch {

		case tt == html.ErrorToken:
			fmt.Println("ErrorToken TOKEN REACHED.")
			// fmt.Println("elementsMap: ", elementsMap)
			// fmt.Println("links: ", links)
			fmt.Println("LEN No. of links: ", len(links))
			fmt.Println("CAP No. of links: ", cap(links))
			fmt.Println("titleStr: ", titleStr)
			// End of the document, we're done
			return
		case tt == html.DoctypeToken:
			t := z.Token()
			fmt.Println("DoctypeToken Type: [", t.Type, "], n Data: [", t.Data, "]")

		case tt == html.SelfClosingTagToken:
			t := z.Token()
			fmt.Println("SelfClosingTagToken BEGIN Type: [", t.Type, "], n Data: [", t.Data, "]")

			for _, attr := range t.Attr {
				fmt.Println("Attr Key: [", attr.Key, "], n Val: [", attr.Val, "]")
			}
			fmt.Println("SelfClosingTagToken END Type: [", t.Type, "], n Data: [", t.Data, "]")

		case tt == html.TextToken:
			// t := z.Token()
			// fmt.Println("TextToken Type: [", t.Type, "], n Data: [", t.Data, "]")
		case tt == html.StartTagToken:
			t := z.Token()

			//fmt.Println("t Type: [", t.Type, "], n Data: [", t.Data, "]")

			v, ok := elementsMap[t.Data]
			if ok {
				elementsMap[t.Data] = v + 1
			} else {
				elementsMap[t.Data] = 1
			}

			if "title" == t.Data {
				fmt.Println("title contents: [", t.String(), "]")

				// for _, attr := range t.Attr {
				// 	fmt.Println("title contents: Key[", attr.Key, "],  Val[", attr.Val, "]")
				// }
				if tt = z.Next(); tt == html.TextToken {
					titleStr = z.Token().Data
					fmt.Println(z.Token().Data)
				}
			}

			if "a" == t.Data {
				for _, attr := range t.Attr {
					if attr.Key == "href" {
						//links = append(links, attr.Val)
						// fmt.Println("a Link: [", attr.Val, "]")
						links = append(links, attr.Val)
					}

				}
			}

			// Check if the token is an <a> tag
			isAnchor := t.Data == "a"
			if !isAnchor {
				continue
			}

			// Extract the href value, if there is one
			// ok, url := getHref(t)
			// if !ok {
			// 	continue
			// }

		}
	}

	// fmt.Println("elementsMap:", elementsMap)
}

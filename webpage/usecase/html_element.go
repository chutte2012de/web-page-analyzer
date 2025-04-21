package usecase

import (
	"fmt"
	"io"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"strings"
	"time"

	"github.com/chutte2012de/web-page-analyzer/webpage/model"
	"golang.org/x/net/html"
)

type (
	IHtmlElement interface {
		ExtractFromUrl(url string) (model.HtmlStat, error)
		ExtractFromIoReader(r io.Reader) (model.HtmlStat, []string, error)
	}

	htmlElement struct {
		LinkMan ILinkManager
	}
)

func NewHtmlElement(linkManager ILinkManager) IHtmlElement {
	return &htmlElement{
		LinkMan: linkManager,
	}
}

func (h *htmlElement) ExtractFromUrl(url string) (model.HtmlStat, error) {

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

	linksInfo, _ := h.LinkMan.GetLinksInfo(url, linksInPage)
	htmlStat.LinksInfo = linksInfo

	htmlStat.Url = url
	slog.Info("Complete ExtractFromUrl", "url", url)
	return htmlStat, nil
}

func (h *htmlElement) ExtractFromIoReader(r io.Reader) (model.HtmlStat, []string, error) {
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

		// Check if the Form is available in web page
		// this will be used to determine if a login form is available
		if currentTokenTag == "form" {
			fmt.Println("FORM tag: [", html.StartTagToken, "]")
		}
		if html.StartTagToken == tokenType && currentTokenTag == "form" {
			htmlStat.Form = "AVAILABLE"
		}

		// Check if it is a Login Form with Input Type Button
		if html.SelfClosingTagToken == tokenType && htmlStat.Form == "AVAILABLE" && currentTokenTag == "input" {
			inputType := ""
			inputValue := ""
			for _, attr := range token.Attr {
				fmt.Println("input Link: Val: [", attr.Val, "], Key: [", attr.Key, "]")
				if attr.Key == "type" {
					inputType = attr.Val
				}
				if attr.Key == "value" {
					inputValue = strings.ToLower(strings.ReplaceAll(attr.Val, " ", ""))
				}
			}
			// Check if the Button is Login or Signin button
			if (inputType == "button" || inputType == "submit") &&
				(inputValue == "login" || inputValue == "signin") {
				htmlStat.LoginForm = "AVAILABLE"
			}
		}

		// Check if it is a Login Form with Button
		if html.TextToken == tokenType && htmlStat.Form == "AVAILABLE" && currentTokenTag == "button" {
			buttonText := strings.ToLower(strings.ReplaceAll(token.Data, " ", ""))
			// Check if the Button is Login or Signin button
			if buttonText == "login" || buttonText == "signin" {
				htmlStat.LoginForm = "AVAILABLE"
			}
		}

		if currentTokenTag == "a" || currentTokenTag == "link" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
				}

			}
		}

		if html.SelfClosingTagToken == tokenType {
			if currentTokenTag == "img" {
				for _, attr := range token.Attr {
					if attr.Key == "src" {
						links = append(links, attr.Val)
					}
				}
			}
			currentTokenTag = ""
		}
	}
}

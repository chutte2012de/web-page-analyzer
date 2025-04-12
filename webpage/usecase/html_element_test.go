package usecase

import (
	"fmt"
	"testing"
)

// go test -v ./...

func TestExtractFromUrl(t *testing.T) {

	h := HtmlElement{}
	htmlStat, err := h.ExtractFromUrl("https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package")
	if err != nil {
		t.Fatalf("unexpected error extracting HTML Stat: %v", err)
	}
	fmt.Println("htmlStat:", htmlStat)
	fmt.Println("htmlStat URL:", htmlStat.Url)
	fmt.Println("htmlStat CreatedAt:", htmlStat.CreatedAt)
	fmt.Println("htmlStat Id:", htmlStat.Id)
}

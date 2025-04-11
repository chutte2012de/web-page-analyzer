package usecase

import "testing"

func TestExtractFromUrl(t *testing.T) {

	h := HtmlElement{}
	h.ExtractFromUrl("https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package")
}

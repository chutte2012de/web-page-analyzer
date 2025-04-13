package usecase

import (
	"fmt"
	"testing"
)

func TestValidateLinks(t *testing.T) {

	linkMan := LinkManager{}
	urls := []string{"http://example.com", "https://www.freecodecamp.org",
		"https://www.behance.net", "https://codepen.io", "http://example.org",
		"http://example.net", "http://exampleapple.net/"}
	results, err := linkMan.ValidateLinks("", urls)
	if err != nil {
		t.Fatalf("unexpected error validating Links: %v", err)
	}
	fmt.Println("results:", results)
}

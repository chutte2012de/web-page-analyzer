package usecase

import "testing"

func TestValidateLinks(t *testing.T) {

	linkMan := LinkManager{}
	urls := []string{"http://example.com", "https://www.freecodecamp.org",
		"https://www.behance.net", "https://codepen.io", "http://example.org",
		"http://example.net", "http://exampleapple.net/"}
	linkMan.ValidateLinks("", urls)
}

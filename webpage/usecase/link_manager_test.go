package usecase

import (
	"fmt"
	"testing"
)

func TestValidateLinks(t *testing.T) {

	return
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

func TestGetHostAndWebsiteBaseNames(t *testing.T) {
	linkMan := LinkManager{}

	hostName, websiteName, err := linkMan.GetHostAndWebsiteBaseNames("https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package")
	if err != nil {
		t.Fatalf("unexpected error Extracting Host name and Website name: %v", err)
	}
	fmt.Println("hostName:", hostName)
	fmt.Println("websiteName:", websiteName)
}

func TestCategorizeLinks(t *testing.T) {
	linkMan := LinkManager{}

	urls := []string{"http://example.com", "https://www.freecodecamp.org",
		"https://www.behance.net", "https://codepen.io", "http://example.org",
		"http://example.net", "http://exampleapple.net/", "/wavemakers", "https://docs.digitalocean.com/products/getting-started",
		"https://docs.digitalocean.com/products/compute", "https://docs.digitalocean.com/products/genai-platform",
		"https://docs.digitalocean.com/products/storage", "https://docs.digitalocean.com/products/databases", "https://docs.digitalocean.com/products/container-registry/", "https://docs.digitalocean.com/products/billing", "https://docs.digitalocean.com/reference/api", "/partners/pod",
		"/partners/services", "https://marketplace.digitalocean.com/", "/hatch", "/partners/directory", "/customers", "https://www.youtube.com/playlist?list=PLseEp7p6Ewibnv09L_48W3bi2HKiY6lrx", "https://ugurus.com/start-here/?utm_source=DO&utm_medium=partners&utm_content=menu",
		"/pricing/calculator", "/resources/articles/cloud-cost-optimization",
		"/resources/cloud-service-providers-how-to-choose", "/resources/articles/digitalocean-vs-awslightsail", "/company/contact/sales?referrer=mainmenu/partners", "/products/ai-ml/1-click-models", "/pricing",
		"https://www.digitalocean.com/api/dynamic-content/v1/login?success_redirect=https%3A%2F%2Fwww.digitalocean.com&error_redirect=https%3A%2F%2Fwww.digitalocean.com%2Fauth-error&type=login",
		"https://cloud.digitalocean.com/login", "https://www.digitalocean.com/api/dynamic-content/v1/login?success_redirect=https%3A%2F%2Fwww.digitalocean.com&error_redirect=https%3A%2F%2Fwww.digitalocean.com%2Fauth-error&type=register",
		"https://cloud.digitalocean.com/registrations/new", "/blog", "https://docs.digitalocean.com/products",
		"/support", "/company/contact/sales?referrer=tophat",
		"https://www.digitalocean.com/api/dynamic-content/v1/login?success_redirect=https%3A%2F%2Fwww.digitalocean.com&error_redirect=https%3A%2F%2Fwww.digitalocean.com%2Fauth-error&type=login", "https://cloud.digitalocean.com/login",
		"https://www.digitalocean.com/api/dynamic-content/v1/login?success_redirect=https%3A%2F%2Fwww.digitalocean.com&error_redirect=https%3A%2F%2Fwww.digitalocean.com%2Fauth-error&type=register",
		"https://cloud.digitalocean.com/registrations/new", "/community/tutorials", "/community/questions", "https://docs.digitalocean.com",
		"/community/pages/cloud-chats", "/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package#prerequisites", "/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package#step-1-mdash-creating-a-sample-program-to-unit-test",
		"/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package#step-2-mdash-writing-unit-tests-in-go",
		"/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package#step-3-mdash-testing-your-go-code-using-the-go-test-command"}

	fmt.Println("Total Input Links Len:", len(urls))
	internalLinksWithoutPreffix, internalLinksWithPreffix, externalLinks, err := linkMan.CategorizeLinks("digitalocean.com", urls)
	if err != nil {
		t.Fatalf("unexpected error Categorizing Links: %v", err)
	}
	fmt.Println("internalLinksWithoutPreffix Len:", len(internalLinksWithoutPreffix))
	fmt.Println("internalLinksWithoutPreffix:", internalLinksWithoutPreffix)
	fmt.Println("internalLinksWithPreffix Len:", len(internalLinksWithPreffix))
	fmt.Println("internalLinksWithPreffix:", internalLinksWithPreffix)
	fmt.Println("externalLinks Len:", len(externalLinks))
	fmt.Println("externalLinks:", externalLinks)

}

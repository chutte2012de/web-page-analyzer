package usecase

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestGetHostAndWebsiteBaseNames_FromLink(t *testing.T) {
	linkMan := LinkManager{}

	hostName, websiteName, err := linkMan.GetHostAndWebsiteBaseNames("https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package")
	if err != nil {
		t.Fatalf("unexpected error Extracting Host name and Website name: %v", err)
	}
	fmt.Println("hostName:", hostName)
	fmt.Println("websiteName:", websiteName)
	assert.Equal(t, "www.digitalocean.com", hostName)
	assert.Equal(t, "digitalocean.com", websiteName)
}

func TestGetHostAndWebsiteBaseNames_FromMainWebSiteUrl(t *testing.T) {
	linkMan := LinkManager{}

	hostName, websiteName, err := linkMan.GetHostAndWebsiteBaseNames("http://www.facebook.com")
	if err != nil {
		t.Fatalf("unexpected error Extracting Host name and Website name: %v", err)
	}
	fmt.Println("hostName:", hostName)
	fmt.Println("websiteName:", websiteName)
	assert.Equal(t, "www.facebook.com", hostName)
	assert.Equal(t, "facebook.com", websiteName)
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

	assert.Equal(t, 23, len(internalLinksWithoutPreffix))
	assert.Equal(t, "/wavemakers", internalLinksWithoutPreffix[0])
	assert.Equal(t, "/partners/pod", internalLinksWithoutPreffix[1])
	assert.Equal(t, "/partners/services", internalLinksWithoutPreffix[2])
	assert.Equal(t, "/hatch", internalLinksWithoutPreffix[3])
	assert.Equal(t, "/partners/directory", internalLinksWithoutPreffix[4])
	assert.Equal(t, "/customers", internalLinksWithoutPreffix[5])

	assert.Equal(t, "/pricing/calculator", internalLinksWithoutPreffix[6])
	assert.Equal(t, "/resources/articles/cloud-cost-optimization", internalLinksWithoutPreffix[7])
	assert.Equal(t, "/resources/cloud-service-providers-how-to-choose", internalLinksWithoutPreffix[8])
	assert.Equal(t, "/resources/articles/digitalocean-vs-awslightsail", internalLinksWithoutPreffix[9])
	assert.Equal(t, "/company/contact/sales?referrer=mainmenu/partners", internalLinksWithoutPreffix[10])
	assert.Equal(t, "/products/ai-ml/1-click-models", internalLinksWithoutPreffix[11])

	assert.Equal(t, "/pricing", internalLinksWithoutPreffix[12])
	assert.Equal(t, "/blog", internalLinksWithoutPreffix[13])
	assert.Equal(t, "/support", internalLinksWithoutPreffix[14])
	assert.Equal(t, "/company/contact/sales?referrer=tophat", internalLinksWithoutPreffix[15])
	assert.Equal(t, "/community/tutorials", internalLinksWithoutPreffix[16])
	assert.Equal(t, "/community/questions", internalLinksWithoutPreffix[17])

	assert.Equal(t, "/community/pages/cloud-chats", internalLinksWithoutPreffix[18])
	assert.Equal(t, "/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package#prerequisites", internalLinksWithoutPreffix[19])
	assert.Equal(t, "/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package#step-1-mdash-creating-a-sample-program-to-unit-test", internalLinksWithoutPreffix[20])
	assert.Equal(t, "/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package#step-2-mdash-writing-unit-tests-in-go", internalLinksWithoutPreffix[21])
	assert.Equal(t, "/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package#step-3-mdash-testing-your-go-code-using-the-go-test-command", internalLinksWithoutPreffix[22])

	assert.Equal(t, 19, len(internalLinksWithPreffix))
	assert.Equal(t, "https://docs.digitalocean.com/products/getting-started", internalLinksWithPreffix[0])
	assert.Equal(t, "https://docs.digitalocean.com/products/compute", internalLinksWithPreffix[1])
	assert.Equal(t, "https://docs.digitalocean.com/products/genai-platform", internalLinksWithPreffix[2])
	assert.Equal(t, "https://docs.digitalocean.com/products/storage", internalLinksWithPreffix[3])
	assert.Equal(t, "https://docs.digitalocean.com/products/databases", internalLinksWithPreffix[4])
	assert.Equal(t, "https://docs.digitalocean.com/products/container-registry/", internalLinksWithPreffix[5])

	assert.Equal(t, "https://docs.digitalocean.com/products/billing", internalLinksWithPreffix[6])
	assert.Equal(t, "https://docs.digitalocean.com/reference/api", internalLinksWithPreffix[7])
	assert.Equal(t, "https://marketplace.digitalocean.com/", internalLinksWithPreffix[8])
	assert.Equal(t, "https://www.digitalocean.com/api/dynamic-content/v1/login?success_redirect=https%3A%2F%2Fwww.digitalocean.com&error_redirect=https%3A%2F%2Fwww.digitalocean.com%2Fauth-error&type=login", internalLinksWithPreffix[9])
	assert.Equal(t, "https://cloud.digitalocean.com/login", internalLinksWithPreffix[10])
	assert.Equal(t, "https://www.digitalocean.com/api/dynamic-content/v1/login?success_redirect=https%3A%2F%2Fwww.digitalocean.com&error_redirect=https%3A%2F%2Fwww.digitalocean.com%2Fauth-error&type=register", internalLinksWithPreffix[11])

	assert.Equal(t, "https://cloud.digitalocean.com/registrations/new", internalLinksWithPreffix[12])
	assert.Equal(t, "https://docs.digitalocean.com/products", internalLinksWithPreffix[13])
	assert.Equal(t, "https://www.digitalocean.com/api/dynamic-content/v1/login?success_redirect=https%3A%2F%2Fwww.digitalocean.com&error_redirect=https%3A%2F%2Fwww.digitalocean.com%2Fauth-error&type=login", internalLinksWithPreffix[14])
	assert.Equal(t, "https://cloud.digitalocean.com/login", internalLinksWithPreffix[15])
	assert.Equal(t, "https://www.digitalocean.com/api/dynamic-content/v1/login?success_redirect=https%3A%2F%2Fwww.digitalocean.com&error_redirect=https%3A%2F%2Fwww.digitalocean.com%2Fauth-error&type=register", internalLinksWithPreffix[16])
	assert.Equal(t, "https://cloud.digitalocean.com/registrations/new", internalLinksWithPreffix[17])

	assert.Equal(t, "https://docs.digitalocean.com", internalLinksWithPreffix[18])

	assert.Equal(t, 9, len(externalLinks))
	assert.Equal(t, "http://example.com", externalLinks[0])
	assert.Equal(t, "https://www.freecodecamp.org", externalLinks[1])
	assert.Equal(t, "https://www.behance.net", externalLinks[2])
	assert.Equal(t, "https://codepen.io", externalLinks[3])
	assert.Equal(t, "http://example.org", externalLinks[4])
	assert.Equal(t, "http://example.net", externalLinks[5])
	assert.Equal(t, "http://exampleapple.net/", externalLinks[6])
	assert.Equal(t, "https://www.youtube.com/playlist?list=PLseEp7p6Ewibnv09L_48W3bi2HKiY6lrx", externalLinks[7])
	assert.Equal(t, "https://ugurus.com/start-here/?utm_source=DO&utm_medium=partners&utm_content=menu", externalLinks[8])
}

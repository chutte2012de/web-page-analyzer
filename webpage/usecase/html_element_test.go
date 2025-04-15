package usecase

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test -v ./...
// https://dev.to/dave3130/golang-html-tokenizer-5fh7
// https://stackoverflow.com/questions/74426546/how-to-calculate-total-code-coverage-for-tests
// https://medium.com/zus-health/mocking-outbound-http-requests-in-go-youre-probably-doing-it-wrong-60373a38d2aa

func TestExtractFromUrl(t *testing.T) {
	//return
	l := NewLinkManager()
	h := NewHtmlElement(l)
	htmlStat, err := h.ExtractFromUrl("https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package")
	if err != nil {
		t.Fatalf("unexpected error extracting HTML Stat: %v", err)
	}
	fmt.Println("htmlStat:", htmlStat)
	fmt.Println("htmlStat URL:", htmlStat.Url)
	fmt.Println("htmlStat CreatedAt:", htmlStat.CreatedAt)
	fmt.Println("htmlStat Id:", htmlStat.Id)
}

func TestExtractFromIoReader(t *testing.T) {

	l := NewLinkManager()
	h := NewHtmlElement(l)
	const sampleHtml = `<!DOCTYPE html><html><head><style> body {background-color: powderblue;} h1 {color: red;} p {color: orange;}</style><title>Sample HTML Code</title><script src="my-script.js">abc</script></head><body><h1>Main title</h1><p id="demo"></p><a href="https://dev.to/">Dev Community</a><script>document.getElementById("demo").innerHTML = "Hello JavaScript!";</script></body></html>`
	htmlStat, links, err := h.ExtractFromIoReader(strings.NewReader(sampleHtml))
	if err != nil {
		t.Fatalf("unexpected error extracting HTML Stat: %v", err)
	}
	fmt.Println("htmlStat:", htmlStat)
	fmt.Println("htmlStat URL:", htmlStat.Url)
	fmt.Println("htmlStat CreatedAt:", htmlStat.CreatedAt)
	fmt.Println("htmlStat Id:", htmlStat.Id)
	fmt.Println("links:", links)

	assert.Equal(t, "", htmlStat.Url)
	assert.Equal(t, "html", htmlStat.Version)
	assert.Equal(t, "Sample HTML Code", htmlStat.Title)

	assert.Equal(t, uint16(1), htmlStat.Headers.H1)
	assert.Equal(t, uint16(0), htmlStat.Headers.H2)
	assert.Equal(t, uint16(0), htmlStat.Headers.H3)
	assert.Equal(t, uint16(0), htmlStat.Headers.H4)
	assert.Equal(t, uint16(0), htmlStat.Headers.H5)
	assert.Equal(t, uint16(0), htmlStat.Headers.H6)

	assert.Equal(t, 1, len(links))
	assert.Equal(t, "https://dev.to/", links[0])
}

func TestExtractFromIoReader_HtmlWithExternalLinks(t *testing.T) {
	l := NewLinkManager()
	h := NewHtmlElement(l)
	const sampleHtml = `<!DOCTYPE html>
<html>
<head>
	<meta charset="utf-8">
	<meta http-equiv="X-UA-Compatible" content="IE=edge">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">

	<title>Dinesh Pandiyan | Awesome Dev</title>

	<link rel="icon" href="favicon.ico" type="image/png" />

	<link href="https://fonts.googleapis.com/css?family=Reem+Kufi|Roboto:300" rel="stylesheet">
	<link href="https://use.fontawesome.com/releases/v5.13.1/css/all.css" rel="stylesheet">
	<link rel="stylesheet" href="css/reset.css">
	<link rel="stylesheet" href="css/styles.css">
	<link rel="stylesheet" href="css/themes/indigo-white.css">
	<!-- <link rel="stylesheet" href="css/themes/green-white.css"> -->
	<!-- <link rel="stylesheet" href="css/themes/red-white.css"> -->
	<!-- <link rel="stylesheet" href="css/themes/grey-white.css"> -->
	<!-- <link rel="stylesheet" href="css/themes/white-indigo.css"> -->
	<!-- <link rel="stylesheet" href="css/themes/white-blue.css"> -->
	<!-- <link rel="stylesheet" href="css/themes/white-grey.css"> -->
	<!-- <link rel="stylesheet" href="css/themes/white-red.css"> -->
	<!-- <link rel="stylesheet" href="css/themes/yellow-black.css"> -->
</head>
<body>
	<main>
		<h1>Heading 1</h1>
		<h2>Heading 2</h2>
		<h3>Heading 3</h3>
		<h4>Heading 4</h4>
		<h5>Heading 5</h5>
		<h6>Heading 6</h6>
		<h2>Heading 2</h2>
		<h3>Heading 3</h3>
		<h4>Heading 4</h4>
		<h5>Heading 5</h5>
		<h6>Heading 6</h6>
		<h3>Heading 3</h3>
		<h4>Heading 4</h4>
		<h5>Heading 5</h5>
		<h6>Heading 6</h6>
		<h4>Heading 4</h4>
		<h5>Heading 5</h5>
		<h6>Heading 6</h6>
		<h5>Heading 5</h5>
		<h6>Heading 6</h6>
		<h6>Heading 6</h6>
		<div class="intro">Hello, I'm Dinesh!</div>
		<div class="tagline">All-Star Dev | Code Fanatic | Linux Hacker | Bleh</div>
		<!-- Find your icons from here - https://fontawesome.com/icons?d=gallery&s=brands -->
		<div class="icons-social">
			<a target="_blank" href="https://github.com/flexdinesh">
        <i class="fab fa-github" aria-hidden="true" title="Github"></i>
        <span class="sr-only">Github</span>
      </a>
      <a target="_blank" href="https://twitter.com/flexdinesh">
        <i class="fab fa-twitter" aria-hidden="true" title="Twitter"></i>
        <span class="sr-only">Twitter</span>
      </a>
      <a target="_blank" href="https://dev.to/flexdinesh">
        <i class="fab fa-dev" aria-hidden="true" title="DEV Community"></i>
        <span class="sr-only">DEV Community</span>
      </a>
      <a target="_blank" href="https://stackoverflow.com/story/flexdinesh">
        <i class="fab fa-stack-overflow" aria-hidden="true" title="StackOverflow"></i>
        <span class="sr-only">StackOverflow</span>
      </a>
      <a target="_blank" href="https://www.linkedin.com/in/dineshpandiyan">
        <i class="fab fa-linkedin" aria-hidden="true" title="LinkedIn"></i>
        <span class="sr-only">LinkedIn</span>
      </a>
      <a target="_blank" href="https://medium.com/@flexdinesh">
        <i class="fab fa-medium" aria-hidden="true" title="Medium"></i>
        <span class="sr-only">Medium</span>
      </a>
      <a target="_blank" href="https://www.freecodecamp.org">
        <i class="fab fa-free-code-camp" aria-hidden="true" title="FreeCodeCamp"></i>
        <span class="sr-only">FreeCodeCamp</span>
      </a>
      <a target="_blank" href="https://www.behance.net">
        <i class="fab fa-behance" aria-hidden="true" title="Behance"></i>
        <span class="sr-only">Behance</span>
      </a>
      <a target="_blank" href="https://codepen.io">
        <i class="fab fa-codepen" aria-hidden="true" title="CodePen"></i>
        <span class="sr-only">CodePen</span>
      </a>
    </div>
	</main>
</body>
</html>`
	htmlStat, links, err := h.ExtractFromIoReader(strings.NewReader(sampleHtml))
	if err != nil {
		t.Fatalf("unexpected error extracting HTML Stat: %v", err)
	}
	fmt.Println("htmlStat:", htmlStat)
	fmt.Println("htmlStat URL:", htmlStat.Url)
	fmt.Println("htmlStat CreatedAt:", htmlStat.CreatedAt)
	fmt.Println("htmlStat Id:", htmlStat.Id)
	fmt.Println("htmlStat Title:", htmlStat.Title)
	fmt.Println("links:", links)

	assert.Equal(t, uint16(1), htmlStat.Headers.H1)
	assert.Equal(t, uint16(2), htmlStat.Headers.H2)
	assert.Equal(t, uint16(3), htmlStat.Headers.H3)
	assert.Equal(t, uint16(4), htmlStat.Headers.H4)
	assert.Equal(t, uint16(5), htmlStat.Headers.H5)
	assert.Equal(t, uint16(6), htmlStat.Headers.H6)

	assert.Equal(t, "", htmlStat.Url)
	assert.Equal(t, "html", htmlStat.Version)
	assert.Equal(t, "Dinesh Pandiyan | Awesome Dev", htmlStat.Title)
}

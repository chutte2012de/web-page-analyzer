"# web-page-analyzer"

"Run locally"
`go run main.go`

"Run in DOcker"
`docker build -t web-page-analyzer .`
`docker run -p 3000:3000 web-page-analyzer`

"Run All Unit Tests"
`go test -v ./...`

"Check Unit Tests Coverage"
`go test ./... -coverprofile=cover.out`

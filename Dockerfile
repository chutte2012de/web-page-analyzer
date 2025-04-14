# docker build -t web-page-analyzer .
# docker run -p 3000:3000 web-page-analyzer
FROM golang:1.24.2
WORKDIR /app
COPY . .
RUN go build -o app main.go
EXPOSE 3000
CMD ["./app"]
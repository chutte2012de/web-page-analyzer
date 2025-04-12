package model

import (
	"time"
)

type HtmlStat struct {
	Id          uint64     `json:"id"`
	Url         string     `json:"url"`
	Version     string     `json:"version"`
	Title       string     `json:"title"`
	Links       Links      `json:"links"`
	Headers     Headers    `json:"headers"`
	Form        string     `json:"form"`
	LoginForm   string     `json:"login_form"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

type Headers struct {
	H1 uint `json:"h1"`
	H2 uint `json:"h2"`
	H3 uint `json:"h3"`
	H4 uint `json:"h4"`
	H5 uint `json:"h5"`
	H6 uint `json:"h6"`
}

type Links struct {
	Internal []Link `json:"internal"`
	External []Link `json:"external"`
}

type Link struct {
	Url    string `json:"url"`
	Status string `json:"status"`
}

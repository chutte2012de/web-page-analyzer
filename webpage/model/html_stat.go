package model

import (
	"time"
)

type HtmlStat struct {
	SummaryId   uint64     `json:"summary_id"`
	Url         string     `json:"url"`
	Title       string     `json:"title"`
	Links       Links      `json:"links"`
	Headers     Headers    `json:"headers"`
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
	InboundLinks  []Link `json:"inbound_links"`
	OutboundLinks []Link `json:"outbound_links"`
}

type Link struct {
	Url    string `json:"url"`
	Status string `json:"status"`
}

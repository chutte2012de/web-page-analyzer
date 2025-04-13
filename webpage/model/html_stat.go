package model

import (
	"time"
)

type HtmlStat struct {
	Id          uint64     `json:"id"`
	Url         string     `json:"url"`
	Version     string     `json:"version"`
	Title       string     `json:"title"`
	LinksInfo   LinksInfo  `json:"links_info"`
	Headers     Headers    `json:"headers"`
	Form        string     `json:"form"`
	LoginForm   string     `json:"login_form"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

type Headers struct {
	H1 uint16 `json:"h1"`
	H2 uint16 `json:"h2"`
	H3 uint16 `json:"h3"`
	H4 uint16 `json:"h4"`
	H5 uint16 `json:"h5"`
	H6 uint16 `json:"h6"`
}

type LinksInfo struct {
	Summary LinksSummary `json:"summary"`
	Detail  LinksDetail  `json:"detail"`
}

type LinksSummary struct {
	Internal SummaryStat `json:"internal"`
	External SummaryStat `json:"external"`
}

type SummaryStat struct {
	TotalCount       uint16 `json:"total_count"`
	ReachableCount   uint16 `json:"reachable_count"`
	UnreachableCount uint16 `json:"unreachable_count"`
}

type LinksDetail struct {
	Internal []Link `json:"internal"`
	External []Link `json:"external"`
}

type Link struct {
	Url       string `json:"url"`
	Reachable string `json:"reachable"`
	Status    string `json:"status"`
}

package handler

import (
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/chutte2012de/web-page-analyzer/webpage/model"
)

type HtmlStat struct {
}

func (h *HtmlStat) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Url string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Println("failed to decode in created:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	now := time.Now().UTC()

	links := model.Links{}

	html_stat := model.HtmlStat{
		Id:    rand.Uint64(),
		Url:   body.Url,
		Links: links,

		CreatedAt: &now,
	}

	res, err := json.Marshal(html_stat)
	if err != nil {
		fmt.Println("failed to marshal:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

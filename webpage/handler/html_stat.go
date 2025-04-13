package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/chutte2012de/web-page-analyzer/webpage/usecase"
)

type HtmlStat struct {
	HtmlElement *usecase.HtmlElement
}

func (h *HtmlStat) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Url string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		slog.Error("Failed to decode Request Body", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	html_stat, err := h.HtmlElement.ExtractFromUrl(body.Url)
	if err != nil {
		slog.Error("Failed to Extract Html Element Statistics", "errpr", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	res, err := json.Marshal(html_stat)
	if err != nil {
		slog.Error("Failed to Marshal Html Element Statistics", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

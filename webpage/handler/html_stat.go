package handler

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/chutte2012de/web-page-analyzer/webpage/usecase"
)

type HtmlStat struct {
	HtmlElement usecase.IHtmlElement
}

func (h *HtmlStat) Create(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Url string `json:"url"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		slog.Error("Failed to decode Request Body", "error", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	htmlStat, err := h.HtmlElement.ExtractFromUrl(body.Url)
	if err != nil {
		slog.Error("Failed to Extract Html Element Statistics", "errpr", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	slog.Info("Html Stat Received", "htmlStat", htmlStat)

	res, err := json.Marshal(htmlStat)
	if err != nil {
		slog.Error("Failed to Marshal Html Element Statistics", "error", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	slog.Info("Json Html Stat", "res", res)

	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

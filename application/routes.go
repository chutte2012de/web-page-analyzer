package application

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/chutte2012de/web-page-analyzer/webpage/handler"
	"github.com/chutte2012de/web-page-analyzer/webpage/usecase"
)

func (a *App) loadRoutes() {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Route("/webpage/htmlstat", a.loadOrderRoutes)
	router.Handle("/metrics", promhttp.Handler())

	a.router = router
}

func (a *App) loadOrderRoutes(router chi.Router) {

	linkManager := usecase.NewLinkManager()
	htmlElement := usecase.NewHtmlElement(linkManager)
	webPageHtmlStatHandler := &handler.HtmlStat{
		HtmlElement: htmlElement,
	}

	router.Post("/", webPageHtmlStatHandler.Create)
}

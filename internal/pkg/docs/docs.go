package docs

import (
	_ "embed"
	"net/http"

	"github.com/go-chi/chi"
)

const (
	contentTypeHTML = "text/html"
	contentTypeJSON = "application/json"
)

//go:embed index.html
var indexPageContent []byte

type ServiceDoc struct {
	URL     string
	Content []byte
}

func NewServiceDocsHandler(docs ...ServiceDoc) http.Handler {
	mux := chi.NewMux()

	for _, doc := range docs {
		mount(mux, doc.URL, doc.Content)
	}

	return mux
}

func mount(mux *chi.Mux, prefix string, content []byte) {
	mux.Route("/"+prefix+"/", func(router chi.Router) {
		router.Get("/swagger.json", handleContent(contentTypeJSON, content))
		router.Mount("/", handleContent(contentTypeHTML, indexPageContent))
	})
}

func handleContent(contentType string, content []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", contentType)
		w.Write(content)
	}
}

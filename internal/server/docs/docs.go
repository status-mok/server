package docs

import (
	_ "embed"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/status-mok/server/internal/server/docs/stoplight_elements"
	serverAPI "github.com/status-mok/server/pkg/server-api"
)

var (
	serverAPIurl = "server-api"
)

func Handler() http.Handler {
	mux := chi.NewMux()

	mount(mux, serverAPIurl, serverAPI.SwaggerJSON)

	return mux
}

func mount(mux *chi.Mux, prefix string, content []byte) {
	mux.Route("/"+prefix, func(router chi.Router) {
		router.Get("/swagger.json", handleJSON(content))
		router.Mount("/", http.StripPrefix(
			"/docs/"+prefix,
			http.FileServer(http.FS(stoplight_elements.FS)),
		))
	})
}

func handleJSON(content []byte) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write(content)
	}
}

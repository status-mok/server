package mok

import (
	"context"
	"net/http"
	"net/http/httputil"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/status-mok/server/internal/pkg/log"
)

func (s *mokServer) httpHandler(ctx context.Context) http.Handler {
	mux := chi.NewMux()

	mux.Use(middleware.RealIP)
	middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: log.StdLogger(ctx), NoColor: true})
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mux.HandleFunc("/", s.httpHandlerFunc)

	return mux
}

func (s *mokServer) httpHandlerFunc(_ http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// 1 log request
	req, err := httputil.DumpRequest(r, true)
	if err != nil {
		log.L(ctx).With(
			"server", s.Name(),
		).Debug("failed to dump request")
	} else {
		log.L(ctx).With(
			"server", s.Name(),
			"request", string(req),
		).Debug("request dump")
	}

	// 2 find endpoint
	// 3 decide if request match any endpoint mock
	// 4 respond
}

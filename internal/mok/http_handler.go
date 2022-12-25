package mok

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/status-mok/server/internal/pkg/log"
)

func (s *server) httpHandler(ctx context.Context) http.Handler {
	mux := chi.NewMux()

	mux.Use(middleware.RealIP)
	middleware.DefaultLogger = middleware.RequestLogger(&middleware.DefaultLogFormatter{Logger: log.StdLogger(ctx), NoColor: true})
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)

	mux.Mount("/", serverHTTPHandler(s))

	return mux
}

func serverHTTPHandler(s *server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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

		s.httpMux.ServeHTTP(w, r)
	})
}

func routeHTTPHandler(rt Route) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if rt.Disabled() {
			notFoundResponse(w)
			return
		}

		exp, err := rt.ExpectationFindMatch(r.Context(), rt.Type(), r)
		if err != nil {
			notFoundResponse(w)
			return
		}

		fmt.Println(exp)
	})
}

func notFoundResponse(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}

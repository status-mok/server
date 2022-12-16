package mok

import (
	"net/http"

	"github.com/status-mok/server/internal/pkg/log"
	"github.com/status-mok/server/internal/pkg/request"
)

func (s *mokServer) httpHandler() http.Handler {
	return http.Handler(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				body, err := request.ReadBody(r)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				// 1 log request
				log.L(r.Context()).With(
					"server", s,
					"request", body,
				).Debug("request received")

				// 2 find endpoint
				// 3 decide if request match any endpoint mock
				// 4 respond
			},
		),
	)
}

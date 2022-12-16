package request

import (
	"bytes"
	"io"
	"net/http"
)

func ReadBody(req *http.Request) ([]byte, error) {
	if req == nil || req.Body == nil {
		return nil, nil
	}

	var err error
	var withGetBody bool

	body := req.Body
	if req.GetBody != nil {
		withGetBody = true

		body, err = req.GetBody()
		if err != nil {
			return nil, err
		}
	}

	bodyContent, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	if !withGetBody {
		req.Body = io.NopCloser(bytes.NewReader(bodyContent))

		req.GetBody = func() (io.ReadCloser, error) {
			return io.NopCloser(bytes.NewReader(bodyContent)), nil
		}
	}

	return bodyContent, nil
}

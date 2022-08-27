package v1

import (
	"io"
	"net/http"
)

func readBody(w http.ResponseWriter, r *http.Request) ([]byte, bool) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return body, false
	}
	return body, true
}

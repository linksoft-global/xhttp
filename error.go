package xhttp

import (
	"fmt"
	"net/http"
)

func Error(w http.ResponseWriter, statusCode int, id, context string) {
	tmpl := `{"id": %q, "ctx": %q, "msg": %q}`

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(statusCode)

	fmt.Fprintf(w, tmpl, id, context, http.StatusText(statusCode))
}

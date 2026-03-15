package middleware

import (
	"net/http"
	"strings"
)

// stripTrailingSlash removes trailing slashes and normalizes empty paths to "/"
func StripTrailingSlash(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if path != "/" && strings.HasSuffix(path, "/") {
			path = strings.TrimSuffix(path, "/")
		}
		if path == "" {
			path = "/"
		}
		r.URL.Path = path
		h.ServeHTTP(w, r)
	})
}

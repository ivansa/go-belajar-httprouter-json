package middleware

import (
	"fmt"
	"net/http"
)

type LogMiddleware struct {
	http.Handler
}

func (lm *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Request: %s %s\n", r.Method, r.URL.Path)
	lm.Handler.ServeHTTP(w, r)
}

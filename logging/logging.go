package logging

// Poor man's logging :)

import (
	"fmt"
	"net/http"
)

func LogRequest(r *http.Request) {
	fmt.Printf("Request: %s %s %s %v %s\n", r.Method, r.URL, r.Proto, r.Header["User-Agent"], r.RemoteAddr)
}

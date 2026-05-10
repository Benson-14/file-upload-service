package middleware

import (
	"fmt"
	"net/http"
	"time"
)

type ReqLogResponseWriter struct {
	http.ResponseWriter
	status int
}

func (rlrw *ReqLogResponseWriter) WriteHeader(status int) {
	rlrw.status = status
	rlrw.ResponseWriter.WriteHeader(status)
}

func (rlrw *ReqLogResponseWriter) Write(b []byte) (int, error) {
	if rlrw.status == 0 {
		rlrw.status = http.StatusOK
	}
	return rlrw.ResponseWriter.Write(b)
}

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rlrw := &ReqLogResponseWriter{
			w,
			http.StatusOK,
		}

		next.ServeHTTP(rlrw, r)

		fmt.Printf("File Upload Service %s %d %s %s\n", r.Method, rlrw.status, r.URL.Path, time.Since(start))
	})
}

package middelware

import (
	"log"
	"net/http"
	"time"
)

// LoggingMiddleware 日志记录中间件
func LoggingMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next(w, r)
		log.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	}
}

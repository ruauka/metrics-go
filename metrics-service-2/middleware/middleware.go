package middleware

import (
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

// MetricsMiddleware prometheus for http-server.
type MetricsMiddleware struct {
	*Metrics
}

// NewMetricsMiddleware returns handler metrics for httprouter.
func NewMetricsMiddleware(metrics *Metrics) func(httprouter.Handle) httprouter.Handle {
	middleware := MetricsMiddleware{metrics}
	return middleware.Handler
}

// Handler calculation custom metrics.
func (metrics MetricsMiddleware) Handler(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()

		rw := newResponseWriter(w)

		next(rw, r, ps)

		code := strconv.Itoa(rw.status)
		labels := []string{code}

		metrics.Duration.WithLabelValues(labels...).Observe(float64(time.Since(start)) / float64(time.Second))
		metrics.Total.WithLabelValues(labels...).Inc()
	}
}

func newResponseWriter(w http.ResponseWriter) *responseWriterDelegator {
	return &responseWriterDelegator{ResponseWriter: w}
}

type responseWriterDelegator struct {
	http.ResponseWriter
	status      int
	written     int64
	wroteHeader bool
}

func (r *responseWriterDelegator) WriteHeader(code int) {
	r.status = code
	r.wroteHeader = true
	r.ResponseWriter.WriteHeader(code)
}

func (r *responseWriterDelegator) Write(b []byte) (int, error) {
	n, err := r.ResponseWriter.Write(b)
	r.written += int64(n)
	return n, err
}

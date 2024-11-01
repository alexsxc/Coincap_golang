package coincap

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type loggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (rt loggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	_, _ = fmt.Fprintf(rt.logger, "[%s] %s %s\n", time.Now().Format(time.ANSIC), req.Method, req.URL)
	return rt.next.RoundTrip(req)
}

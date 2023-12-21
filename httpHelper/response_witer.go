package httpHelper

import (
	"bytes"
	"io"
	"net/http"
)

type ResponseWriterLazy struct {
	Writer http.ResponseWriter
	Buffer bytes.Buffer
	Code   int
}

// Header
// Deprecated: The function is deprecated use the httpHelper2 package.
func (rw *ResponseWriterLazy) Header() http.Header {
	return rw.Writer.Header()
}

// WriteHeader
// Deprecated: The function is deprecated use the httpHelper2 package.
func (rw *ResponseWriterLazy) WriteHeader(statusCode int) {
	rw.Code = statusCode
}

// Write
// Deprecated: The function is deprecated use the httpHelper2 package.
func (rw *ResponseWriterLazy) Write(data []byte) (int, error) {
	return rw.Buffer.Write(data)
}

// Done
// Deprecated: The function is deprecated use the httpHelper2 package.
func (rw *ResponseWriterLazy) Done() (int64, error) {
	if rw.Code > 0 {
		rw.Writer.WriteHeader(rw.Code)
	}
	return io.Copy(rw.Writer, &rw.Buffer)
}

package httpHelper

import (
	"bytes"
	"io"
	"net/http"
)

type ResponseWriterLazy struct {
	w    http.ResponseWriter
	buf  bytes.Buffer
	code int
}

func (rw *ResponseWriterLazy) Header() http.Header {
	return rw.w.Header()
}

func (rw *ResponseWriterLazy) WriteHeader(statusCode int) {
	rw.code = statusCode
}

func (rw *ResponseWriterLazy) Write(data []byte) (int, error) {
	return rw.buf.Write(data)
}

func (rw *ResponseWriterLazy) Done() (int64, error) {
	if rw.code > 0 {
		rw.w.WriteHeader(rw.code)
	}
	return io.Copy(rw.w, &rw.buf)
}

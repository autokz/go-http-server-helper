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

func (rw *ResponseWriterLazy) Header() http.Header {
	return rw.Writer.Header()
}

func (rw *ResponseWriterLazy) WriteHeader(statusCode int) {
	rw.Code = statusCode
}

func (rw *ResponseWriterLazy) Write(data []byte) (int, error) {
	return rw.Buffer.Write(data)
}

func (rw *ResponseWriterLazy) Done() (int64, error) {
	if rw.Code > 0 {
		rw.Writer.WriteHeader(rw.Code)
	}
	return io.Copy(rw.Writer, &rw.Buffer)
}

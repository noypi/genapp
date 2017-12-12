package genapp

import (
	"bytes"
	"net/http"
)

type Writer struct {
	http.ResponseWriter
	header map[string][]string
	buf    *bytes.Buffer
	status int
}

func newWriter() *Writer {
	return &Writer{
		header: map[string][]string{},
		buf:    new(bytes.Buffer),
	}
}

func (w *Writer) Header() http.Header {
	return w.header
}

func (w *Writer) Write(bb []byte) (int, error) {
	w.buf.Write(bb)
	return w.buf.Len(), nil
}

func (w *Writer) WriteHeader(n int) {
	w.status = n
}

package gzip

import (
	"compress/gzip"
	"io"
	"sync"
)

var gzipBuffers = sync.Pool{
	New: func() interface{} {
		return new(gzip.Writer)
	},
}

func getGzip(w io.Writer) *gzip.Writer {
	gz := gzipBuffers.Get().(*gzip.Writer)
	gz.Reset(w)
	return gz
}

func putGzip(gz *gzip.Writer) {
	gzipBuffers.Put(gz)
}

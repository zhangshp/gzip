package gzip

import (
	"compress/gzip"
	"io"
	"sync"
)

var (
	gzipBuffers sync.Pool
)

func getGzip(dst io.Writer) (w *gzip.Writer) {
	if gz := gzipBuffers.Get(); gz != nil {
		w = gz.(*gzip.Writer)
		w.Reset(dst)
	} else {
		w, _ = gzip.NewWriterLevel(dst, gzip.BestCompression)
	}
	return
}

func putGzip(gz *gzip.Writer) {
	gz.Close()
	gzipBuffers.Put(gz)
}

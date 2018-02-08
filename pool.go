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
	g := gzipBuffers.Get().(*gzip.Writer)
	g.Reset(w)
	return g
}

func putGzip(g *gzip.Writer) {
	gzipBuffers.Put(g)
}

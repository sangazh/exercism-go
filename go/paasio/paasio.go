package paasio

import (
	"io"
	"sync"
)

type readCounter struct {
	r    io.Reader
	n    int64
	nops int
	mux  sync.Mutex
}

func (r *readCounter) ReadCount() (n int64, nops int) {
	r.mux.Lock()
	defer r.mux.Unlock()

	return r.n, r.nops
}

func (r *readCounter) Read(b []byte) (n int, err error) {
	r.mux.Lock()
	defer r.mux.Unlock()
	n, err = r.r.Read(b)

	r.nops += 1
	r.n += int64(n)
	return

}

func NewReadCounter(r io.Reader) ReadCounter {
	return &readCounter{r: r}
}

type writeCounter struct {
	w    io.Writer
	n    int64
	nops int
	mux  sync.Mutex
}

func (w *writeCounter) WriteCount() (n int64, nops int) {
	w.mux.Lock()
	defer w.mux.Unlock()
	return w.n, w.nops
}

func (w *writeCounter) Write(b []byte) (n int, err error) {
	w.mux.Lock()
	defer w.mux.Unlock()
	n, err = w.w.Write(b)

	w.nops += 1
	w.n += int64(n)

	return
}

func NewWriteCounter(w io.Writer) WriteCounter {
	return &writeCounter{w: w}
}

type readWriteCounter struct {
	rc ReadCounter
	wc WriteCounter
}

func (rw *readWriteCounter) Read(b []byte) (int, error) {
	return rw.rc.Read(b)
}
func (rw *readWriteCounter) Write(b []byte) (int, error) {
	return rw.wc.Write(b)
}
func (rw *readWriteCounter) ReadCount() (n int64, nops int) {
	return rw.rc.ReadCount()
}
func (rw *readWriteCounter) WriteCount() (n int64, nops int) {
	return rw.wc.WriteCount()
}

func NewReadWriteCounter(rw io.ReadWriter) ReadWriteCounter {
	return &readWriteCounter{rc: NewReadCounter(rw), wc: NewWriteCounter(rw)}
}

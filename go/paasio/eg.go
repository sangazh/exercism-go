package paasio

import (
	"io"
	"sync"
)

type MyWriteCounter struct {
	w    io.Writer
	n    int64
	nops int
	l    sync.Mutex
}

func (wc *MyWriteCounter) Write(p []byte) (n int, err error) {
	wc.l.Lock()
	defer wc.l.Unlock()
	n, err = wc.w.Write(p)
	wc.n += int64(n)
	wc.nops++
	return
}

func (wc *MyWriteCounter) WriteCount() (n int64, nops int) {
	wc.l.Lock()
	defer wc.l.Unlock()
	n = wc.n
	nops = wc.nops
	return
}

func NewWriteCounter1(w io.Writer) WriteCounter {
	wc := MyWriteCounter{w: w}
	return &wc
}

type MyReadCounter struct {
	r    io.Reader
	n    int64
	nops int
	l    sync.Mutex
}

func (rc *MyReadCounter) Read(p []byte) (n int, err error) {
	rc.l.Lock()
	defer rc.l.Unlock()
	n, err = rc.r.Read(p)
	rc.n += int64(n)
	rc.nops++
	return
}

func (rc *MyReadCounter) ReadCount() (n int64, nops int) {
	rc.l.Lock()
	defer rc.l.Unlock()
	n = rc.n
	nops = rc.nops
	return
}

func NewReadCounter1(r io.Reader) ReadCounter {
	return &MyReadCounter{r: r}
}

type MyReadWriteCounter struct {
	rc ReadCounter
	wc WriteCounter
}

func (rwc *MyReadWriteCounter) Read(p []byte) (n int, err error) {
	return rwc.rc.Read(p)
}

func (rwc *MyReadWriteCounter) Write(p []byte) (n int, err error) {
	return rwc.wc.Write(p)
}

func (rwc *MyReadWriteCounter) ReadCount() (n int64, nops int) {
	return rwc.rc.ReadCount()
}

func (rwc *MyReadWriteCounter) WriteCount() (n int64, nops int) {
	return rwc.wc.WriteCount()
}

func NewReadWriteCounter1(rw io.ReadWriter) ReadWriteCounter {
	return &MyReadWriteCounter{rc: NewReadCounter(rw), wc: NewWriteCounter(rw)}
}

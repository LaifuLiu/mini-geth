package rlp

import (
	"fmt"
	"io"
	"reflect"
	"sync"
)

// set a globe sync.pool
// write different write funcs

type EncBuffer struct {
	str []byte
}

var encBuffPool = sync.Pool{
	New: func() interface{} { return &EncBuffer{str: make([]byte, 1024)} },
}

func getEncBufferer() *EncBuffer {
	buf := encBuffPool.Get().(*EncBuffer)
	buf.reset()
	return buf
}

func (buf *EncBuffer) reset() {
	buf.str = buf.str[:0]
}

func (buf *EncBuffer) encode(val interface{}) error {
	rval := reflect.ValueOf(val)

	writer, err := makeWriter(rval.Type())
	if err != nil {
		return err
	}

	writer(rval, buf)

	return nil
}

func (buf *EncBuffer) writeTo(w io.Writer) error {
	_, err := w.Write(buf.str)
	if err != nil {
		return err
	}

	return nil
}

func (buf *EncBuffer) writeUint64(w io.Writer) error {
	fmt.Println("int64")
	return nil
}

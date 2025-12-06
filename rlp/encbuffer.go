package rlp

import (
	"io"
	"reflect"
	"sync"
)

// set a globe sync.pool
// write different write funcs

type EncBuff struct {
	str []byte
}

var encBuffPool = sync.Pool{
	New: func() interface{} { return &EncBuff{str: make([]byte, 1024)} },
}

func getEncBuffer() *EncBuff {
	buf := encBuffPool.Get().(*EncBuff)
	buf.reset()
	return buf
}

func (buf *EncBuff) reset() {
	buf.str = buf.str[:0]
}

func (buf *EncBuff) encode(val interface{}) error {
	rval := reflect.ValueOf(val)

	writer, err := makeWriter(rval.Type())
	if err != nil {
		return err
	}

	writer(rval, buf)

	return nil
}

func (buf *EncBuff) writeTo(w io.Writer) error {
	_, err := w.Write(buf.str)
	if err != nil {
		return err
	}

	return nil
}

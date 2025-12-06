package rlp

import (
	"fmt"
	"io"
	"reflect"
)

type writer func(reflect.Value, *EncBuff) error

func Encode(w io.Writer, val interface{}) error {
	buf := getEncBuffer()
	if err := buf.encode(val); err != nil {
		return err
	}

	return buf.writeTo(w)
}

func makeWriter(typ reflect.Type) (writer, error) {
	kind := typ.Kind()

	switch {
	case kind == reflect.Uint64:
		return write64, nil
	default:
		return nil, fmt.Errorf("rlp: type %v is not RLP-serializable", typ)
	}
}

func write64(rval reflect.Value, buf *EncBuff) error {

}

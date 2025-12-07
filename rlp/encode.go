package rlp

import (
	"fmt"
	"io"
	"log"
	"reflect"
)

func Encode(w io.Writer, val interface{}) error {
	buf := getEncBufferer()
	if err := buf.encode(val); err != nil {
		return err
	}

	return buf.writeTo(w)
}

func makeWriter(typ reflect.Type) (writer, error) {
	kind := typ.Kind()

	switch {
	case kind == reflect.Struct:
		return makeStructWriter(typ)
	case kind == reflect.String:
		return writeString, nil
	case kind == reflect.Uint64:
		return writeUint64, nil
	default:
		return nil, fmt.Errorf("rlp: type %v is not RLP-serializable", typ)
	}
}

func writeString(rval reflect.Value, buf *EncBuffer) error {
	log.Println("string")
	return nil
}

func writeUint64(rval reflect.Value, buf *EncBuffer) error {
	log.Println("int64")
	return nil
}

func makeStructWriter(typ reflect.Type) (writer, error) {
	fields, err := structFields(typ)
	if err != nil {
		return nil, err
	}

	writer := func(rval reflect.Value, buf *EncBuffer) error {
		for _, f := range fields {
			if err := f.writer(rval.Field(f.index), buf); err != nil {
				return err
			}
		}

		return nil
	}

	return writer, nil
}

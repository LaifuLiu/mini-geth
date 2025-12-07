package rlp

import "reflect"

type writer func(reflect.Value, *EncBuffer) error

type Field struct {
	writer writer
	index  int
}

func structFields(typ reflect.Type) ([]Field, error) {
	fields := make([]Field, typ.NumField())

	for i := 0; i < typ.NumField(); i++ {
		writer, err := makeWriter(typ.Field(i).Type)
		if err != nil {
			return nil, err
		}
		fields = append(fields, Field{writer: writer, index: i})
	}

	return fields, nil
}

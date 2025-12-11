package ethdb

type KeyValueReader interface {
	Has(key []byte) (bool, error)

	Get(key []byte) ([]byte, error)
}

type KeyValueWriter interface {
	Put(key []byte, values []byte) error

	Delete(key []byte) error
}

type KeyValueRangeDeleter interface {
	DeleteRange(start, end []byte) error
}

type KeyValueStore interface {
	KeyValueReader
	KeyValueWriter
	KeyValueRangeDeleter
}

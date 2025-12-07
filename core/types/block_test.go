package types

import (
	"bytes"
	"testing"
)

func TestEncodeRlp(t *testing.T) {
	b := new(bytes.Buffer)
	h := &Header{
		parentHash: "abcde",
		basefee:    10,
	}
	tx := NewTransaction("eee", "alice", "bob")
	txs := []*Transaction{tx}
	block := NewBlock(h, txs)
	block.EncodeRlp(b)
}

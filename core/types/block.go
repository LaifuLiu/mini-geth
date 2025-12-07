package types

import (
	"io"

	"github.com/LaifuLiu/mini-geth/rlp"
)

type Header struct {
	parentHash string
	// txhash      string
	// receipthash string
	// number      uint64
	// gaslimit    uint64
	// gasused     uint64
	basefee uint64
}

type Block struct {
	header *Header
	txs    []*Transaction
}

type ExtBlock struct {
	header *Header
	txs    []*Transaction
}

func NewBlock(header *Header, txs []*Transaction) *Block {
	return &Block{
		header: header,
		txs:    txs,
	}
}

func (b *Block) EncodeRlp(w io.Writer) error {
	return rlp.Encode(w, &ExtBlock{
		header: b.header,
		txs:    b.txs,
	})
}

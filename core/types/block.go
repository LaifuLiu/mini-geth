package types

import "io"

type Header struct {
	parentHash  string
	txHash      string
	receiptHash string
	number      uint64
	gasLimit    uint64
	gasUsed     uint64
	baseFee     uint64
}

type Block struct {
	header *Header
	txs    []*Transaction
}

type ExtBlock struct {
	header *Header
	txs    []*Transaction
}

func (b *Block) EncodeRlp(w io.Writer) error {
	return rlp.Encode(w, &ExtBlock{
		header: b.header,
		txs:    b.txs,
	})
}

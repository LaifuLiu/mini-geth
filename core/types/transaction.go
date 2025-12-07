package types

type Transaction struct {
	hash string
	from string
	to   string
}

func NewTransaction(hash, from, to string) *Transaction {
	return &Transaction{
		hash: hash,
		from: from,
		to:   to,
	}
}

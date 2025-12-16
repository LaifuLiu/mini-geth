package core

import (
	"github.com/LaifuLiu/mini-geth/core/types"
)

type BlockChain struct{}

func NewBlockChain() (*BlockChain, error) {
	return nil, nil
}

func (bc *BlockChain) InsertChain(chain types.Block) error {
	return nil
}

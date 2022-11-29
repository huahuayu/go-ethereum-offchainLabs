package arbitrum

import (
	"context"

	"github.com/huahuayu/go-ethereum-offchainLabs/core"
	"github.com/huahuayu/go-ethereum-offchainLabs/core/types"
)

type ArbInterface interface {
	PublishTransaction(ctx context.Context, tx *types.Transaction) error
	BlockChain() *core.BlockChain
	ArbNode() interface{}
}

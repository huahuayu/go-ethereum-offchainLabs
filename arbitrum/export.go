package arbitrum

import (
	"context"

	"github.com/huahuayu/go-ethereum-offchainLabs/common/hexutil"
	"github.com/huahuayu/go-ethereum-offchainLabs/core"
	"github.com/huahuayu/go-ethereum-offchainLabs/internal/ethapi"
	"github.com/huahuayu/go-ethereum-offchainLabs/rpc"
)

type TransactionArgs = ethapi.TransactionArgs

func EstimateGas(ctx context.Context, b ethapi.Backend, args TransactionArgs, blockNrOrHash rpc.BlockNumberOrHash, gasCap uint64) (hexutil.Uint64, error) {
	return ethapi.DoEstimateGas(ctx, b, args, blockNrOrHash, gasCap)
}

func NewRevertReason(result *core.ExecutionResult) error {
	return ethapi.NewRevertError(result)
}

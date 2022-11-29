package eth

import (
	"github.com/huahuayu/go-ethereum-offchainLabs/core"
	"github.com/huahuayu/go-ethereum-offchainLabs/core/state"
	"github.com/huahuayu/go-ethereum-offchainLabs/core/types"
	"github.com/huahuayu/go-ethereum-offchainLabs/core/vm"
	"github.com/huahuayu/go-ethereum-offchainLabs/ethdb"
)

func NewArbEthereum(
	blockchain *core.BlockChain,
	chainDb ethdb.Database,
) *Ethereum {
	return &Ethereum{
		blockchain: blockchain,
		chainDb:    chainDb,
	}
}

func (eth *Ethereum) StateAtTransaction(block *types.Block, txIndex int, reexec uint64) (core.Message, vm.BlockContext, *state.StateDB, error) {
	return eth.stateAtTransaction(block, txIndex, reexec)
}

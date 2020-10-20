// Package supervisor manages running modules and takes care
// of the inter-modules communication, if required.
package supervisor

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
)

// Lachesis returns a lachesis full node connection if available.
func (ws *WatchdogSupervisor) Lachesis() *rpc.Client {
	// try to establish the lachesis connection
	if ws.client == nil {
		ws.dialLachesis()
	}

	return ws.client
}

// Lachesis returns a lachesis full node connection if available.
func (ws *WatchdogSupervisor) dialLachesis() {
	// do we have a connection?
	if ws.client != nil {
		return
	}

	// log the action
	ws.log.Debugf("connecting lachesis at %s", ws.nodeRpcPath)

	// establish connection
	client, err := rpc.Dial(ws.nodeRpcPath)
	if err != nil {
		ws.log.Criticalf("can not connect lachesis node; %s", err.Error())
		return
	}

	// notice connection
	ws.log.Noticef("lachesis connection established at %s", ws.nodeRpcPath)
	ws.client = client
}

// BlockHeight returns the current block height of the Opera blockchain.
func (ws *WatchdogSupervisor) BlockHeight() (*big.Int, error) {
	// call for data
	var height hexutil.Big
	err := ws.Lachesis().Call(&height, "ftm_blockNumber")
	if err != nil {
		ws.Log().Errorf("block height not available; %s", err.Error())
		return nil, err
	}

	// inform and return
	ws.Log().Debugf("current block height is %s", height.ToInt().String())
	return height.ToInt(), nil
}

// ContractCallOpts provides call options if possible
func (ws *WatchdogSupervisor) ContractCallOpts(from common.Address) *bind.CallOpts {
	// get the latest block available
	block, err := ws.BlockHeight()
	if err != nil {
		ws.Log().Errorf("call options not available; %s", err.Error())
		return nil
	}

	// prep the call
	call := bind.CallOpts{
		Pending:     false,
		From:        from,
		BlockNumber: block,
		Context:     nil,
	}

	return &call
}

// Package supervisor manages running modules and takes care
// of the inter-modules communication, if required.
package supervisor

import (
	"github.com/ethereum/go-ethereum/rpc"
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

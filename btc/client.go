package btc

import (
	"github.com/fanguanghui/coinrpc/rpc"
)



type BtcClient struct {
	ConnCfg *rpc.ConnConfig
}

func NewBtcClient(connCfg *rpc.ConnConfig) *BtcClient {
	return &BtcClient{connCfg}
}

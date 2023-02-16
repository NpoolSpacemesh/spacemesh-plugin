package client

import (
	"context"

	apitypes "github.com/spacemeshos/api/release/go/spacemesh/v1"
)

// GlobalStateHash returns the current global state hash
func (c *gRPCClient) GlobalStateHash() (*apitypes.GlobalStateHash, error) {
	gsc := c.getGlobalStateServiceClient()
	if resp, err := gsc.GlobalStateHash(context.Background(), &apitypes.GlobalStateHashRequest{}); err != nil {
		return nil, err
	} else {
		return resp.Response, nil
	}
}

// AccountState returns basic account data such as balance and nonce from the global state
func (c *gRPCClient) AccountState(address apitypes.AccountId) (*apitypes.Account, error) {
	gsc := c.getGlobalStateServiceClient()
	resp, err := gsc.Account(context.Background(), &apitypes.AccountRequest{
		AccountId: &address})
	if err != nil {
		return nil, err
	}

	return resp.AccountWrapper, nil
}

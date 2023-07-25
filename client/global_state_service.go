package client

import (
	"context"

	v1 "github.com/spacemeshos/api/release/go/spacemesh/v1"
)

// GlobalStateHash returns the current global state hash
func (c *Client) GlobalStateHash(ctx context.Context) (*v1.GlobalStateHash, error) {
	gsc := c.getGlobalStateServiceClient()
	if resp, err := gsc.GlobalStateHash(ctx, &v1.GlobalStateHashRequest{}); err != nil {
		return nil, err
	} else {
		return resp.Response, nil
	}
}

// AccountState returns basic account data such as balance and nonce from the global state
func (c *Client) AccountState(ctx context.Context, address v1.AccountId) (*v1.Account, error) {
	gsc := c.getGlobalStateServiceClient()
	resp, err := gsc.Account(ctx, &v1.AccountRequest{
		AccountId: &address})
	if err != nil {
		return nil, err
	}

	return resp.AccountWrapper, nil
}

package client

import (
	"context"

	apitypes "github.com/spacemeshos/api/release/go/spacemesh/v1"
)

// GetMeshTransactions returns the MeshTransactions on the mesh to or from an address.
func (c *Client) GetGenesisID(ctx context.Context) ([]byte, error) {
	ms := c.getMeshServiceClient()
	resp, err := ms.GenesisID(ctx, &apitypes.GenesisIDRequest{})
	if err != nil {
		return nil, err
	}
	return resp.GenesisId, err
}

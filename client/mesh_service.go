package client

import (
	"context"

	apitypes "github.com/spacemeshos/api/release/go/spacemesh/v1"
)

// GetMeshTransactions returns the MeshTransactions on the mesh to or from an address.
func (c *Client) GetGenesisID() ([]byte, error) {
	ms := c.getMeshServiceClient()
	resp, err := ms.GenesisID(context.Background(), &apitypes.GenesisIDRequest{})
	if err != nil {
		return nil, err
	}
	return resp.GenesisId, err
}

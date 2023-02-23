package client

import (
	"errors"

	"github.com/golang/protobuf/ptypes/empty"
	apitypes "github.com/spacemeshos/api/release/go/spacemesh/v1"
	"golang.org/x/net/context"
)

type NodeInfo struct {
	Version string
	Build   string
}

// Echo is a basic api sanity test. It verifies that the client can connect to
// the node service and get a response from it to an echo request.
// todo: change this to api health-check service as node service might not be available
func (c *Client) Echo() error {
	service := c.getNodeServiceClient()
	const msg = "hello spacemesh"
	resp, err := service.Echo(context.Background(), &apitypes.EchoRequest{
		Msg: &apitypes.SimpleString{Value: msg}})

	if err != nil {
		return err
	}

	if resp.Msg.Value != msg {
		return errors.New("unexpected node service echo response")
	}

	return nil
}

// NodeInfo returns static node info such as build, version and api server url
func (c *Client) NodeInfo() (*NodeInfo, error) {
	info := &NodeInfo{}
	s := c.getNodeServiceClient()
	resp, err := s.Version(context.Background(), &empty.Empty{})
	if err != nil {
		return nil, err
	}
	info.Version = resp.VersionString.Value

	resp1, err := s.Build(context.Background(), &empty.Empty{})
	if err != nil {
		return nil, err
	}
	info.Build = resp1.BuildString.Value

	return info, nil
}

// NodeStatus returns dynamic node status such as sync status and number of connected peers
func (c *Client) NodeStatus() (*apitypes.NodeStatus, error) {
	s := c.getNodeServiceClient()
	resp, err := s.Status(context.Background(), &apitypes.StatusRequest{})
	if err != nil {
		return nil, err
	}

	return resp.Status, nil
}

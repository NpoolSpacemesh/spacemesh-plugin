package main

import (

	// "crypto/ed25519"

	"context"
	"encoding/hex"
	"fmt"

	v1 "github.com/spacemeshos/api/release/go/spacemesh/v1"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/genvm/sdk"
	"github.com/spacemeshos/go-spacemesh/genvm/sdk/wallet"

	"github.com/NpoolSpacemesh/spacemesh-plugin/account"
	"github.com/NpoolSpacemesh/spacemesh-plugin/client"
	"github.com/NpoolSpacemesh/spacemesh-plugin/util"
	"github.com/spacemeshos/go-spacemesh/signing"
)

const (
	endpoint    = "172.16.3.90:9092"
	endpointNet = account.StandaloneHRP
)

var (
	fromPriStr = "e0aef1ce781f28ed8a88f5a3fd87e4ffe399f3750b3fee640e2606c5e2922217f7d6c5e814c89faf8a2866e65000f31253f5c42568ca3cf4f0c6e70cbc878559"
	// fromPubStr := "f7d6c5e814c89faf8a2866e65000f31253f5c42568ca3cf4f0c6e70cbc878559"
	fromAddressStestStr      = "standalone1qqqqqqru6pcet8crur4qw52ac73w3mamj0mk0vq0qxtem"
	fromAddressStandaloneStr = "stest1qqqqqqru6pcet8crur4qw52ac73w3mamj0mk0vqu5zg3w"

	toPriStr = "01c6f1db4dcbf1b900cb3ff0411a1cf4b1279f95bf33103a01996e1be1d7754f09815bea25c2c74754b3438e8217da3c6799e13e9e1cedba6e7f30e079269ca4"
	// toPubStr := "09815bea25c2c74754b3438e8217da3c6799e13e9e1cedba6e7f30e079269ca4"
	toAddressStandaloneStr = "standalone1qqqqqqx5xtqgy44pgeweeya02yszawqhc0w9ulqwuk2qv"
	toAddressStestStr      = "stest1qqqqqqx5xtqgy44pgeweeya02yszawqhc0w9ulqagjfge"

	ctx = context.Background()
)

func NewEdSignerByPriStr(priStr string) (*signing.EdSigner, error) {
	priBytes, err := hex.DecodeString(priStr)
	if err != nil {
		return nil, err
	}
	signer, err := signing.NewEdSigner(signing.WithPrivateKey(priBytes))
	if err != nil {
		return nil, err
	}
	return signer, err
}

func main() {
	// tttt()
	// justTx()
	// justTx()
	// time.Sleep(time.Second * 20)
	// justTx()
	// justTx()
	// getTransaction()
	// _txID, err := base64.StdEncoding.DecodeString("dtTHPZJo7mqdlCgoJ1x9j1Bk4o+V7qDpKbqqrJvydJQ=")
	// h32 := types.EmptyLayerHash
	// h32.SetBytes(_txID)
	// fmt.Println(h32.Hex(), err)

	// time.Sleep(2 * time.Second)
	justSpawn()
	// testAcc()
	// testAcc()
}

func tttt() {
	acc, err := account.CreateAccountFromHexPri(fromPriStr)
	fmt.Println(acc, err)
	fmt.Println(acc.GetAddress(endpointNet).String())

	acc, err = account.CreateAccountFromHexPri(toPriStr)
	fmt.Println(acc, err)
	fmt.Println(acc.GetAddress(endpointNet).String())
}

func testAcc() {
	client := client.NewClient(endpoint, false)
	err := client.Connect(ctx)
	fmt.Println(1, err)
	acc, err := client.AccountState(ctx, v1.AccountId{Address: fromAddressStandaloneStr})
	fmt.Println(2, acc, err)

}

func justSpawn() {
	client := client.NewClient(endpoint, false)
	err := client.Connect(ctx)
	fmt.Println(1, err)

	genesisID, err := client.GetGenesisID(ctx)
	fmt.Println(2, genesisID, err)

	signer, err := NewEdSignerByPriStr(toPriStr)
	fmt.Println(3, err)

	_genesisID := types.EmptyLayerHash
	_genesisID.SetBytes(genesisID)
	h20 := types.Hash20{}
	copy(h20[:], _genesisID[12:])

	rawTx := types.NewRawTx(wallet.SelfSpawn(signer.PrivateKey(), 0, sdk.WithGenesisID(h20), sdk.WithGasPrice(2)))

	txState, err := client.SubmitCoinTransaction(ctx, rawTx.Raw)
	fmt.Println(util.PrettyStruct(txState), err)
}

func justTx() {
	client := client.NewClient(endpoint, false)
	err := client.Connect(ctx)
	fmt.Println(1, err)

	// pepare to account
	toAcc, err := account.CreateAccountFromHexPri(toPriStr)
	fmt.Println(2, toAcc.GetAddress(endpointNet), err)

	toAddr := toAcc.GetAddress(endpointNet)
	// pepare from account
	fromSigner, err := NewEdSignerByPriStr(fromPriStr)
	fmt.Println(4, err)

	// pepare genessisID
	genesisID, err := client.GetGenesisID(ctx)
	fmt.Println(5, genesisID, err)
	_genesisID := types.EmptyLayerHash
	_genesisID.SetBytes(genesisID)
	h20 := types.Hash20{}
	copy(h20[:], _genesisID[12:])

	// get rawTX
	rawTx := types.NewRawTx(wallet.Spend(fromSigner.PrivateKey(), toAddr, 300, 2, sdk.WithGenesisID(h20), sdk.WithGasPrice(2)))

	// submit tx
	txState, err := client.SubmitCoinTransaction(ctx, rawTx.Raw)
	fmt.Println(6, util.PrettyStruct(txState), err)

	// get tx state
	txState, tx, err := client.TransactionState(ctx, txState.Id.GetId(), true)
	fmt.Println(7, util.PrettyStruct(txState))
	fmt.Println(8, util.PrettyStruct(tx))
	fmt.Println(9, err)
	hash32 := types.EmptyLayerHash
	hash32.SetBytes(txState.Id.GetId())
	fmt.Println(10, hash32.Hex())
}

func getTransaction() {
	client := client.NewClient(endpoint, false)
	err := client.Connect(ctx)
	fmt.Println(1, err)
	txID := "0x6d81b1871c09f73b8856475230a00c33cc8929f7c479e69c54c52f575f95f02e"
	_txID := types.HexToHash32(txID)

	// _txID, err := base64.StdEncoding.DecodeString("dtTHPZJo7mqdlCgoJ1x9j1Bk4o+V7qDpKbqqrJvydJQ=")
	// _txID, err := base64.StdEncoding.DecodeString("7+XMPKGhPuQ1gG8L+TraFSa0/XXMp1ln8j3VzkGwiDA=")
	fmt.Println(2.5, err)
	txState, tx, err := client.TransactionState(ctx, _txID[:], true)
	fmt.Println(2, err)
	// fmt.Println(base64.StdEncoding.EncodeToString(txState.Id.GetId()))

	// hhhash := types.EmptyLayerHash
	// hhhash.SetBytes(txState.Id.GetId())
	// fmt.Println(hhhash.String())
	fmt.Println(3, util.PrettyStruct(txState.State))
	fmt.Println(4, util.PrettyStruct(tx))

	// fmt.Println(client.AccountState(v1.AccountId{Address: "stest1qqqqqq8ahm02zlsyvl9n2era9h9jlhy9hmrkfnqjyfrj7"}))

}

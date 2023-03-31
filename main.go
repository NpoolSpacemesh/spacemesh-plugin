package main

import (
	"bytes"

	// "crypto/ed25519"

	"encoding/base64"
	"encoding/hex"
	"fmt"

	xdr "github.com/davecgh/go-xdr/xdr2"
	v1 "github.com/spacemeshos/api/release/go/spacemesh/v1"
	"github.com/spacemeshos/go-scale"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/genvm/core"
	"github.com/spacemeshos/go-spacemesh/genvm/sdk"
	"github.com/spacemeshos/go-spacemesh/genvm/sdk/wallet"

	tmpWallet "github.com/spacemeshos/go-spacemesh/genvm/templates/wallet"

	smAccount "github.com/NpoolSpacemesh/spacemesh-plugin/account"
	"github.com/NpoolSpacemesh/spacemesh-plugin/client"
	"github.com/NpoolSpacemesh/spacemesh-plugin/util"
	"github.com/spacemeshos/go-spacemesh/signing"
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

type InnerTransaction struct {
	AccountNonce uint64
	Recipient    types.Address
	// GasLimit     uint64
	Price  uint64
	Amount uint64
}
type InTransaction struct {
	InnerTransaction
	Signature [64]byte
}

func interfaceToBytes(i interface{}) ([]byte, error) {
	var w bytes.Buffer
	if _, err := xdr.Marshal(&w, &i); err != nil {
		return nil, err
	}
	return w.Bytes(), nil
}

func encode(fields ...scale.Encodable) []byte {
	buf := bytes.NewBuffer(nil)
	encoder := scale.NewEncoder(buf)
	for _, field := range fields {
		_, err := field.EncodeScale(encoder)
		if err != nil {
			panic(err)
		}
	}
	return buf.Bytes()
}

func byteArrToBody(txs []byte) (body []byte) {
	bodyStr := fmt.Sprint("{\"tx\":[")
	comma := ""
	for _, b := range txs {
		bodyStr += fmt.Sprint(comma)
		bodyStr += fmt.Sprint(b)
		comma = ","
	}
	bodyStr += fmt.Sprintln("]}")
	return []byte(bodyStr)
}

func main() {
	// getTransaction()
	// justTx()
	// time.Sleep(2 * time.Second)
	// justSpawn()
	// testAcc()
	tttt()
}

func tttt() {
	types.DefaultTestAddressConfig()
	types.DefaultAddressConfig().NetworkHRP = smAccount.MainNet

	public := &core.PublicKey{}

	fmt.Println(core.ComputePrincipal(tmpWallet.TemplateAddress, public).String())
	fmt.Println(types.GenerateAddress([]byte("ssss")).String())

}

func testAcc() {
	acc, _ := smAccount.CreateAccount(smAccount.TestNet)
	fmt.Println(util.PrettyStruct(acc))
}

func justSpawn() {
	types.DefaultTestAddressConfig()
	client := client.NewClient("172.16.3.90:9092", false)
	err := client.Connect()
	fmt.Println(1, err)

	genesisID, err := client.GetGenesisID()
	fmt.Println(2, genesisID, err)

	priStr := "e0aef1ce781f28ed8a88f5a3fd87e4ffe399f3750b3fee640e2606c5e2922217f7d6c5e814c89faf8a2866e65000f31253f5c42568ca3cf4f0c6e70cbc878559"
	// pubStr := "f7d6c5e814c89faf8a2866e65000f31253f5c42568ca3cf4f0c6e70cbc878559"
	// stestStr := "stest1qqqqqqru6pcet8crur4qw52ac73w3mamj0mk0vqu5zg3w"
	types.DefaultTestAddressConfig()

	signer, err := NewEdSignerByPriStr(priStr)
	fmt.Println(3, err)

	_genesisID := types.EmptyLayerHash
	_genesisID.SetBytes(genesisID)
	h20 := types.Hash20{}
	copy(h20[:], _genesisID[12:])

	rawTx := types.NewRawTx(wallet.SelfSpawn(signer.PrivateKey(), 0, sdk.WithGenesisID(h20), sdk.WithGasPrice(2)))

	txState, err := client.SubmitCoinTransaction(rawTx.Raw)
	fmt.Println(util.PrettyStruct(txState), err)
}

func justTx() {
	types.DefaultTestAddressConfig()
	client := client.NewClient("172.16.3.90:9092", false)
	err := client.Connect()
	fmt.Println(1, err)

	genesisID, err := client.GetGenesisID()
	fmt.Println(2, genesisID, err)

	priStr := "e0aef1ce781f28ed8a88f5a3fd87e4ffe399f3750b3fee640e2606c5e2922217f7d6c5e814c89faf8a2866e65000f31253f5c42568ca3cf4f0c6e70cbc878559"
	// pubStr := "f7d6c5e814c89faf8a2866e65000f31253f5c42568ca3cf4f0c6e70cbc878559"
	stestStr := "stest1qqqqqqru6pcet8crur4qw52ac73w3mamj0mk0vqu5zg3w"
	toStr := "stest1qqqqqqx5xtqgy44pgeweeya02yszawqhc0w9ulqagjfge"
	toAddr, err := types.StringToAddress(toStr)
	fmt.Println(2, toAddr, err)

	types.DefaultTestAddressConfig()

	signer, err := NewEdSignerByPriStr(priStr)

	fmt.Println(3, err)
	_genesisID := types.EmptyLayerHash
	_genesisID.SetBytes(genesisID)
	h20 := types.Hash20{}
	copy(h20[:], _genesisID[12:])

	acc, err := client.AccountState(v1.AccountId{Address: stestStr})
	fmt.Println(4, acc, err)

	rawTx := types.NewRawTx(wallet.Spend(signer.PrivateKey(), toAddr, 100, 1, sdk.WithGenesisID(h20), sdk.WithGasPrice(2)))

	txState, err := client.SubmitCoinTransaction(rawTx.Raw)
	fmt.Println(util.PrettyStruct(txState), err)
}

func getTransaction() {
	client := client.NewClient("172.16.3.90:9092", false)
	err := client.Connect()
	fmt.Println(1, err)
	// txID := "0x7eaa2bd72608d438210d3df3dfb4dd02ed5b0b74385098691b036f9e412faaca"
	// _txID := types.HexToHash32(txID)
	_txID, err := base64.StdEncoding.DecodeString("LCy3zlmBZoJSpme/OolSk0WF8+ypzP8I+EpbwHxQ1YE=")
	// _txID, err := base64.StdEncoding.DecodeString("IUdo9KDVmzf0JdOKsatfO6aEpcB99xmkdyT7NmG3iqM=")
	fmt.Println(2.5, err)
	txState, tx, err := client.TransactionState(_txID[:], true)
	fmt.Println(2, err)
	// fmt.Println(base64.StdEncoding.EncodeToString(txState.Id.GetId()))

	// hhhash := types.EmptyLayerHash
	// hhhash.SetBytes(txState.Id.GetId())
	// fmt.Println(hhhash.String())
	fmt.Println(3, util.PrettyStruct(txState))
	fmt.Println(4, util.PrettyStruct(tx))

	// fmt.Println(client.AccountState(v1.AccountId{Address: "stest1qqqqqq8ahm02zlsyvl9n2era9h9jlhy9hmrkfnqjyfrj7"}))

}

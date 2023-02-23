package main

import (
	"bytes"
	// "crypto/ed25519"
	"encoding/hex"
	"fmt"

	xdr "github.com/davecgh/go-xdr/xdr2"
	"github.com/spacemeshos/ed25519"
	"github.com/spacemeshos/go-scale"
	"github.com/spacemeshos/go-spacemesh/codec"
	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/genvm/core"
	"github.com/spacemeshos/go-spacemesh/genvm/sdk"
	"github.com/spacemeshos/go-spacemesh/genvm/sdk/wallet"
	tplWallet "github.com/spacemeshos/go-spacemesh/genvm/templates/wallet"

	"github.com/NpoolSpacemesh/spacemesh-plugin/client"
	"github.com/NpoolSpacemesh/spacemesh-plugin/util"
	"github.com/spacemeshos/go-spacemesh/hash"
	"github.com/spacemeshos/go-spacemesh/signing"
)

func NewAccount() (pri, pub, acc string) {
	signer, err := signing.NewEdSigner()
	if err != nil {
		fmt.Println(err)
	}

	// just use in testnet
	types.DefaultTestAddressConfig()

	pubStr := signer.PublicKey().String()
	priStr := hex.EncodeToString(signer.PrivateKey())
	accStr := types.GenerateAddress([]byte(pubStr)).String()

	return priStr, pubStr, accStr
}

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
	spawn()
}

func spawn() {
	// account()
	priStr := "eadcdac9cf65a2d5e7899ab2090dbb216ce06653255a42075d8f5c5c147e591407b46e6f13edb93757be7df38c72b945d5f7904f31aaff6737b84b2a8a2a58ee"
	// pubStr := "07b46e6f13edb93757be7df38c72b945d5f7904f31aaff6737b84b2a8a2a58ee"
	// stestStr := "stest1qqqqqqrxvcmrwvehvgurgc3jvyuxzvnpx5ux2eghq7crd"
	types.DefaultTestAddressConfig()

	signer, err := NewEdSignerByPriStr(priStr)
	fmt.Println(1, err)

	// pubStr := signer.PublicKey().String()
	// stestStr := types.GenerateAddress([]byte(pubStr)).String()
	// fmt.Println(priStr)
	// fmt.Println(pubStr)
	// fmt.Println(stestStr)

	client := client.NewClient("172.16.3.90:9092", false)
	err = client.Connect()

	fmt.Println(2, err)
	// fmt.Println(client.AccountState(v1.AccountId{Address: "stest1qqqqqqq28n6fw97jclu3tna6syxxy4elga2jtqgrf94zd"}))

	tx := &types.Transaction{TxHeader: &types.TxHeader{}}
	tx.Principal = wallet.Address(signer.PublicKey().Bytes())
	args := tplWallet.SpawnArguments{}
	copy(args.PublicKey[:], signer.PublicKey().Bytes())
	payload := core.Payload{}
	payload.Nonce = 0
	payload.GasPrice = 0
	public := &core.PublicKey{}
	copy(public[:], signer.PublicKey().Bytes())
	principal := core.ComputePrincipal(tplWallet.TemplateAddress, public)
	_tx := encode(&sdk.TxVersion, &principal, &sdk.MethodSpawn, &tplWallet.TemplateAddress, &payload, &args)

	genesisID, err := client.GetGenesisID()
	hh := hash.Sum(genesisID[:], _tx)
	sig := ed25519.Sign(signer.PrivateKey(), hh[:])
	tx.RawTx = types.NewRawTx(append(_tx, sig...))
	// serializedTx, err := codec.Encode(tx)
	txState, err := client.SubmitCoinTransaction(tx.Raw)

	fmt.Println(util.PrettyStruct(txState), err)
}

func spend() {
	// account()
	priStr := "eadcdac9cf65a2d5e7899ab2090dbb216ce06653255a42075d8f5c5c147e591407b46e6f13edb93757be7df38c72b945d5f7904f31aaff6737b84b2a8a2a58ee"
	// pubStr := "07b46e6f13edb93757be7df38c72b945d5f7904f31aaff6737b84b2a8a2a58ee"
	// stestStr := "stest1qqqqqqrxvcmrwvehvgurgc3jvyuxzvnpx5ux2eghq7crd"
	types.DefaultTestAddressConfig()

	signer, err := NewEdSignerByPriStr(priStr)
	fmt.Println(1, err)

	// pubStr := signer.PublicKey().String()
	// stestStr := types.GenerateAddress([]byte(pubStr)).String()
	// fmt.Println(priStr)
	// fmt.Println(pubStr)
	// fmt.Println(stestStr)

	client := client.NewClient("172.16.3.90:9092", false)
	err = client.Connect()

	fmt.Println(2, err)
	// fmt.Println(client.AccountState(v1.AccountId{Address: "stest1qqqqqqq28n6fw97jclu3tna6syxxy4elga2jtqgrf94zd"}))

	to := "stest1qqqqqqxdwamxnf80v5qsfgkntjqgx5auxxteu7scd4txz"
	toAddr, err := types.StringToAddress(to)
	fmt.Println(3, err)

	spawnargs := tplWallet.SpawnArguments{}
	copy(spawnargs.PublicKey[:], signer.PublicKey().PublicKey)
	principal := core.ComputePrincipal(tplWallet.TemplateAddress, &spawnargs)

	payload := core.Payload{}
	payload.GasPrice = 1
	payload.Nonce = 0

	args := tplWallet.SpendArguments{}
	args.Destination = toAddr
	args.Amount = 1000000

	genesisID, err := client.GetGenesisID()
	tx := &types.Transaction{TxHeader: &types.TxHeader{}}
	_tx := encode(&sdk.TxVersion, &principal, &sdk.MethodSpend, &payload, &args)
	hh := hash.Sum(genesisID[:], _tx)
	sig := ed25519.Sign(signer.PrivateKey(), hh[:])
	_tx = append(_tx, sig...)
	tx.RawTx = types.NewRawTx(_tx)
	tx.MaxSpend = 1

	serializedTx, err := codec.Encode(tx)
	txState, err := client.SubmitCoinTransaction(serializedTx)

	fmt.Println(util.PrettyStruct(txState), err)
}

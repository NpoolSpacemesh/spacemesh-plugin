package account

import (
	"crypto/ed25519"
	"encoding/hex"
	"errors"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/genvm/templates/wallet"

	"github.com/spacemeshos/go-spacemesh/genvm/core"
	"github.com/spacemeshos/go-spacemesh/signing"
)

const (
	MainHRP       = "sm"
	TestHRP       = "stest"
	StandaloneHRP = "standalone"
)

type Account struct {
	Pri    string
	Pub    string
	signer *signing.EdSigner
}

func CreateAccount() (*Account, error) {
	signer, err := signing.NewEdSigner()
	if err != nil {
		return nil, err
	}
	return createAccount(signer)
}

func PubkeyToAddress(pubkey []byte, hrp string) types.Address {
	types.SetNetworkHRP(hrp)
	key := [ed25519.PublicKeySize]byte{}
	copy(key[:], pubkey)
	walletArgs := &wallet.SpawnArguments{PublicKey: key}
	return core.ComputePrincipal(wallet.TemplateAddress, walletArgs)
}

func CreateAccountFromHexPri(pri string) (*Account, error) {
	priBytes, err := hex.DecodeString(pri)
	if err != nil {
		return nil, err
	}
	signer, err := signing.NewEdSigner(signing.WithPrivateKey(priBytes))
	if err != nil {
		return nil, err
	}
	return createAccount(signer)
}

func createAccount(signer *signing.EdSigner) (*Account, error) {
	if signer == nil {
		return nil, errors.New("edsigner is nil")
	}
	public := &core.PublicKey{}
	copy(public[:], signing.Public(signer.PrivateKey()))

	return &Account{Pri: hex.EncodeToString(signer.PrivateKey()),
		Pub:    hex.EncodeToString(signer.PublicKey().Bytes()),
		signer: signer,
	}, nil
}

func (a *Account) GetSigner() *signing.EdSigner {
	return a.signer
}

func (a *Account) GetAddress(hrp string) types.Address {
	return PubkeyToAddress(a.signer.PublicKey().PublicKey, hrp)
}

package account

import (
	"encoding/hex"
	"errors"

	tmpWallet "github.com/spacemeshos/go-spacemesh/genvm/templates/wallet"

	"github.com/spacemeshos/go-spacemesh/common/types"
	"github.com/spacemeshos/go-spacemesh/genvm/core"
	"github.com/spacemeshos/go-spacemesh/signing"
)

const (
	MainNet = "sm"
	TestNet = "stest"
)

type Account struct {
	Pri       string
	Pub       string
	Principal string
	signer    *signing.EdSigner
}

func CreateAccount(net string) (*Account, error) {

	signer, err := signing.NewEdSigner()
	if err != nil {
		return nil, err
	}
	return createAccount(signer, net)
}

func CreateAccountFromHexPri(pri string, net string) (*Account, error) {
	priBytes, err := hex.DecodeString(pri)
	if err != nil {
		return nil, err
	}
	signer, err := signing.NewEdSigner(signing.WithPrivateKey(priBytes))
	if err != nil {
		return nil, err
	}
	return createAccount(signer, net)
}

func createAccount(signer *signing.EdSigner, net string) (*Account, error) {
	if signer == nil {
		return nil, errors.New("edsigner is nil")
	}

	types.DefaultAddressConfig().NetworkHRP = net
	public := &core.PublicKey{}
	copy(public[:], signing.Public(signer.PrivateKey()))
	principal := core.ComputePrincipal(tmpWallet.TemplateAddress, public)

	return &Account{Pri: hex.EncodeToString(signer.PrivateKey()),
		Pub:       hex.EncodeToString(signer.PublicKey().Bytes()),
		Principal: principal.String(),
		signer:    signer,
	}, nil
}

func (a *Account) GetSigner() *signing.EdSigner {
	return a.signer
}

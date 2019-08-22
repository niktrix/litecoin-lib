package account

import (
	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
)

type Account struct {
	Address    btcutil.Address
	PrivateKey *btcec.PrivateKey
}

//NewAccount Creates new account
func NewAccount(privateKey string, chainConfig *chaincfg.Params, isCompressed bool) (*Account, error) {
	// decode WIF
	decodedWif, err := btcutil.DecodeWIF(privateKey)
	if err != nil {
		return nil, err
	}

	decodedWif.PrivKey.PubKey().Params()

	serialisedPubKey := decodedWif.PrivKey.PubKey().SerializeUncompressed()
	if isCompressed {
		serialisedPubKey = decodedWif.PrivKey.PubKey().SerializeCompressed()
	}
	addresspubkey, err := btcutil.NewAddressPubKey(serialisedPubKey, chainConfig)
	if err != nil {
		return nil, err
	}

	fromAddress, err := btcutil.DecodeAddress(addresspubkey.EncodeAddress(), chainConfig)
	if err != nil {
		return nil, err
	}

	return &Account{Address: fromAddress, PrivateKey: decodedWif.PrivKey}, nil

}

package account

import (
	"github.com/ltcsuite/ltcd/chaincfg"
	"github.com/ltcsuite/ltcutil"
	"github.com/ltcsuite/ltcd/btcec"
	"github.com/ltcsuite/ltcutil/hdkeychain"
)

type Account struct {
	Address ltcutil.Address
	 PrivateKey *btcec.PrivateKey
}

//NewAccount Creates new account
func NewAccount(privateKey string, chainConfig *chaincfg.Params, isCompressed bool) (*Account, error) {
	// decode WIF
	decodedWif, err := ltcutil.DecodeWIF(privateKey)
	if err != nil {
		return nil, err
	}

	decodedWif.PrivKey.PubKey().Params()

	serialisedPubKey := decodedWif.PrivKey.PubKey().SerializeUncompressed()
	if isCompressed {
		serialisedPubKey = decodedWif.PrivKey.PubKey().SerializeCompressed()
	}
	addresspubkey, err := ltcutil.NewAddressPubKey(serialisedPubKey, chainConfig)
	if err != nil {
		return nil, err
	}

	fromAddress, err := ltcutil.DecodeAddress(addresspubkey.EncodeAddress(), chainConfig)
	if err != nil {
		return nil, err
	}

	return &Account{Address: fromAddress, PrivateKey: decodedWif.PrivKey}, nil

}

func NewAccountFromSeed(seed string, chainConfig *chaincfg.Params) (*Account, error) {
	// decode WIF

	key, err := hdkeychain.NewMaster([]byte("Hello Hello Hello"), chainConfig)
	if err != nil {
		return nil, err
	}

	// amount := int64(32000000000)
	// txFee := int64(1000)

	// Show that the generated master node extended key is private.
	address, _ := key.Address(chainConfig)

	return &Account{Address: address}, nil

}

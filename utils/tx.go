package utils

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/btcsuite/btcd/txscript"
	"github.com/btcsuite/btcd/wire"
	"github.com/btcsuite/btcutil"
	"github.com/niktrix/bitcoin-lib/account"

	"bytes"

	"encoding/hex"
)

type BTCTransaction struct {
	unspentTX   []UTXO
	amount      int64
	fee         int64
	from        *account.Account
	to          string
	chainConfig *chaincfg.Params
	txOut       *wire.MsgTx
	isCompress  bool
}

func NewTx() *BTCTransaction {
	return &BTCTransaction{isCompress: false}
}

func (tx *BTCTransaction) SetUnspentTxs(utxo []UTXO) *BTCTransaction {
	tx.unspentTX = utxo
	return tx
}

func (tx *BTCTransaction) SetIsCompress(is bool) *BTCTransaction {
	tx.isCompress = is
	return tx
}

func (tx *BTCTransaction) SetAmount(amount int64) *BTCTransaction {
	tx.amount = amount
	return tx
}

func (tx *BTCTransaction) SetFrom(from *account.Account) *BTCTransaction {
	tx.from = from
	return tx
}

func (tx *BTCTransaction) SetTo(to string) *BTCTransaction {
	tx.to = to
	return tx
}

func (tx *BTCTransaction) SetFee(fee int64) *BTCTransaction {
	tx.fee = fee
	return tx
}

func (tx *BTCTransaction) SetConfig(chainConfig *chaincfg.Params) *BTCTransaction {
	tx.chainConfig = chainConfig
	return tx
}

func (tx *BTCTransaction) Execute() error {
	// GET first UTXO
	firstUnspentTx := tx.unspentTX[0]
	sourceUTXOHash, err := chainhash.NewHashFromStr(firstUnspentTx.Txid)
	if err != nil {
		return err
	}

	sourceUTXO := wire.NewOutPoint(sourceUTXOHash, firstUnspentTx.Vout)
	//calculate left amout to get that utxo as out
	leftAmount := firstUnspentTx.Satoshis - (tx.amount + tx.fee)

	sourceTxIn := wire.NewTxIn(sourceUTXO, nil, nil)

	destinationAddress, err := btcutil.DecodeAddress(tx.to, tx.chainConfig)
	if err != nil {
		return err
	}

	destinationPkScript, err := txscript.PayToAddrScript(destinationAddress)
	if err != nil {
		return err
	}

	sourcePkScript, err := txscript.PayToAddrScript(tx.from.Address)
	if err != nil {
		return err
	}

	sourceTxOut := wire.NewTxOut(tx.amount, sourcePkScript)

	redeemTx := wire.NewMsgTx(wire.TxVersion)
	redeemTx.AddTxIn(sourceTxIn)
	leftAmout := wire.NewTxOut(leftAmount, sourcePkScript)
	redeemTx.AddTxOut(leftAmout)
	redeemTxOut := wire.NewTxOut((tx.amount - tx.fee), destinationPkScript)
	redeemTx.AddTxOut(redeemTxOut)

	sigScript, err := txscript.SignatureScript(redeemTx, 0, sourceTxOut.PkScript, txscript.SigHashAll, tx.from.PrivateKey, tx.isCompress)
	if err != nil {
		return err
	}

	redeemTx.TxIn[0].SignatureScript = sigScript

	// validate signature
	flags := txscript.StandardVerifyFlags
	vm, err := txscript.NewEngine(sourceTxOut.PkScript, redeemTx, 0, flags, nil, nil, tx.amount)
	if err != nil {
		return err
	}

	if err := vm.Execute(); err != nil {
		return err
	}
	tx.txOut = redeemTx

	return nil
}

func (tx *BTCTransaction) GetRaw() string {
	buf := bytes.NewBuffer(make([]byte, 0, tx.txOut.SerializeSize()))
	tx.txOut.Serialize(buf)

	rawTx := hex.EncodeToString(buf.Bytes())

	return rawTx
}

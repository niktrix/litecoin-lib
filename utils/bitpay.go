package utils

import (
	"encoding/json"

	"github.com/niktrix/bitcoin-lib/request"
)

const TEST_URL = "https://test-insight.bitpay.com"
const MAINNET_URL = "https://insight.bitpay.com"

type BitPay struct {
	chain string
	url   string
}

func NewBitPay(chain string) *BitPay {
	if chain == "mainnet" {
		return &BitPay{chain: chain, url: MAINNET_URL}
	}
	return &BitPay{chain: chain, url: TEST_URL}

}

func (bp *BitPay) GetUnspentTxs(address string) (ut []UTXO, err error) {
	url := bp.url + "/api/addrs/utxo"
	utxo, err := request.New().SetURL(url).SetRequestType("POST").SetBody("{\n\t\"addrs\":\"" + address + "\"\n}").Execute()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(utxo, &ut)
	if err != nil {
		return nil, err
	}
	return

}

func (bp *BitPay) BroadCastTX(rawtx string) (resp string, err error) {
	url := bp.url + "/api/tx/send"
	response, err := request.New().SetURL(url).SetRequestType("POST").SetBody("{\n\t\"rawtx\":\"" + rawtx + "\"\n}").Execute()
	if err != nil {
		return "", err
	}
	resp = string(response)

	return

}

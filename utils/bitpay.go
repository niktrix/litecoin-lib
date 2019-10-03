package utils

import (
	"encoding/json"
	"fmt"

	"github.com/niktrix/litecoin-lib/request"
)

const TEST_URL = "https://testnet.litecore.io/api/addr"
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
	url := bp.url + "/" + address + "/utxo"
	fmt.Println(url)
	utxo, err := request.New().SetURL(url).SetRequestType("GET").Execute()
	if err != nil {
		return ut, err
	}
	//fmt.Println(string(utxo))
	err = json.Unmarshal(utxo, &ut)
	if err != nil {
		return ut, err
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

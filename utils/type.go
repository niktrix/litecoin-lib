package utils

type UTXO struct {
	Address       string  `json:"address"`
	Txid          string  `json:"txid"`
	Vout          uint32     `json:"vout"`
	ScriptPubKey  string  `json:"scriptPubKey"`
	Amount        float64 `json:"amount"`
	Satoshis      int64     `json:"satoshis"`
	Height        int     `json:"height"`
	Confirmations int     `json:"confirmations"`
}

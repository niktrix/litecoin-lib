package utils

type BTCUtils interface {
	GetUnspentTxs(string) []UTXO
}

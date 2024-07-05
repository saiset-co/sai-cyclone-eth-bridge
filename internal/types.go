package internal

import "math/big"

type TxNotificationRequest struct {
	From string      `json:"from"`
	Tx   interface{} `json:"tx"`
}

type NewCycloneTx struct {
	Message string
}

type EthereumParameter struct {
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

type NewEthereumTx struct {
	Contract string              `json:"contract"`
	Method   string              `json:"method"`
	Value    string              `json:"value"`
	Params   []EthereumParameter `json:"params"`
}

type SendTxRequest struct {
	Method string      `json:"method"`
	Data   interface{} `json:"data"`
}

type VmResponse struct {
	C []map[string]interface{} `json:"C"`
	D map[string]interface{}   `json:"D"`
	R map[string]interface{}   `json:"R"`
	T map[string]interface{}   `json:"T"`
	V map[string]interface{}   `json:"V"`
}

type CycloneTx struct {
	Hash      string `json:"hash"`
	Block     string `json:"block"`
	Nonce     string `json:"nonce"`
	Vm        string `json:"vm"`
	Sender    string `json:"sender"`
	Signature string `json:"signature"`
	Message   string `json:"message"`
	Exec      struct {
		Hash       string     `json:"hash"`
		VmResponse VmResponse `json:"vmResponse"`
	} `json:"exec"`
	FeeCurrency string `json:"feeCurrency"`
}

type EthereumTx struct {
	Number    int                      `json:"Number"`
	Hash      string                   `json:"NumHash"`
	From      string                   `json:"From"`
	To        string                   `json:"To"`
	Amount    big.Int                  `json:"Amount"`
	Events    []map[string]interface{} `json:"Events"`
	Status    string                   `json:"Status"`
	Operation string                   `json:"Operation"`
	Input     string                   `json:"Input"`
}

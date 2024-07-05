package internal

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"

	"github.com/saiset-co/sai-cyclone-eth-bridge/logger"
	"github.com/saiset-co/sai-cyclone-eth-bridge/utils"
	saiService "github.com/saiset-co/sai-service/service"
)

func (is *InternalService) NewHandler() saiService.Handler {
	return saiService.Handler{
		"notify": saiService.HandlerElement{
			Name:        "Get transaction",
			Description: "Get transaction from the indexer",
			Function: func(data, meta interface{}) (interface{}, int, error) {
				return is.handleTransaction(data)
			},
		},
	}
}

func (is *InternalService) handleTransaction(data interface{}) (interface{}, int, error) {
	var request = new(TxNotificationRequest)

	dataJson, err := json.Marshal(data)
	if err != nil {
		logger.Logger.Error("handleTransaction", zap.Error(err))
		return nil, 500, err
	}

	err = json.Unmarshal(dataJson, request)
	if err != nil {
		logger.Logger.Error("handleTransaction", zap.Error(err))
		return nil, 500, err
	}

	err = validator.New().Struct(request)
	if err != nil {
		logger.Logger.Error("handleTransaction", zap.Error(err))
		return nil, 500, err
	}

	switch request.From {
	case "Cyclone":
		err = is.callEthContract(request.Tx)
		if err != nil {
			logger.Logger.Error("handleTransaction", zap.Error(err))
			return nil, 500, err
		}
	case "Ethereum":
		err = is.callCycloneContract(request.Tx)
		if err != nil {
			logger.Logger.Error("handleTransaction", zap.Error(err))
			return nil, 500, err
		}
	}

	return data, 200, nil
}

func (is *InternalService) callEthContract(txData interface{}) error {
	var tx = new(CycloneTx)

	txJson, err := json.Marshal(txData)
	if err != nil {
		logger.Logger.Error("callEthContract", zap.Error(err))
		return err
	}

	err = json.Unmarshal(txJson, tx)
	if err != nil {
		logger.Logger.Error("callEthContract", zap.Error(err))
		return err
	}

	for _, mapC := range tx.Exec.VmResponse.C {
		for _, data := range mapC {
			dataMap, ok := data.(map[string]interface{})
			if !ok {
				//logger
				continue
			}
			for _, dataMapValue := range dataMap {
				dataMapValueStr, ok := dataMapValue.(string)
				if !ok {
					//logger
					continue
				}

				str := strings.Split(dataMapValueStr, "-")

				if len(str) < 4 {
					//logger
					continue
				}

				newTxRequest := SendTxRequest{
					Method: "api",
					Data: NewEthereumTx{
						Contract: "WCYCL",
						Method:   "recordData",
						Value:    "0",
						Params: []EthereumParameter{
							{Type: "address", Value: str[1]},
							{Type: "string", Value: str[2]},
							{Type: "uint256", Value: str[3]},
						},
					},
				}

				payload, err := jsoniter.Marshal(&newTxRequest)
				if err != nil {
					logger.Logger.Error("Notify", zap.Error(err))
					return err
				}

				_, err = utils.SaiQuerySender(bytes.NewReader(payload), is.ethereum.Interaction, is.ethereum.Token)
				if err != nil {
					logger.Logger.Error("Notify", zap.Error(err))
					continue
				}
			}
		}
	}

	return err
}

func (is *InternalService) callCycloneContract(txData interface{}) error {
	var tx = new(EthereumTx)

	txJson, err := json.Marshal(txData)
	if err != nil {
		logger.Logger.Error("callEthContract", zap.Error(err))
		return err
	}

	err = json.Unmarshal(txJson, tx)
	if err != nil {
		logger.Logger.Error("callEthContract", zap.Error(err))
		return err
	}

	var sb strings.Builder
	sb.WriteString("callContract('" + is.cyclone.Contract + "'")
	sb.WriteString(",")
	sb.WriteString("recordData")
	sb.WriteString(",")
	sb.WriteString(tx.From)
	//sb.WriteString(",")
	//sb.WriteString(tx.Events)

	newTxRequest := SendTxRequest{
		Method: "send_tx",
		Data: NewCycloneTx{
			Message: sb.String(),
		},
	}

	payload, err := jsoniter.Marshal(&newTxRequest)
	if err != nil {
		logger.Logger.Error("Notify", zap.Error(err))
		return err
	}

	_, err = utils.SaiQuerySender(bytes.NewReader(payload), is.cyclone.Interaction, is.cyclone.Token)

	return err
}

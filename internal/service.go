package internal

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	"github.com/saiset-co/sai-cyclone-eth-bridge/logger"
	saiService "github.com/saiset-co/sai-service/service"
)

type Cyclone struct {
	Interaction string
	Contract    string
	Token       string
}

type Ethereum struct {
	Interaction string
	Contract    string
	Token       string
}

type InternalService struct {
	Context  *saiService.Context
	client   http.Client
	cyclone  Cyclone
	ethereum Ethereum
}

func (is *InternalService) Init() {
	cyclone := is.Context.GetConfig("cyclone", Cyclone{})
	ethereum := is.Context.GetConfig("ethereum", Ethereum{})

	cycloneBytes, err := json.Marshal(cyclone)
	if err != nil {
		logger.Logger.Error("Init", zap.Error(err))
		panic(err)
	}

	err = json.Unmarshal(cycloneBytes, &is.cyclone)
	if err != nil {
		logger.Logger.Error("Init", zap.Error(err))
		panic(err)
	}

	ethereumBytes, err := json.Marshal(ethereum)
	if err != nil {
		logger.Logger.Error("Init", zap.Error(err))
		panic(err)
	}

	err = json.Unmarshal(ethereumBytes, &is.ethereum)
	if err != nil {
		logger.Logger.Error("Init", zap.Error(err))
		panic(err)
	}
}

func (is *InternalService) Process() {

}

package kcckit

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
)

// An AppService provides an HTTP client and a signer to make a HTTP request with the signature to KuCoin API.
type AppService struct {
	rpcClient  *ethclient.Client
	ChainId    string
	rpcUrl     string
	userAddr   string
	privateKey string
	Info       *Config
}

type NetworkParams struct {
	ChainName       string
	ChainID         int
	RpcUri          string
	WebSocketRpcUri string
	ExplorerUri     string
}

// An AppServiceOption is an option parameter to create the instance of AppService.
type AppServiceOption func(service *AppService)

func NewAppService(opts ...AppServiceOption) (*AppService, error) {
	as := &AppService{}
	for _, opt := range opts {
		opt(as)
	}
	if as.ChainId == "" || as.ChainId == "321" {
		as.Info = GetConfigForKCC()
		if as.ChainId == "" {
			as.ChainId = "321"
		}
		if as.rpcUrl == "" {
			as.rpcUrl = as.Info.Chain.RPCUrl
		}
		ethClient, err := getClient(as.rpcUrl)
		if err != nil {
			return nil, err
		}
		as.rpcClient = ethClient
	} else {
		as.Info = GetConfigForKCCTestnet()
		if as.ChainId == "" {
			as.ChainId = "322"
		}
		if as.rpcUrl == "" {
			as.rpcUrl = as.Info.Chain.RPCUrl
		}
		ethClient, err := getClient(as.rpcUrl)
		if err != nil {
			return nil, err
		}
		as.rpcClient = ethClient
	}

	if GasPriceMultiPoint < 1 {
		GasPriceMultiPoint = 1
	}

	chainIdString, _ := decimal.NewFromString(as.ChainId)
	ChainId = chainIdString.BigInt()

	return as, nil
}

// AppChainIdOption creates an instance of ApiServiceOption about chainId.
func AppChainIdOption(chainId string) AppServiceOption {
	return func(service *AppService) {
		service.ChainId = chainId
	}
}

// AppKeyOption creates an instance of ApiServiceOption about key.
func AppKeyOption(key string) AppServiceOption {
	return func(service *AppService) {
		service.privateKey = key
	}
}

// AppRpcUriOption creates an instance of ApiServiceOption about rpc uri.
func AppRpcUriOption(rpcUri string) AppServiceOption {
	return func(service *AppService) {
		service.rpcUrl = rpcUri
	}
}

// AppGasPriceOption creates an instance of ApiServiceOption about gas price.
func AppGasPriceOption(gasPrice int64) AppServiceOption {
	return func(service *AppService) {
		GasPrice = gasPrice
	}
}

// AppGasPriceMultiOption creates an instance of ApiServiceOption about gas price multipoint.
func AppGasPriceMultiOption(gasPriceMulti float64) AppServiceOption {
	return func(service *AppService) {
		GasPriceMultiPoint = gasPriceMulti
	}
}

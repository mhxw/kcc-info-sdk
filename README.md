# KCC Kit

## Install

```bash
go get github.com/mhxw/kcc-kit
```

## Usage

```go
s, err := kcckit.NewAppService(
    //services.AppChainIdOption(testChainId),
    //services.AppRpcUriOption(testRpc),
    services.AppChainIdOption(mainChainId),
    services.AppRpcUriOption(mainRpc),
    services.AppKeyOption(key),
    services.AppGasPriceOption(gasPrice),
    services.AppGasPriceMultiOption(gasPriceMultipoint),
)
if err != nil {
    return
}
```

## Faucet

- Mainnet

https://faucet.kcc.io

- Testnet

https://faucet-testnet.kcc.network

## License

[MIT](LICENSE)
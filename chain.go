package kcckit

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"math/big"
	"strings"
)

func getClient(rpcUri string) (*ethclient.Client, error) {
	client, err := ethclient.Dial(rpcUri)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (as *AppService) SpiltTx(txHash string) string {
	return as.Info.Chain.TXUrl + txHash
}

func (as *AppService) GetNativeBalance(userAddr string) (*big.Int, error) {
	client := as.rpcClient
	balance, err := client.BalanceAt(context.Background(), common.HexToAddress(userAddr), nil)
	if err != nil {
		return nil, err
	}
	return balance, nil
}

func (as *AppService) TransferNativeBalance(toAddr string, amount *big.Int) (string, error) {
	client := as.rpcClient
	privateKey := as.privateKey

	addrLowerStr := strings.ToLower(privateKey)
	if strings.HasPrefix(addrLowerStr, "0x") {
		addrLowerStr = addrLowerStr[2:]
		privateKey = privateKey[2:]
	}
	userPrivateKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return "", errors.Wrap(err, "private key err")
	}

	publicKey := userPrivateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	userAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), userAddress)
	if err != nil {
		return "", errors.Wrap(err, "PendingNonceAt err")
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "", errors.Wrap(err, "SuggestGasPrice err")
	}
	gasPrice = decimal.NewFromBigInt(gasPrice, 0).Mul(decimal.NewFromFloat(1.2)).BigInt()

	maxBalance, err := client.BalanceAt(context.Background(), userAddress, nil)
	if err != nil {
		return "", errors.Wrap(err, "BalanceAt err")
	}

	toAddress := common.HexToAddress(toAddr)
	gasLimit := int64(21000)
	// value=balance-(gasLimit * gasPrice)
	fee := decimal.NewFromInt(gasLimit).Mul(decimal.NewFromBigInt(gasPrice, 0))
	valueAndFee := decimal.NewFromBigInt(amount, 0).Add(fee).BigInt()
	if maxBalance.Cmp(valueAndFee) < 0 {
		return "", errors.New("insufficient value balance")
	}

	var data []byte
	tx := types.NewTransaction(nonce, toAddress, amount, uint64(gasLimit), gasPrice, data)

	chainID := ChainId

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), userPrivateKey)
	if err != nil {
		return "", errors.Wrap(err, "123")
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return "", errors.Wrap(err, "222")
	}

	return signedTx.Hash().Hex(), nil

}

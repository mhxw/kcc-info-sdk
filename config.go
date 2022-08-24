package kcckit

import "github.com/ethereum/go-ethereum/common"

type Config struct {
	Token  Token
	Mojito Mojito
	Chain  Chain
}

type Token struct {
	BTCAddress   string
	ETHAddress   string
	SKCSAddress  string
	WKCSAddress  string
	WOKTAddress  string
	OKBAddress   string
	USDTAddress  string
	USDCAddress  string
	MJTAddress   string
	CHEAddress   string
	APEAddress   string
	MULTIAddress string
	SAXAddress   string
	SANDAddress  string
	MANAAddress  string
	MLSAddress   string
	AAVEAddress  string
	CRVAddress   string
	UNIAddress   string
	CFXAddress   string
	LINKAddress  string
}

type Mojito struct {
	MojitoFactory    string
	Masterchef       string
	MasterchefV2     string
	MojitoRouter     string
	SwapMiningAddr   string
	MojitoOracle     string
	LP               Lp
	LpPid            LpPid
	Path             Path
	EIP712DomainHash string
	PermitTypeHash   string
}

type Lp struct {
	LpMJTUSDT  string
	LpMJTUSDC  string
	LpUSDTUSDC string
	LpWKCSMJT  string
	LpWKCSUSDC string
	LpWKCSUSDT string
}

type LpPid struct {
	LpMJTUSDTPid  int64
	LpMJTUSDCPid  int64
	LpUSDTUSDCPid int64
}

type Path struct {
	PathUsdcUsdt []common.Address
	PathUsdtUsdc []common.Address
	PathUsdtMjt  []common.Address
	PathMjtUsdt  []common.Address
	PathUsdtChe  []common.Address
	PathCheUsdt  []common.Address
}

type Chain struct {
	Network    string
	ChainId    int64
	TXUrl      string
	RPCUrl     string
	WSUrl      string
	Multicall  string
	Multicall2 string
}

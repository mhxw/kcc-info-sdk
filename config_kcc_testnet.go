package kcckit

import (
	"github.com/ethereum/go-ethereum/common"
)

func GetConfigForKCCTestnet() *Config {
	var (
		BTCAddr   = "0xf57a7eE19a628e4d475b72d6c9DD847c50636e01"
		ETHAddr   = "0xF8Cb9f1D136Ff4c883320b5B4fa80048b888F459"
		WKCSAddr  = "0x6551358EDC7fee9ADAB1E2E49560E68a12E82d9e"
		USDTAddr  = "0x67f6a7bbe0da067a747c6b2bedf8abbf7d6f60dc"
		USDCAddr  = "0xD6c7E27a598714c2226404Eb054e0c074C906Fc9"
		MJTAddr   = "0x208EecDBc49C137D0174B848DEf5F8cB74d6951E"
		APEAddr   = "0xdae6c2a48bfaa66b43815c5548b10800919c993e"
		MULTIAddr = "0x9fb9a33956351cf4fa040f65a13b835a3c8764e3"
		SKCSAddr  = "0x311dD61dF0E88dDc6803e7353F5d9B71522AedA9"
	)

	paramsConfig := Config{
		Token: Token{
			BTCAddress:   BTCAddr,
			ETHAddress:   ETHAddr,
			SKCSAddress:  SKCSAddr,
			WKCSAddress:  WKCSAddr,
			USDTAddress:  USDTAddr,
			USDCAddress:  USDCAddr,
			MJTAddress:   MJTAddr,
			APEAddress:   APEAddr,
			MULTIAddress: MULTIAddr,
		},
		Mojito: Mojito{
			MojitoFactory:    "0x0B9427F175EfB95C3EbB7b9009B12dc685F517E6",
			MojitoRouter:     "0x59a4210Dd69FDdE1457905098fF03E0617A548C5",
			SwapMiningAddr:   "0x4af2211877FbAb15289291ACf26075583AF695EC",
			MojitoOracle:     "0x4254c4ef13Bd3aea3643835D5f2CE5C4f227cF5B",
			Masterchef:       "0x84F10c60Aa2d69aA38Ae307D3bd57e2825BD5617",
			MasterchefV2:     "0xfdfcE767aDD9dCF032Cbd0DE35F0E57b04495324",
			EIP712DomainHash: "0x8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f",
			PermitTypeHash:   "0x6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9",
			LP: Lp{
				LpMJTUSDT:  "0xdd170Beed47e550190cD80C1Bb57F4CD369bD3C1",
				LpMJTUSDC:  "0xA232918Ca4064667F9230Eb30Cd593c7c03959d7",
				LpUSDTUSDC: "0x8726ce53B4850E5154e0dDAF8e80EE97A61f0400",
				LpWKCSMJT:  "0x99Ae42cADE7116cE3d08515977F643534687Eb12",
				LpWKCSUSDT: "0x4047C095D63397Dfc44C9183C2cE01B38ae3c72A",
				LpWKCSUSDC: "0xb349a3429E30C0Cf2C01BCeba9CF05FC5101Ef0f",
			},
			LpPid: LpPid{
				LpMJTUSDTPid:  2,
				LpMJTUSDCPid:  4,
				LpUSDTUSDCPid: 6,
			},
			Path: Path{
				PathUsdcUsdt: []common.Address{
					common.HexToAddress(USDCAddr),
					common.HexToAddress(USDTAddr),
				},
				PathUsdtUsdc: []common.Address{
					common.HexToAddress(USDTAddr),
					common.HexToAddress(USDCAddr),
				},
				PathUsdtMjt: []common.Address{
					common.HexToAddress(USDTAddr),
					common.HexToAddress(MJTAddr),
				},
				PathMjtUsdt: []common.Address{
					common.HexToAddress(MJTAddr),
					common.HexToAddress(USDTAddr),
				},
			},
		},
		Chain: Chain{
			Multicall2: "0x665683D9bd41C09cF38c3956c926D9924F1ADa97",
			Multicall:  "0x54089A493613b57a0799793A8b1CF38F362C2647",
			ChainId:    322,
			TXUrl:      "https://scan-testnet.kcc.network/tx/",
			RPCUrl:     "https://rpc.sdk.wang/kcc-test",
			WSUrl:      "wss://rpc-ws-testnet.kcc.network",
		},
	}
	return &paramsConfig
}

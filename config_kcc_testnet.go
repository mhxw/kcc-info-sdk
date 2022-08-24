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
		Torches: Torches{
			TorchesTroller: "0xf431ABdf916a370dB31Bb97EF74655EfDd80bceF",
			CompoundLens:   "0x67CFeBc7E25b11f65F43C44319Ae8663d7D72ef2",
			TToken: TToken{
				TKCS:  "0xDC9bB1465B9Eb66cD05a4964f2721559894529B1",
				TUSDT: "0xfC49791fF96187ad3260f47CeDdB9c440f25cB6E",
				TUSDC: "0x3d41B01a94cd8538099005565b38de558CAE5EDA",
				TBTC:  "0xDBdAf501d752ef1CB408059D01D05Db3A1811632",
				TETH:  "0x899A2189fB04B3Ed69B568f31D6505340A8910EB",
				TSKCS: "0xE3A92f9113cDD652Ea3AD0A57835b41e8b88dAff",
			},
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
		MojitoOracle: MojitoOracle{
			MojitoOracleAddr: "0xd6c5ac4A130E3B296cCD6443F99d894b8c9dbb23",
			MOProxyAddr:      "0xF0C516F6192FAbBD85398a4EAc9b5f9805011fdF",
			BtcUsdtPairId:    "0x9f3a02ad48a53d7a78ecde8447369718dad4cb83d1dbb6e852af344b6958696f",
			EthUsdtPairId:    "0xc27a9ec66481b5654fa106f0ff4953f464c930ddeadc29dff13c50670c0d8cf7",
			KcsUsdtPairId:    "0x10e4b55f0e4ec1ba07c58b6aac7d37f5a218efc2c3f4f4b637934b65573eb19d",
			MjtUsdtPairId:    "0xcef32d3651d74bf39b0f08c22a7ec83cc66c029645f1a915ee050fb74242e21a",
			UsdcUsdtPairId:   "0x0cf25c1cf37d7aba87b730218eba67b54738879302f4ebbc32f8b11505db2084",
		},
		TorchesOralce: TorchesOralce{
			BtcUsd: EACAndOCR{
				OCR: "0xC51898c389d3591B9A15b9302A79B143058b4255",
				EAC: "0xBb3423a913a9a69aD7Dba09B62abdFDE4643BAe4",
			},
			EthUsd: EACAndOCR{
				OCR: "0xC86C325bd7889b222e86a19C719FbF59E3035DD4",
				EAC: "0x22337a9a305E081c0C801dd7B7b8eCF4966660bB",
			},
			KcsUsd: EACAndOCR{
				OCR: "0x5E71dbD99E1b1F4b2d17Ae5908D17481980aA712",
				EAC: "0xae3DB39196012a7bF6D38737192F260cdFE1E7Ec",
			},
			MjtUsd: EACAndOCR{
				OCR: "0xEd27394f20FC294dAA334Cd8eef2260f0241D7F0",
				EAC: "0x11Eb72402ABA2031dAc555F158e23614009b1b6f",
			},
			UsdtUsd: EACAndOCR{
				OCR: "0xC51898c389d3591B9A15b9302A79B143058b4255",
				EAC: "0xBb3423a913a9a69aD7Dba09B62abdFDE4643BAe4",
			},
			UsdcUsd: EACAndOCR{
				OCR: "0xC86C325bd7889b222e86a19C719FbF59E3035DD4",
				EAC: "0x22337a9a305E081c0C801dd7B7b8eCF4966660bB",
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

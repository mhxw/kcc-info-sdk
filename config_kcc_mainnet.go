package kcckit

import (
	"github.com/ethereum/go-ethereum/common"
)

func GetConfigForKCC() *Config {
	var (
		BTCAddr   = "0xfa93c12cd345c658bc4644d1d4e1b9615952258c"
		ETHAddr   = "0xf55af137a98607f7ed2efefa4cd2dfe70e4253b1"
		WKCSAddr  = "0x4446Fc4eb47f2f6586f9fAAb68B3498F86C07521"
		USDTAddr  = "0x0039f574eE5cC39bdD162E9A88e3EB1f111bAF48"
		USDCAddr  = "0x980a5afef3d17ad98635f6c5aebcbaeded3c3430"
		MJTAddr   = "0x2cA48b4eeA5A731c2B54e7C3944DBDB87c0CFB6F"
		APEAddr   = "0xdae6c2a48bfaa66b43815c5548b10800919c993e"
		MULTIAddr = "0x9fb9a33956351cf4fa040f65a13b835a3c8764e3"
		SAXAddr   = "0x8367B781316eDb622DB3847f863e610edf1f4a20"
		SANDAddr  = "0xb12c13e66ade1f72f71834f2fc5082db8c091358"
		MANAAddr  = "0xc19a5cacc2bb68ff09f2fcc695f31493a039fa5e"
		MLSAddr   = "0x974E54266708292383C2989DCA5FDbb115666D4f"
		CRVAddr   = "0x4500e16da66b99e0c55d7b46ebbd59bc413ba171"
		AAVEAddr  = "0xe76e97c157658004ee22e01c03a5e21a4655a2fd"
		UNIAddr   = "0xee58e4d62b10a92db1089d4d040b759c28ae16cd"
		CFXAddr   = "0x4792c1ecb969b036eb51330c63bd27899a13d84e"
		LinkAddr  = "0x47841910329aaa6b88d5e9dcde9000195151dc72"
		SKCSAddr  = "0x00eE2d494258D6C5A30d6B6472A09b27121Ef451"
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
			SAXAddress:   SAXAddr,
			SANDAddress:  SANDAddr,
			MANAAddress:  MANAAddr,
			MLSAddress:   MLSAddr,
			AAVEAddress:  AAVEAddr,
			UNIAddress:   UNIAddr,
			CRVAddress:   CRVAddr,
			MULTIAddress: MULTIAddr,
			CFXAddress:   CFXAddr,
			LINKAddress:  LinkAddr,
		},
		Torches: Torches{
			TorchesTroller: "0xfbAFd34A4644DC4f7c5b2Ae150279162Eb2B0dF6",
			CompoundLens:   "0x5B0Ad1D8d438B00b0730134c8F0BaEf3ed445a0A",
			TToken: TToken{
				TKCS:  "0xf9401F5246185eD3Fd0EF48f4775250d32069AEf",
				TUSDT: "0x92dBEA1Ac6278a0b4AEC11388C94F8fAFBE246C1",
				TUSDC: "0x23eBfcC4aB0Ddf47E7b99B1eb1B9eA790f3d646D",
				TBTC:  "0x413170815cF21B16ddeB180aAd0FFD00C80D31f8",
				TETH:  "0xF7310bD46D48e8d2BF8d12f69B6cCdD12fB4E0C5",
				TSKCS: "0x0868713842d2e296CeF26c86d736AC7C374A5199",
			},
		},
		Mojito: Mojito{
			MojitoFactory:    "0x79855a03426e15ad120df77efa623af87bd54ef3",
			MojitoRouter:     "0x8c8067ed3bc19acce28c1953bfc18dc85a2127f7",
			SwapMiningAddr:   "0x79aa527f8d54016a75f5799da4ee198bb522413d",
			MojitoOracle:     "0xC8a01e731eCC1443028F6BbA2706c077D544a6Fd",
			Masterchef:       "0x25C6d6A65C3ae5d41599Ba2211629B24604Fea4F",
			MasterchefV2:     "0xfdfcE767aDD9dCF032Cbd0DE35F0E57b04495324",
			EIP712DomainHash: "0x8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f",
			PermitTypeHash:   "0x6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c9",
			LP: Lp{
				LpMJTUSDT:  "0xdd170Beed47e550190cD80C1Bb57F4CD369bD3C1",
				LpMJTUSDC:  "0xA232918Ca4064667F9230Eb30Cd593c7c03959d7",
				LpUSDTUSDC: "0xeb06057E2405c8819e2cfFEA5Dea07A54ad569e3",
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
			MojitoOracleAddr: "0x6ED03dE8D8BA4c90E7C8A2317bED899AF1959491",
			MOProxyAddr:      "0x09515646b5833526Dff12FAE9B03754c7f0E0a8F",
			BtcUsdtPairId:    "0x9f3a02ad48a53d7a78ecde8447369718dad4cb83d1dbb6e852af344b6958696f",
			EthUsdtPairId:    "0xc27a9ec66481b5654fa106f0ff4953f464c930ddeadc29dff13c50670c0d8cf7",
			KcsUsdtPairId:    "0x10e4b55f0e4ec1ba07c58b6aac7d37f5a218efc2c3f4f4b637934b65573eb19d",
			MjtUsdtPairId:    "0xcef32d3651d74bf39b0f08c22a7ec83cc66c029645f1a915ee050fb74242e21a",
			UsdcUsdtPairId:   "0x0cf25c1cf37d7aba87b730218eba67b54738879302f4ebbc32f8b11505db2084",
		},
		Witnet: Witnet{
			WitnetPriceRouter: "0xD39D4d972C7E166856c4eb29E54D3548B4597F53",
			BtcUsdPairId:      "0x24beead43216e490aa240ef0d32e18c57beea168f06eabb94f5193868d500946",
			EthUsdPairId:      "0x3d15f7018db5cc80838b684361aaa100bfadf8a11e02d5c1c92e9c6af47626c8",
			KcsUsdtPairId:     "0x31debffc453c5d04a78431e7bc28098c606d2bbeea22f10a35809924a201a977",
			MjtKcsPairId:      "0x2dcfd5546926b857978957b40dcd5164cc788079b46ce9c1abbaedac07f96837",
			UsdtUsdPairId:     "0x538f5a25b39995a23c24037d2d38f979c8fa7b00d001e897212d936e6f6556ef",
			UsdcUsdPairId:     "0x4c80cf2e5b3d17b98f6f24fc78f661982b8ef656c3b75a038f7bfc6f93c1b20e",
		},
		TorchesOralce: TorchesOralce{
			BtcUsd: EACAndOCR{
				EAC: "0xFAce3f85602A8dc013217b61a97a9AFE7B2F276F",
				OCR: "0x13cc7A061Ee93157f2bcE90645cD9f754428e08A",
			},
			EthUsd: EACAndOCR{
				EAC: "0x72E10386eBE0E3175f62BF3Edfc9A64aC3c5918a",
				OCR: "0x0948C486F5CEDfb22CBc55DD2a3C9841043550Dd",
			},
			KcsUsd: EACAndOCR{
				EAC: "0xAFC9c849b1a784955908d91EE43A3203fBC1f950",
				OCR: "0xF12d64D3d51AC263cd5e41fd6b634E9A55F2A0d9",
			},
			MjtUsd: EACAndOCR{
				EAC: "0x5eF7D0B6C63c9F0b0b056416B2dBA95cC02473a3",
				OCR: "0xa10DD65F2231Aea80f7d3BC53789d520620A4291",
			},
			UsdtUsd: EACAndOCR{
				EAC: "0x001c1a168ba2a36D01a99542740C375c51615161",
				OCR: "0xbFBE130ce1302118e90148eeaeb5Ce0ad73a0c11",
			},
			UsdcUsd: EACAndOCR{
				EAC: "0x1A165db46d431804B0082eb5BEbc307ffb97e31b",
				OCR: "0x0f10d056b488ED7ad72EabE677A2fa5Bd665ef67",
			},
		},
		Chain: Chain{
			Multicall2: "0x7c1c85c39d3d6b6ecb811dfe949b9c23f6e818b0",
			Multicall:  "0x83a7e42a1190346073323dfaf40d69f8f8c076a4",
			ChainId:    321,
			TXUrl:      "https://explorer.kcc.io/zh/tx/",
			RPCUrl:     "https://rpc-mainnet.kcc.network",
			WSUrl:      "wss://rpc-ws-mainnet.kcc.network",
		},
	}
	return &paramsConfig
}

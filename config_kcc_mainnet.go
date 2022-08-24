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

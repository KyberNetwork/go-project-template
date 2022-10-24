package main

import (
	"bytes"
	"context"
	"log"
	"math/big"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/KyberNetwork/go-project-template/pkg/contracts/erc20"
	"github.com/KyberNetwork/go-project-template/pkg/contracts/iuniswapv2router"
	"github.com/KyberNetwork/go-project-template/pkg/contracts/simutil"
	"github.com/KyberNetwork/go-project-template/pkg/contracts/uniswapv3quoter"
	"github.com/KyberNetwork/go-project-template/pkg/contracts/uniswapv3router"
	"github.com/KyberNetwork/go-project-template/pkg/convert"
	"github.com/KyberNetwork/go-project-template/pkg/onchain/kyberswap"
	"github.com/KyberNetwork/go-project-template/pkg/onchain/simclient"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/avmclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	rpcURL   = "http://192.168.11.4:8545"
	myWallet = common.HexToAddress("0x0000000000000000000000000000000000111101")

	kyberswapRouterAddress = common.HexToAddress("0x617Dee16B86534a5d792A4d7A62FB491B544111E") // 0x617Dee16B86534a5d792A4d7A62FB491B544111E
	univ2routerAddress     = common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")
	uniswapV3Router        = common.HexToAddress("0xE592427A0AEce92De3Edee1F18E0157C05861564")
	uniswapV3Quoter        = common.HexToAddress("0xb27308f9F90D607463bb33eA1BeBb41C27CE5AB6")

	kncAddress  = common.HexToAddress("0xdeFA4e8a7bcBA345F687a2f1456F5Edd9CE97202")
	usdtAddress = common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7")
	usdcAddress = common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")
	wethAddress = common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")

	overrideAccounts = simclient.OverrideAccounts{
		myWallet: {
			Balance: "0x8ac7230489e80000",
		},
		kncAddress: {
			StateDiff: map[string]string{
				"0xee62773c1388dc8fdd9d3b67edd9f007212b1bd8331e5996b697ac5315d3cd90": "0x00000000000000000000000000000000000000000000993635c9adc5dea00000", // balance
				"0xcae1dfe8c33a033da2b4c601c9036189105639b4b09809f8868d75ccb3c69478": "0x00000000000000000000000000000000000000000000993635c9adc5dea00000", // allownace
			},
		},
	}

	univ3Overrides = simclient.OverrideAccounts{
		myWallet: {
			Balance: "0x8ac7230489e80000",
		},
		usdcAddress: {
			StateDiff: map[string]string{
				"0x12ad0a2e9dbaa72fddc19f4ed13c3647f1ae39b16be9478c6a1eca1bf074de65": "0x00000000000000000000000000000000000000000000993635c9adc5dea00000",
				"0x4e6c44dbae1976917ab2cb32265dcd82cd04024b0596a22ef872a96cffbeb913": "0x00000000000000000000000000000000000000000000993635c9adc5dea00000",
			},
		},
	}
)

func requireNoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func TestKyberswapAgg() {
	httpClient := &http.Client{Transport: &http.Transport{
		MaxConnsPerHost: 32,
	}}
	kclient := kyberswap.NewClient(httpClient)
	tokenIn := kncAddress
	tokenOut := wethAddress
	amountIn := convert.FloatToWei(100, 18)

	swapInfo, err := kclient.GetSwapInfo(context.Background(), "ethereum", kyberswap.GetSwapInfoInputParams{
		AmountIn:          amountIn,
		To:                &myWallet,
		TokenIn:           &tokenIn,
		TokenOut:          &tokenOut,
		SaveGas:           false,
		ChargeFeeBy:       "currency_in",
		SlippageTolerance: "50",
		ClientData:        "",
		Deadline:          "",
		FeeAmount:         big.NewInt(0),
		FeeReceiver:       &myWallet,
		IsInBps:           true,
	})
	requireNoErr(err)
	if len(swapInfo.EncodedSwapData) == 0 {
		log.Fatalf("api return empty encoded data")
	}

	swapData := hexutil.MustDecode(swapInfo.EncodedSwapData)
	// log.Println("swap data", swapInfo.EncodedSwapData)
	router := swapInfo.RouterAddress
	if *router != kyberswapRouterAddress {
		log.Fatalf("router address changed to %s", swapInfo.RouterAddress.String())
	}

	simClient, err := simclient.NewSimClient(rpcURL, httpClient, overrideAccounts)
	requireNoErr(err)

	block, err := simClient.BlockNumber(context.Background())
	requireNoErr(err)
	log.Println("current block", block)

	inERC20Token, err := erc20.NewErc20(tokenIn, simClient)
	requireNoErr(err)
	balance, err := inERC20Token.BalanceOf(&bind.CallOpts{}, myWallet)
	requireNoErr(err)
	log.Println("KNC balance", balance.String())
	allow, err := inERC20Token.Allowance(&bind.CallOpts{}, myWallet, kyberswapRouterAddress)
	requireNoErr(err)
	log.Println("allowance on router: ", allow.String())

	// deployedContract := common.HexToAddress("0x81926273bb7393340aba83f4acda4af47c0862d7")
	msg := ethereum.CallMsg{
		From:      myWallet,
		To:        &kyberswapRouterAddress,
		Gas:       5000_000,
		GasPrice:  convert.FloatToWei(100, 9),
		GasFeeCap: convert.FloatToWei(100, 9),
		GasTipCap: convert.FloatToWei(100, 9),
		Value:     big.NewInt(0), // convert.FloatToTokenAmount(0.1, 18),
		Data:      swapData,
	}

	var counter = uint64(0)
	const numRoutine = 1

	for i := 0; i < numRoutine; i++ {
		go func() {
			for {
				_, err := simClient.CallContract(context.Background(), msg, nil)
				requireNoErr(err)
				atomic.AddUint64(&counter, 1)
				break
			}
		}()
	}
	for range time.NewTicker(time.Second).C {
		count := atomic.SwapUint64(&counter, 0)
		log.Println("request rate", count)
	}
}

func TestUniswapv2() {

	c := &http.Client{Transport: &http.Transport{
		MaxConnsPerHost:     128,
		MaxIdleConns:        64,
		MaxIdleConnsPerHost: 64,
	}}
	amountIn := convert.FloatToWei(5000, 6)
	r, err := rpc.DialHTTPWithClient(rpcURL, c)
	requireNoErr(err)
	simClient := ethclient.NewClient(r) // simclient.NewSimClient(rpcURL, c, univ2Overrides)

	uniabi, _ := abi.JSON(bytes.NewBuffer([]byte(iuniswapv2router.Iuniswapv2routerMetaData.ABI)))
	data, err := uniabi.Pack("getAmountsOut", amountIn, []common.Address{usdcAddress, wethAddress})
	requireNoErr(err)

	var counter = uint64(0)
	const numRoutine = 12

	for i := 0; i < numRoutine; i++ {
		go func() {
			for {
				_, err := simClient.CallContract(context.Background(), ethereum.CallMsg{
					From: myWallet,
					To:   &univ2routerAddress,
					Data: data,
					Gas:  8000000,
				}, nil)
				requireNoErr(err)
				atomic.AddUint64(&counter, 1)
			}
		}()
	}
	for range time.NewTicker(time.Second).C {
		count := atomic.SwapUint64(&counter, 0)
		log.Println("request rate", count)
	}
}

func TestFakeAccount() {

	c := &http.Client{Transport: &http.Transport{
		MaxConnsPerHost:     128,
		MaxIdleConns:        64,
		MaxIdleConnsPerHost: 64,
	}}
	fakeSimUtilAddress := common.HexToAddress("0x11111111111111111111111111111111111111aa")
	fakeUserAddress := common.HexToAddress("0x1111111111111111111111111111111111111100")

	avmClient := avmclient.New(c, "http://192.168.11.4:8345")

	uniabi, _ := abi.JSON(bytes.NewBuffer([]byte(simutil.SimUtilABI)))
	data, err := uniabi.Pack("getBalances",
		[]common.Address{fakeUserAddress},
		[]common.Address{usdcAddress, wethAddress})
	requireNoErr(err)

	rdata, err := avmClient.CustomCall(avmclient.CallMsg{
		From: fakeUserAddress,
		To:   fakeSimUtilAddress,
		Data: data,
		Overrides: []avmclient.Override{
			{
				Address: fakeSimUtilAddress,
				Code:    hexutil.MustDecode("0x608060405234801561001057600080fd5b506004361061002b5760003560e01c8063ef5bfc3714610030575b600080fd5b6100fc6004803603604081101561004657600080fd5b810190808035906020019064010000000081111561006357600080fd5b82018360208201111561007557600080fd5b8035906020019184602083028401116401000000008311171561009757600080fd5b9091929391929390803590602001906401000000008111156100b857600080fd5b8201836020820111156100ca57600080fd5b803590602001918460208302840111640100000000831117156100ec57600080fd5b9091929391929390505050610153565b6040518080602001828103825283818151815260200191508051906020019060200280838360005b8381101561013f578082015181840152602081019050610124565b505050509050019250505060405180910390f35b60608083839050868690500267ffffffffffffffff8111801561017557600080fd5b506040519080825280602002602001820160405280156101a45781602001602082028036833780820191505090505b50905060005b868690508110156103c45760005b858590508110156103b65773eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee73ffffffffffffffffffffffffffffffffffffffff168686838181106101fa57fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156102975787878381811061023f57fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16318382888890508502018151811061028657fe5b6020026020010181815250506103a9565b8585828181106102a357fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166370a082318989858181106102e757fe5b9050602002013573ffffffffffffffffffffffffffffffffffffffff166040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561034e57600080fd5b505afa158015610362573d6000803e3d6000fd5b505050506040513d602081101561037857600080fd5b81019080805190602001909291905050508382888890508502018151811061039c57fe5b6020026020010181815250505b80806001019150506101b8565b5080806001019150506101aa565b508091505094935050505056fea26469706673582212200b8738e9ce84a5af761886e464d14ce62b0d8bce0316bc15a43ba74267a846ea64736f6c634300060c0033"),
			},
		},
	})

	requireNoErr(err)
	var out []*big.Int
	err = uniabi.UnpackIntoInterface(&out, "getBalances", rdata)
	requireNoErr(err)
	log.Println(out)
}

func TestUniswapv2CustomRate() {

	c := &http.Client{Transport: &http.Transport{
		MaxConnsPerHost:     128,
		MaxIdleConns:        64,
		MaxIdleConnsPerHost: 64,
	}}
	amountIn := convert.FloatToWei(5000, 6)
	avmClient := avmclient.New(c, "http://192.168.11.4:8345")
	uniabi, _ := abi.JSON(bytes.NewBuffer([]byte(iuniswapv2router.Iuniswapv2routerMetaData.ABI)))
	// data, err := uniabi.Pack("swapExactETHForTokens", big.NewInt(0),
	//	[]common.Address{wethAddress, usdcAddress}, myWallet, big.NewInt(time.Now().Unix()+600))
	data, err := uniabi.Pack("getAmountsOut", amountIn, []common.Address{usdcAddress, wethAddress})
	requireNoErr(err)

	var counter = uint64(0)
	const numRoutine = 12

	for i := 0; i < numRoutine; i++ {
		go func() {
			for {
				_, err := avmClient.CustomCall(avmclient.CallMsg{
					From: myWallet,
					To:   univ2routerAddress,
					Data: data,
				})
				requireNoErr(err)
				atomic.AddUint64(&counter, 1)
			}
		}()
	}
	for range time.NewTicker(time.Second).C {
		count := atomic.SwapUint64(&counter, 0)
		log.Println("request rate", count)
	}
}

func TestUniswapv3() {

	c := &http.Client{Transport: &http.Transport{
		MaxConnsPerHost:     128,
		MaxIdleConns:        64,
		MaxIdleConnsPerHost: 64,
	}}
	amountIn := convert.FloatToWei(5000, 6)

	simClient, err := simclient.NewSimClient(rpcURL, c, univ3Overrides)
	requireNoErr(err)

	var counter = uint64(0)
	const numRoutine = 1
	uniabi, _ := abi.JSON(bytes.NewBuffer([]byte(uniswapv3router.Uniswapv3MetaData.ABI)))
	data, err := uniabi.Pack("exactInputSingle", uniswapv3router.ISwapRouterExactInputSingleParams{
		TokenIn:           usdcAddress,
		TokenOut:          wethAddress,
		Fee:               big.NewInt(3000),
		Recipient:         myWallet,
		Deadline:          big.NewInt(time.Now().Unix() + 600),
		AmountIn:          amountIn,
		AmountOutMinimum:  big.NewInt(0),
		SqrtPriceLimitX96: big.NewInt(0),
	})
	requireNoErr(err)
	for i := 0; i < numRoutine; i++ {
		go func() {
			for {
				buff, err := simClient.CallContract(context.Background(), ethereum.CallMsg{
					From: myWallet,
					To:   &uniswapV3Router,
					Gas:  8000000,
					Data: data,
				}, nil)
				requireNoErr(err)
				atomic.AddUint64(&counter, 1)

				var res *big.Int
				uniabi.UnpackIntoInterface(&res, "exactInputSingle", buff)
				log.Println(res)
				break
			}
		}()
	}
	for range time.NewTicker(time.Second).C {
		count := atomic.SwapUint64(&counter, 0)
		log.Println("request rate", count)
	}
}

func TestUniswapv3Custom() {

	c := &http.Client{Transport: &http.Transport{
		MaxConnsPerHost:     128,
		MaxIdleConns:        64,
		MaxIdleConnsPerHost: 64,
	}}

	amountIn := convert.FloatToWei(5000, 6)

	avmClient := avmclient.New(c, "http://192.168.11.4:8345")

	uniabi, _ := abi.JSON(bytes.NewBuffer([]byte(uniswapv3router.Uniswapv3MetaData.ABI)))
	data, err := uniabi.Pack("exactInputSingle", uniswapv3router.ISwapRouterExactInputSingleParams{
		TokenIn:           usdcAddress,
		TokenOut:          wethAddress,
		Fee:               big.NewInt(3000),
		Recipient:         myWallet,
		Deadline:          big.NewInt(time.Now().Unix() + 600),
		AmountIn:          amountIn,
		AmountOutMinimum:  big.NewInt(0),
		SqrtPriceLimitX96: big.NewInt(0),
	})
	requireNoErr(err)

	var counter = uint64(0)
	const numRoutine = 12
	overrides := []avmclient.Override{
		{
			Address: myWallet,
			Balance: convert.FloatToWei(1, 18),
		},
		{
			Address: usdcAddress,
			StateDiff: map[common.Hash]common.Hash{
				common.HexToHash("0x12ad0a2e9dbaa72fddc19f4ed13c3647f1ae39b16be9478c6a1eca1bf074de65"): common.HexToHash("0x00000000000000000000000000000000000000000000993635c9adc5dea00000"),
				common.HexToHash("0x4e6c44dbae1976917ab2cb32265dcd82cd04024b0596a22ef872a96cffbeb913"): common.HexToHash("0x00000000000000000000000000000000000000000000993635c9adc5dea00000"),
			},
		},
	}
	for i := 0; i < numRoutine; i++ {
		go func() {
			for {
				_, err := avmClient.CustomCall(avmclient.CallMsg{
					From:      myWallet,
					To:        uniswapV3Router,
					Gas:       avmclient.Uint64Ptr(8000000),
					Data:      data,
					Overrides: overrides,
				})
				requireNoErr(err)
				atomic.AddUint64(&counter, 1)
			}
		}()
	}
	for range time.NewTicker(time.Second).C {
		count := atomic.SwapUint64(&counter, 0)
		log.Println("request rate", count)
	}
}

func TestUniswapv3QuoterCustom() {

	c := &http.Client{Transport: &http.Transport{
		MaxConnsPerHost:     128,
		MaxIdleConns:        64,
		MaxIdleConnsPerHost: 64,
	}}

	amountIn := convert.FloatToWei(500000, 6)
	avmClient := avmclient.New(c, "http://192.168.11.4:8345")

	uniabi, _ := abi.JSON(bytes.NewBuffer([]byte(uniswapv3quoter.Uniswapv3quoterMetaData.ABI)))
	data, err := uniabi.Pack("quoteExactInputSingle", usdcAddress, wethAddress, big.NewInt(3000), amountIn, big.NewInt(0))
	requireNoErr(err)

	var counter = uint64(0)
	const numRoutine = 1
	for i := 0; i < numRoutine; i++ {
		go func() {
			for {
				_, err := avmClient.CustomCall(avmclient.CallMsg{
					From: myWallet,
					To:   uniswapV3Quoter,
					Gas:  avmclient.Uint64Ptr(8000000),
					Data: data,
					// Overrides: overrides,
				})
				requireNoErr(err)
				atomic.AddUint64(&counter, 1)
			}
		}()
	}
	for range time.NewTicker(time.Second).C {
		count := atomic.SwapUint64(&counter, 0)
		log.Println("request rate", count)
	}
}

func main() {
	// TestSimcall() // 2022/10/12 07:20:25 current block 15730500
	// TestUniswapv2()
	// TestUniswapv2Custom()
	// TestFakeAccount()
	// TestSimUtilcall()
	// TestUniswapv3Custom()
	// TestUniswapv3()
	TestUniswapv3QuoterCustom()
	// TestUniswapv2CustomRate()
}

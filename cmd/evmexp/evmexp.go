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

	kncAddress         = common.HexToAddress("0xdeFA4e8a7bcBA345F687a2f1456F5Edd9CE97202")
	usdtAddress        = common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7")
	usdcAddress        = common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")
	wethAddress        = common.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")
	univ2routerAddress = common.HexToAddress("0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D")

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
)

func requireNoErr(err error) {
	if err != nil {
		panic(err)
	}
}

func TestSimcall() {
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
	rc, err := rpc.DialHTTPWithClient(rpcURL, c)
	requireNoErr(err)
	ethClient := ethclient.NewClient(rc)
	univ2Client, err := iuniswapv2router.NewIuniswapv2router(univ2routerAddress, ethClient)
	requireNoErr(err)

	block, err := ethClient.BlockNumber(context.Background())
	requireNoErr(err)
	log.Println("current block", block)

	var counter = uint64(0)
	const numRoutine = 12

	for i := 0; i < numRoutine; i++ {
		go func() {
			for {
				_, err := univ2Client.GetAmountsOut(&bind.CallOpts{}, amountIn, []common.Address{usdcAddress, wethAddress})
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

func TestUniswapv2Custom() {

	c := &http.Client{Transport: &http.Transport{
		MaxConnsPerHost:     128,
		MaxIdleConns:        64,
		MaxIdleConnsPerHost: 64,
	}}
	amountIn := convert.FloatToWei(5000, 6)

	avmClient := avmclient.New(c, "http://192.168.11.4:8345")

	uniabi, _ := abi.JSON(bytes.NewBuffer([]byte(iuniswapv2router.Iuniswapv2routerMetaData.ABI)))
	data, err := uniabi.Pack("getAmountsOut", amountIn, []common.Address{usdcAddress, wethAddress})
	requireNoErr(err)

	var counter = uint64(0)
	const numRoutine = 12

	for i := 0; i < numRoutine; i++ {
		go func() {
			for {
				_, err := avmClient.CustomCall(avmclient.CallMsg{
					To:   univ2routerAddress,
					Data: data,
					Gas:  8000000,
				})
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

func main() {
	// TestSimcall() // 2022/10/12 07:20:25 current block 15730500
	// TestUniswapv2()
	TestUniswapv2Custom()
}

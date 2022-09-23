package main

import (
	"context"
	"log"
	"math/big"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/KyberNetwork/go-project-template/pkg/contracts/erc20"
	"github.com/KyberNetwork/go-project-template/pkg/convert"
	"github.com/KyberNetwork/go-project-template/pkg/onchain/kyberswap"
	"github.com/KyberNetwork/go-project-template/pkg/onchain/simclient"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	rpcURL   = "http://localhost:8545"
	myWallet = common.HexToAddress("0x0000000000000000000000000000000000111101")

	ksRouterAddress = common.HexToAddress("0x617Dee16B86534a5d792A4d7A62FB491B544111E") // 0x617Dee16B86534a5d792A4d7A62FB491B544111E

	kncAddress     = common.HexToAddress("0xdeFA4e8a7bcBA345F687a2f1456F5Edd9CE97202")
	usdtAddress    = common.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7")
	CommonContract = simclient.OverrideAccounts{
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
	tokenOut := usdtAddress
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
		IsInBps:           false,
	})
	requireNoErr(err)
	if len(swapInfo.EncodedSwapData) == 0 {
		log.Fatalf("api return empty encoded data")
	}

	swapData := hexutil.MustDecode(swapInfo.EncodedSwapData)
	log.Println("swap data", swapInfo.EncodedSwapData)
	router := swapInfo.RouterAddress
	if *router != ksRouterAddress {
		log.Fatalf("router address changed to %s", swapInfo.RouterAddress.String())
	}

	client, err := ethclient.Dial(rpcURL)
	requireNoErr(err)
	block, err := client.BlockNumber(context.Background())
	requireNoErr(err)
	log.Println("block", block)
	simClient, err := simclient.NewSimClient(rpcURL, httpClient, CommonContract)
	requireNoErr(err)

	inERC20Token, err := erc20.NewErc20(tokenIn, simClient)
	requireNoErr(err)
	balance, err := inERC20Token.BalanceOf(&bind.CallOpts{}, myWallet)
	requireNoErr(err)
	log.Println("KNC balance", balance.String())
	allow, err := inERC20Token.Allowance(&bind.CallOpts{}, myWallet, ksRouterAddress)
	requireNoErr(err)
	log.Println("allowance on router: ", allow.String())

	// deployedContract := common.HexToAddress("0x81926273bb7393340aba83f4acda4af47c0862d7")
	msg := ethereum.CallMsg{
		From:      myWallet,
		To:        &ksRouterAddress,
		Gas:       5000_000,
		GasPrice:  convert.FloatToWei(100, 9),
		GasFeeCap: convert.FloatToWei(100, 9),
		GasTipCap: convert.FloatToWei(100, 9),
		Value:     big.NewInt(0), // convert.FloatToTokenAmount(0.1, 18),
		Data:      swapData,
	}

	var counter = uint64(0)

	for i := 0; i < 8; i++ {
		go func() {
			for {
				_, err := simClient.CallContract(context.Background(), msg, nil)
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
	TestSimcall()
}

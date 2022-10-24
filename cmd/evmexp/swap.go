package main

import (
	"bytes"
	"log"
	"math/big"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/KyberNetwork/go-project-template/pkg/contracts/iuniswapv2router"
	"github.com/KyberNetwork/go-project-template/pkg/convert"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient/avmclient"
)

func TestUniswapv2Custom() {

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
	data, err := uniabi.Pack("swapExactTokensForETH", amountIn, big.NewInt(0),
		[]common.Address{usdcAddress, wethAddress}, myWallet, big.NewInt(time.Now().Unix()+600))
	requireNoErr(err)

	var counter = uint64(0)
	const numRoutine = 12
	overrides := []avmclient.Override{
		{
			Address: myWallet,
			Balance: convert.FloatToWei(100, 18),
		},
		{
			Address: usdcAddress,
			StateDiff: map[common.Hash]common.Hash{
				common.HexToHash("0x12ad0a2e9dbaa72fddc19f4ed13c3647f1ae39b16be9478c6a1eca1bf074de65"): common.HexToHash("0x0000000000000000000000000000000000000000ffff993635c9adc5dea00000"),
				common.HexToHash("0xa75d6ac5bed3f4f8130d1ca3a985bccd4c92c846677a8a7b913893a4be23b6f8"): common.HexToHash("0x0000000000000000000000000000000000000000ffff993635c9adc5dea00000"),
			},
		},
	}

	for i := 0; i < numRoutine; i++ {
		go func() {
			for {
				_, err := avmClient.CustomCall(avmclient.CallMsg{
					From:      myWallet,
					To:        univ2routerAddress,
					Data:      data,
					Value:     nil,
					Gas:       avmclient.Uint64Ptr(8000000),
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

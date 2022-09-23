package kyberswap

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

// DTO: the struct that has closest data type the internal object using

type GetSwapInfoInputParams struct {
	AmountIn *big.Int
	To       *common.Address
	TokenIn  *common.Address
	TokenOut *common.Address
	SaveGas  bool

	ChargeFeeBy       string
	SlippageTolerance string
	ClientData        string
	Deadline          string
	FeeAmount         *big.Int
	FeeReceiver       *common.Address
	IsInBps           bool
	Dexes             string
}

type GetSwapInfoOutput struct {
	InputAmount  *big.Int
	OutputAmount *big.Int
	TotalGas     *big.Int

	GasPriceGwei    string
	GasUsd          float64
	AmountInUsd     float64
	AmountOutUsd    float64
	ReceivedUsd     float64
	Swaps           [][]SwapResponse
	EncodedSwapData string
	RouterAddress   *common.Address
}

func ToGetSwapInfoRequestParams(input GetSwapInfoInputParams) *GetSwapInfoRequestParams {
	var (
		amountIn, to, tokenIn, tokenOut, feeAmount, feeReceiver string
		saveGas                                                 = "0"
	)

	if input.AmountIn != nil {
		amountIn = input.AmountIn.String()
	}
	if input.To != nil {
		to = input.To.String()
	}
	if input.TokenIn != nil {
		tokenIn = input.TokenIn.String()
	}
	if input.TokenOut != nil {
		tokenOut = input.TokenOut.String()
	}
	if input.FeeAmount != nil {
		feeAmount = input.FeeAmount.String()
	}
	if input.FeeReceiver != nil {
		feeReceiver = input.FeeReceiver.String()
	}
	if input.SaveGas {
		saveGas = "1"
	}

	return &GetSwapInfoRequestParams{
		AmountIn: amountIn,
		To:       to,
		TokenIn:  tokenIn,
		TokenOut: tokenOut,
		SaveGas:  saveGas,

		ChargeFeeBy:       input.ChargeFeeBy,
		SlippageTolerance: input.SlippageTolerance,
		ClientData:        input.ClientData,
		Deadline:          input.Deadline,
		FeeAmount:         feeAmount,
		FeeReceiver:       feeReceiver,
		IsInBps:           input.IsInBps,
		Dexes:             input.Dexes,
	}
}

func ToGetSwapInfoOutput(input GetSwapInfoResponse) (*GetSwapInfoOutput, error) {
	inputAmount, ok := new(big.Int).SetString(input.InputAmount, 10)
	if !ok {
		return nil, fmt.Errorf("convert InputAmount failed: %v", input.InputAmount)
	}
	outputAmount, ok := new(big.Int).SetString(input.OutputAmount, 10)
	if !ok {
		return nil, fmt.Errorf("convert InputAmount failed: %v", input.OutputAmount)
	}
	totalGas := big.NewInt(int64(input.TotalGas))
	routerAddress := common.HexToAddress(input.RouterAddress)

	return &GetSwapInfoOutput{
		InputAmount:  inputAmount,
		OutputAmount: outputAmount,
		TotalGas:     totalGas,

		GasPriceGwei:    input.GasPriceGwei,
		GasUsd:          input.GasUsd,
		AmountInUsd:     input.AmountInUsd,
		AmountOutUsd:    input.AmountOutUsd,
		ReceivedUsd:     input.ReceivedUsd,
		Swaps:           input.Swaps,
		EncodedSwapData: input.EncodedSwapData,
		RouterAddress:   &routerAddress,
	}, nil
}

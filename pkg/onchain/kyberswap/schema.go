package kyberswap

// Schema: the struct that has closest data type with the external API

type GetSwapInfoRequestParams struct {
	AmountIn string `url:"amountIn"`
	To       string `url:"to"`
	TokenIn  string `url:"tokenIn"`
	TokenOut string `url:"tokenOut"`

	SlippageTolerance string `url:"slippageTolerance,omitempty"`
	ChargeFeeBy       string `url:"chargeFeeBy,omitempty"`
	ClientData        string `url:"clientData,omitempty"`
	Deadline          string `url:"deadline,omitempty"`
	FeeAmount         string `url:"feeAmount,omitempty"`
	FeeReceiver       string `url:"feeReceiver,omitempty"`
	IsInBps           bool   `url:"isInBps,omitempty"`
	SaveGas           string `url:"saveGas,omitempty"`
	// Dexes: By default, if empty, MetaAggregator version will be used.
	// If dexes are filled, MetaAggregator version will not be used.
	Dexes string `url:"dexes,omitempty"`
}

type SwapResponse struct {
	Pool              string `json:"pool"`
	TokenIn           string `json:"tokenIn"`
	TokenOut          string `json:"tokenOut"`
	SwapAmount        string `json:"swapAmount"`
	AmountOut         string `json:"amountOut"`
	LimitReturnAmount string `json:"limitReturnAmount"`
	MaxPrice          string `json:"maxPrice"`
	Exchange          string `json:"exchange"`
	PoolLength        int    `json:"poolLength"`
	PoolType          string `json:"poolType"`
}

type GetSwapInfoResponse struct {
	InputAmount     string           `json:"inputAmount"`
	OutputAmount    string           `json:"outputAmount"`
	TotalGas        int              `json:"totalGas"`
	GasPriceGwei    string           `jsons:"gasPriceGwei"`
	GasUsd          float64          `json:"gasUsd"`
	AmountInUsd     float64          `json:"amountInUsd"`
	AmountOutUsd    float64          `json:"amountOutUsd"`
	ReceivedUsd     float64          `json:"receivedUsd"`
	Swaps           [][]SwapResponse `json:"swaps"`
	EncodedSwapData string           `json:"encodedSwapData"`
	RouterAddress   string           `json:"routerAddress"`
}

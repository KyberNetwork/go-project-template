package kyberswap

import (
	"context"
	"fmt"
	"net/http"

	"github.com/KyberNetwork/go-project-template/pkg/httputil"
	"github.com/google/go-querystring/query"
)

const (
	// https://kyber-network.stoplight.io/
	kyberSwapAggregatorAPI = "https://aggregator-api.kyberswap.com"
)

type client struct {
	httpC *http.Client
}

type Client interface {
	GetSwapInfo(ctx context.Context, chainName string, inputParams GetSwapInfoInputParams) (*GetSwapInfoOutput, error)
}

func NewClient(httpC *http.Client) Client {
	if httpC == nil {
		httpC = http.DefaultClient
	}

	return &client{
		httpC: httpC,
	}
}

func (c *client) GetSwapInfo(
	ctx context.Context, chainName string, inputParams GetSwapInfoInputParams,
) (*GetSwapInfoOutput, error) {
	params := ToGetSwapInfoRequestParams(inputParams)
	paramBuilder, err := query.Values(params)
	if err != nil {
		return nil, fmt.Errorf("parse params err: %w", err)
	}
	apiURL := fmt.Sprintf("%s/%s/route/encode/?%s", kyberSwapAggregatorAPI, chainName, paramBuilder.Encode())

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodGet, apiURL, nil)
	if err != nil {
		return nil, fmt.Errorf("new request err: %w", err)
	}
	httpReq.Header.Add("accept-version", "Latest")

	var resp GetSwapInfoResponse
	if err = httputil.DoHTTPRequest(c.httpC, httpReq, &resp); err != nil {
		return nil, err
	}
	output, err := ToGetSwapInfoOutput(resp)
	if err != nil {
		return nil, err
	}
	return output, nil
}

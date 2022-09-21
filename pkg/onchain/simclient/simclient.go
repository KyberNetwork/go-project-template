package simclient

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/pkg/errors"
)

type reqMessage struct {
	JSONRPC string            `json:"jsonrpc"`
	ID      int               `json:"id"`
	Method  string            `json:"method"`
	Params  []json.RawMessage `json:"params"`
}

type Account struct {
	Nonce     string            `json:"nonce,omitempty"`
	Balance   string            `json:"balance,omitempty"`
	Code      string            `json:"code,omitempty"`
	State     map[string]string `json:"state,omitempty"`
	StateDiff map[string]string `json:"stateDiff,omitempty"`
}

type OverrideAccounts map[common.Address]Account

type roundTripperExt struct {
	c          *http.Client
	appendData json.RawMessage
}

func newRoundTripExt(c *http.Client, accounts OverrideAccounts) (http.RoundTripper, error) {
	data, err := json.Marshal(accounts)
	if err != nil {
		return nil, err
	}

	return &roundTripperExt{
		c:          c,
		appendData: data,
	}, nil
}

func (r roundTripperExt) RoundTrip(request *http.Request) (*http.Response, error) {
	rt := request.Clone(context.Background())
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}
	_ = request.Body.Close()

	if len(body) > 0 {
		rt.Body = io.NopCloser(bytes.NewBuffer(body))
		request.Body = io.NopCloser(bytes.NewBuffer(body))
	}

	var req reqMessage

	if err := json.Unmarshal(body, &req); err == nil {
		if req.Method == "eth_call" {
			req.Params = append(req.Params, r.appendData)
		}

		d2, err := json.Marshal(req)
		if err != nil {
			return nil, err
		}
		rt.ContentLength = int64(len(d2))
		rt.Body = io.NopCloser(bytes.NewBuffer(d2))
	}

	return r.c.Do(rt)
}

func NewSimClient(url string, client *http.Client, accounts OverrideAccounts) (*ethclient.Client, error) {
	round, err := newRoundTripExt(client, accounts)
	if err != nil {
		return nil, err
	}

	cc := &http.Client{Transport: round}
	r, err := rpc.DialHTTPWithClient(url, cc)
	if err != nil {
		return nil, errors.WithMessage(err, "simclient: dial rpc")
	}

	return ethclient.NewClient(r), nil
}

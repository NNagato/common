package blockchain

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	ethereum "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	// CustomContractAddress ...
	CustomContractAddress = ethereum.HexToAddress("0x221CA93C327eDe7d8f9296E2a790905CD7021105")
)

type reqMessage struct {
	JSONRPC string            `json:"jsonrpc"`
	ID      int               `json:"id"`
	Method  string            `json:"method"`
	Params  []json.RawMessage `json:"params"`
}

// AccountData ...
type AccountData struct {
	Nonce string `json:"nonce"`
	Code  string `json:"code"`
}

// accountDataWithNewRunTimeCode ...
func accountDataWithNewRunTimeCode(code string) json.RawMessage {
	m := make(map[string]AccountData)
	m["0x221CA93C327eDe7d8f9296E2a790905CD7021105"] = AccountData{
		Nonce: "0x112233",
		Code:  code,
	}
	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return b
}

type roundTripperExt struct {
	code string
	c    *http.Client
}

func (r *roundTripperExt) RoundTrip(request *http.Request) (*http.Response, error) {
	rt := request.Clone(context.Background())
	body, _ := ioutil.ReadAll(request.Body)
	_ = request.Body.Close()
	if len(body) > 0 {
		rt.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	}
	var req reqMessage
	if err := json.Unmarshal(body, &req); err == nil {
		if req.Method == "eth_call" {
			req.Params = append(req.Params, accountDataWithNewRunTimeCode(r.code))
		}
		d2, err := json.Marshal(req)
		if err != nil {
			panic(err)
		}
		rt.ContentLength = int64(len(d2))
		rt.Body = ioutil.NopCloser(bytes.NewBuffer(d2))
	}
	return r.c.Do(rt)
}

// NewCustomEtherClient ...
func NewCustomEtherClient(nodeEndpoint, code string) (*ethclient.Client, error) {
	cc := &http.Client{Transport: &roundTripperExt{c: &http.Client{}, code: code}}
	r, err := rpc.DialHTTPWithClient(nodeEndpoint, cc)
	if err != nil {
		return nil, err
	}
	client := ethclient.NewClient(r)
	return client, nil
}

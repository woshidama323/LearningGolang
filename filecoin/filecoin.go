package filecoin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/lotus/api/v0api"
)

type GetAddressByMinerID struct {
	ChainProvider struct {
		Api    v0api.FullNodeStruct
		Closer jsonrpc.ClientCloser
	}
}

func NewMiner(providerUrl string) (*GetAddressByMinerID, error) {

	var res v0api.FullNodeStruct
	closer, err := jsonrpc.NewMergeClient(context.Background(), providerUrl, "Filecoin",
		[]interface{}{
			&res.CommonStruct.Internal,
			&res.Internal,
		}, nil)

	if err != nil {
		return nil, err
	}
	return &GetAddressByMinerID{
		ChainProvider: struct {
			Api    v0api.FullNodeStruct
			Closer jsonrpc.ClientCloser
		}{
			Api:    res,
			Closer: closer,
		},
	}, nil
}

func GetMinerInfo(urlstr, minerid string) (interface{}, error) {
	getMinerInfo := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "Filecoin.StateMinerInfo",
		"params": []interface{}{
			minerid,
			nil,
		},
		"id": 1,
	}

	// res := &HeadInfo{}
	value, _ := json.Marshal(getMinerInfo)

	resp, err := http.Post(urlstr, "application/json", bytes.NewBuffer(value))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var FResp interface{}
	err = json.Unmarshal(body, &FResp)
	if err != nil {
		return nil, err
	}
	// if FResp.Error.IsError() {
	// 	return nil, FResp.Error.Error()
	// }
	// if FResp.Result == nil {
	// 	return nil, errors.New("Result is nil ")
	// }
	// _d, err := json.Marshal(FResp.Result)
	// if err != nil {
	// 	return nil, err
	// }

	// err = json.Unmarshal(_d, res)

	fmt.Printf("current result is :%v", FResp)
	fmt.Printf("current result body is :%v", string(body))
	return FResp, nil
}

func GetMinerAddressInfo(urlstr, minerid string) (interface{}, error) {

	getAddresses := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "Filecoin.ChainHead",
		"params":  []interface{}{},
		"id":      1,
	}

	// res := &HeadInfo{}
	value, _ := json.Marshal(getAddresses)

	resp, err := http.Post(urlstr, "application/json", bytes.NewBuffer(value))
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var FResp interface{}
	err = json.Unmarshal(body, &FResp)
	if err != nil {
		return nil, err
	}
	return FResp, nil

}

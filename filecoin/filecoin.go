package filecoin

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

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

	Sum(string(body))
	return FResp, nil
}

func GetMinerAddressInfo(urlstr, address string) (interface{}, error) {

	getAddresses := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "Filecoin.StateAccountKey",
		"params": []interface{}{
			address,
			nil,
		},
		"id": 1,
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

type ResType struct {
}
type MinerInfo struct {
	Owner                      string
	Worker                     string
	NewWorker                  string
	ControlAddresses           []string
	WorkerChangeEpoch          uint64
	PeerId                     *string
	Multiaddrs                 []string
	WindowPoStProofType        int
	SectorSize                 int
	WindowPoStPartitionSectors uint64
	ConsensusFaultElapsed      int
}

func Sum(in interface{}) (int64, error) {
	res := int64(0)
	v := reflect.ValueOf(in)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		res = v.Int()
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			sliceRes, err := Sum(v.Index(i).Interface())
			if err != nil {
				return 0, err
			}
			res = res + sliceRes
		}
	case reflect.Map:
		for _, k := range v.MapKeys() {
			mapRes, err := Sum(v.MapIndex(k).Interface())
			if err != nil {
				return 0, err
			}
			res = res + mapRes
		}
	default:
		return 0, fmt.Errorf("input passed was invalid.")
	}
	return res, nil
}

// {
//     "jsonrpc": "2.0",
//     "result": {
//         "Owner": "t03083",
//         "Worker": "t03083",
//         "NewWorker": "\u003cempty\u003e",
//         "ControlAddresses": null,
//         "WorkerChangeEpoch": -1,
//         "PeerId": "12D3KooWAcpx2A7SHM5dHtw5oyVm3Xk9Cu74uNhVdG361nSnQvgS",
//         "Multiaddrs": null,
//         "WindowPoStProofType": 8,
//         "SectorSize": 34359738368,
//         "WindowPoStPartitionSectors": 2349,
//         "ConsensusFaultElapsed": -1
//     },
//     "id": 1
// }

package dingtalk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

const (
	FilecoinLocalGroup = "https://oapi.dingtalk.com/robot/send?access_token=56155f0e75d5a64116f03272eddfc92fd1f5c9ddc73eb9338b1927726ef19ec8"
	FilcoinTestGroup   = "https://oapi.dingtalk.com/robot/send?access_token=e237d0b0cf64bedb7040d0ec431ff56842276dc84997ba1f55befb16f0c84ed9"
)

type DingTalkTextMsg struct {
	MsgType string                 `json:"msgtype"`
	Text    map[string]interface{} `json:"text"`
}

type DingTalkMarkDownMsg struct {
	MsgType  string                 `json:"msgtype"`
	MarkDown map[string]interface{} `json:"markdown"`
}

type DingTalkCardMsg struct {
	MsgType    string `json:"msgtype"`
	CardAction struct {
		Title       string `json:"title"`
		MarkDown    string `json:"markdown"`
		SingleTitle string `json:"single_title"`
		SingleUrl   string `json:"single_url"`
	} `json:"action_card"`
}
type DingTalkImageMsg struct {
	MsgType    string `json:"msgtype"`
	CardAction struct {
		Title       string `json:"title"`
		MarkDown    string `json:"markdown"`
		SingleTitle string `json:"single_title"`
		SingleUrl   string `json:"single_url"`
	} `json:"action_card"`
}

func SendToDingTalkTextMsg(input string) error {

	var jsonData = DingTalkTextMsg{
		MsgType: "text",
		Text: map[string]interface{}{
			"content": "robot",
		},
	}

	SendPostRequest(jsonData)
	return nil
}

func SendToDingTalkMarkDownMsg(input string) error {

	var jsonData = DingTalkMarkDownMsg{
		MsgType: "markdown",
		MarkDown: map[string]interface{}{
			"title": "robot",
			"text":  input,
		},
	}
	SendPostRequest(jsonData)
	return nil
}

func SendPostRequest(jsonData interface{}) error {
	jsonBytes, err := json.Marshal(jsonData)
	fmt.Printf("string is:%v\n", string(jsonBytes))
	if err != nil {
		return err
	}
	log.Println("....", string(jsonBytes))
	request, err := http.NewRequest("POST", FilcoinTestGroup, bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Printf("Error creating request")
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Charset", "UTF-8")
	client := &http.Client{}

	resp, err := client.Do(request)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New("what's it ")
	}
	defer resp.Body.Close()
	var j interface{}
	err = json.NewDecoder(resp.Body).Decode(&j)
	if err != nil {
		return err
	}
	if j != nil && j.(map[string]interface{})["errcode"].(float64) > 0 {
		log.Printf("++++ %v:", j)
		return errors.New("geterror")
	}
	log.Printf("++++ %v:", j)

	return nil
}

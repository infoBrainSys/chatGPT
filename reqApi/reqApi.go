package reqapi

import (
	"Learn/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/spf13/cast"
)

const (
	apiURL    = "https://api.openai.com/v1/completions"
	modelName = "text-davinci-003"
	prompt    = "什么是web3？web3时代哪些编程语言比较吃香？使用中文回答我"
)

func ReqApi() (string, error) {
	responseChan := make(chan config.ResponseJson)
	errorChan := make(chan error)
	go func() {
		client := http.Client{}
		data, err := json.Marshal(config.Config{
			Model:       modelName,
			Prompt:      prompt,
			MaxTokens:   2048,
			Temperature: 0.7,
			TopP:        0,
			N:           1,
		})
		if err != nil {
			errorChan <- err
			return
		}
		req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(data))
		if err != nil {
			errorChan <- err
			return
		}

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer KEY")
		since := time.Now()
		resp, err := client.Do(req)
		if err != nil {
			errorChan <- err
			return
		}
		fmt.Println(time.Since(since))
		defer resp.Body.Close()

		b, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			errorChan <- err
			return
		}
		var respJson config.ResponseJson
		json.Unmarshal(b, &respJson)

		responseChan <- respJson
	}()

	select {
	case respJson := <-responseChan:
		return cast.ToString(&respJson.Choices[0].Text), nil
	case err := <-errorChan:
		return "", err
	}
}

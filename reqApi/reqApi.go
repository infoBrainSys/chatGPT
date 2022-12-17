package reqapi

import (
	"bytes"
	"chatGPT/config"
	model "chatGPT/models"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cast"
)

const (
	apiURL = "https://api.openai.com/v1/completions"
)

var _cfg *config.Config

func ReqApi(prompt string) (string, error) {
	p := prompt

	// 定义 responseChan 和 errorChan 用于接受 goroutine 返回的请求内容
	responseChan := make(chan model.ResponseJson)
	errorChan := make(chan error)

	// 获取 viper.Viper 对象用于操作获取配置文件参数
	viper, err := _cfg.GetCfg()
	if err != nil {
		log.Println("viper get data failed, err :", err)
	}

	// 开启协程
	go func() {
		// 实例化一个 client 用于发起请求
		client := http.Client{}

		// 序列化配置文件作为 request 请求体
		data, err := json.Marshal(model.Request{
			Model:       viper.GetString("model"),
			Prompt:      p,
			MaxTokens:   viper.GetInt("max_tokens"),
			Temperature: viper.GetFloat64("temperature"),
			TopP:        float64(viper.GetInt("top_p")),
			N:           viper.GetInt("n"),
		})
		if err != nil {
			// 出错则返回 err 到 errorChan
			errorChan <- err
			return
		}

		// 设置请求 method, URl, 请求体
		req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(data))
		if err != nil {
			// 出错则返回 err 到 errorChan
			errorChan <- err
			return
		}

		// 设置请求头参数
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", "Bearer "+viper.GetString("key"))
		// fmt.Println("================")
		// fmt.Println("Bearer " + viper.GetString("key"))
		// fmt.Println("================")

		// 开发过程中用于检测整体一个相应消耗的时间
		// since := time.Now()

		// 发起请求并且获取相应内容
		resp, err := client.Do(req)
		if err != nil {
			errorChan <- err
			return
		}
		// fmt.Println(time.Since(since))
		defer resp.Body.Close()

		// 读取响应体
		b, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			// 出错则返回 err 到 errorChan
			errorChan <- err
			return
		}

		// 将读取内容反序列化到 ResponseJson 结构体
		var respJson model.ResponseJson
		json.Unmarshal(b, &respJson)

		// 将内容写入到 responseChan 中
		responseChan <- respJson
	}()

	select {

	// 只读响应体中的 Text 并且 return
	case respJson := <-responseChan:
		return cast.ToString(&respJson.Choices[0].Text), nil

	// 出错则 return err
	case err := <-errorChan:
		return "", err
	}
}

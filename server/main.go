package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// 受け取るURLを指定（接続確認にも使用）
	url := "http://host.docker.internal:1323"

	// GETリクエストを送信して接続確認
	response, err := http.Get(url)
	if err != nil {
		handleError(url, err)
	}
	// CloseしてTCPコネクションを開きっぱなしにしない
	defer response.Body.Close()

	// ステータスコードを確認
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error Response:", response.Status)
		return
	}

	// レスポンスボディを読み込み
	body, err := io.ReadAll(response.Body)
	if err != nil {
		handleError(url, err)
		return
	}

	// JSONを構造体にエンコード
	var responseData map[string]interface{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		handleError(url, err)
		return
	}
	fmt.Println(responseData["message"])

	// JSON文字列を定義
	devices := []map[string]interface{}{
		{"id": 1, "name": "device1"},
		{"id": 2, "name": "device2"},
	}

	v, err := json.Marshal(devices)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return
	}

	response, err = http.Post(url, "application/json", bytes.NewBuffer(v))
	if err != nil {
		fmt.Println("HTTP POST error:", err)
		return
	}
	defer response.Body.Close()
	fmt.Println("Success! Status:", response.Status)
}

func handleError(url string, err error) {
	log.Fatal(err)
}

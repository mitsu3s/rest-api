package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	// 受け取るURLを指定（接続確認にも使用）
	url := "http://localhost:1323"
	// url := "http://host.docker.internal:1323"

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

	// JSONレスポンスを保存するファイルを作成
	file, err := os.Create("response.json")
	if err != nil {
		handleError(url, err)
		return
	}

	defer file.Close()

	// JSONをエンコードしてファイルに書き込み
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(responseData); err != nil {
		handleError(url, err)
	}
	handleSuccess(url)

	fmt.Println("Success!")
}

func handleError(url string, err error) {
	sendNotification(url, "Error", err)
	log.Fatal(err)
}

func handleSuccess(url string) {
	sendNotification(url, "OK", nil)
}

func sendNotification(url, status string, err error) {
	// 送信するJSONを作成
	notification := map[string]string{
		"status":  status,
		"message": "",
	}
	if err != nil {
		notification["message"] = err.Error()
	}

	// JSON形式にエンコード
	jsonValue, _ := json.Marshal(notification)

	// POSTリクエストを送信
	_, err = http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatal(err)
	}
}

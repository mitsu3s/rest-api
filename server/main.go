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
		log.Fatal(err)
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
		log.Fatal(err)
	}

	// JSONを構造体にエンコード
	var responseData map[string]interface{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		log.Fatal(err)
	}

	// JSONレスポンスを保存するファイルを作成
	file, err := os.Create("response.json")
	if err != nil {
		writeError(url, err)
		log.Fatal(err)
	}

	defer file.Close()

	// JSONをエンコードしてファイルに書き込み
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(responseData); err != nil {
		writeError(url, err)
		log.Fatal(err)
	}
	writeSuccess(url)

	fmt.Println("Success!")
}

func writeSuccess(url string) {
	notification := map[string]string{"status": "OK"}
	jsonValue, _ := json.Marshal(notification)

	_, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatal(err)
	}
}

func writeError(url string, err error) {
	notifivation := map[string]string{"status": "Error", "message": err.Error()}
	jsonValue, _ := json.Marshal(notifivation)

	_, err = http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatal(err)
	}
}

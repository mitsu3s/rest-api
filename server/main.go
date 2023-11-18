package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// 受け取るURLを指定（接続確認にも使用）
	// url := "http://localhost:1323"
	url := "http://host.docker.internal:1323"

	// GETリクエストを送信して接続確認
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	// CloseしてTCPコネクションを開きっぱなしにしない
	defer response.Body.Close()

	// ステータスコードを確認
	if response.StatusCode != 200 {
		fmt.Println("Error Response:", response.Status)
		return
	}

	// レスポンスボディを読み込む
	body, _ := io.ReadAll(response.Body)

	// JSONを構造体にエンコード
	if json.Unmarshal(body, &response) != nil {
		log.Fatal(err)
	}

	fmt.Println(response)
}

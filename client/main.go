package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instanceを作成
	e := echo.New()

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	// ルーティングを設定
	e.GET("/", func(c echo.Context) error {
		message := map[string]string{
			"message": "Connected!",
		}
		return c.JSON(http.StatusOK, message)
	})

	e.POST("/", func(c echo.Context) error {
		body, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}

		if err := writeDevice(body); err != nil {
			fmt.Println("Error writing to JSON file:", err)
			return err
		}
		return c.String(http.StatusOK, "OK")
	})

	// サーバーを開始
	e.Start(":1323")
}

func writeDevice(data []byte) error {
	var parsedData interface{}
	if err := json.Unmarshal(data, &parsedData); err != nil {
		return err
	}

	// JSONファイルに書き込む
	file, err := os.Create("devices.json")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(parsedData); err != nil {
		return err
	}

	return nil
}

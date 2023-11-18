package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instanceを作成
	e := echo.New()

	// ルーティングを設定
	e.GET("/", func(c echo.Context) error {
		message := map[string]string{
			"message": "Hello, World!",
		}
		return c.JSON(http.StatusOK, message)
	})
	// サーバーを開始
	e.Logger.Fatal(e.Start(":1323"))
}

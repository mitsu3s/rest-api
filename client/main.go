package main

import (
	"fmt"
	"net/http"

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
			"ID":   "1",
			"Name": "User1",
		}
		return c.JSON(http.StatusOK, message)
	})

	e.POST("/", func(c echo.Context) error {
		var notification map[string]string
		if err := c.Bind(&notification); err != nil {
			return err
		}

		if notification["status"] == "OK" {
			fmt.Println(notification["status"])
		} else {
			fmt.Println(notification["message"])
		}
		return nil
	})

	// サーバーを開始
	// e.Logger.Fatal(e.Start(":1323"))
	e.Start(":1323")
}

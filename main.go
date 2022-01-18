package main

import (
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/PersonalSmartContractTest/api"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}
	conn, err := api.NewApi(common.HexToAddress("0x549129cecd93c366a4fb327970caa8604a54dbb2"), client)
	if err != nil {
		panic(err)
	}
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/greet/:message", func(c echo.Context) error {
		message := c.Param("message")
		reply, err := conn.Greet(&bind.CallOpts{}, message)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, reply)
	})
	e.GET("/hello", func(c echo.Context) error {
		reply, err := conn.Hello(&bind.CallOpts{})
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, reply) // Hello World
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

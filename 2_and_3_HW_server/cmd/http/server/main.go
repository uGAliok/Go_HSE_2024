package main

import (
	"Go_HSE_2024/2_and_3_HW_server/accounts"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	connectionString := "host=0.0.0.0 port=5432 dbname=postgres user=postgres password=mysecretpassword"

	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	defer func() {
		_ = conn.Close(context.Background())
	}()

	// для http сервера 2 дз
	accountsHandler := accounts.New()

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/account/get", func(c echo.Context) error {
		return accountsHandler.GetAccount(c, conn)
	})
	e.POST("/account/create", func(c echo.Context) error {
		return accountsHandler.CreateAccount(c, conn)
	})
	e.DELETE("/account/delete", func(c echo.Context) error {
		return accountsHandler.DeleteAccount(c, conn)
	})
	e.PATCH("/account/patch", func(c echo.Context) error {
		return accountsHandler.PatchAccount(c, conn)
	})
	e.PUT("/account/change", func(c echo.Context) error {
		return accountsHandler.ChangeAccount(c, conn)
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

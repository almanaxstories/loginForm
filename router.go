package main

import (
	// jwt "github.com/golang-jwt/jwt/v4"

	"database/sql"
	"fmt"

	"github.com/almanaxstories/loginForm/api"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var db *sql.DB

func main() {
	db, err := sqlx.Open("pgx", "postgres://admin:qwert@127.0.0.1:5432/usersInfoDB?sslmode=disable")
	if err != nil {
		fmt.Printf("Connection to DB failed. Error message: %s", err)
	} else {
		fmt.Println("Successfully connected to DB!")
	}
	defer db.Close()

	userRepo := api.NewUserRepo(db)
	createUserHandler := api.CreateUser(userRepo)

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.File("./assets/index.html")
	})

	e.POST("/", api.LoginHandler(userRepo))

	e.GET("/register", func(c echo.Context) error {
		return c.File("./assets/register.html")
	})

	e.POST("/register", createUserHandler)

	e.Use(middleware.Static("./static"))
	e.Logger.Fatal(e.Start(":8090"))
}

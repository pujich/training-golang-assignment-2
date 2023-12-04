package main

import (
	"golang-second-assigment/config"
	"golang-second-assigment/controllers"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var PORT = os.Getenv("PORTAPP")

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	PORT = ":" + os.Getenv("PORTAPP")
	config.ConnectGorm()
	config.StartingApps()

	r := echo.New()

	r.GET("/", func(ctx echo.Context) error {
		data := "Seperti Menarik wkwk"
		return ctx.String(http.StatusOK, data)
	})
	r.GET("/order", controllers.GetAllOrders)

	r.POST("/addorder", controllers.AddOrderByGorm)
	r.PUT("/updateorder/:id", controllers.UpdateOrderByGorm)
	r.DELETE("/deleteorder/:id", controllers.DeleteOrderByGorm)

	r.Logger.Fatal(r.Start(PORT))
}
